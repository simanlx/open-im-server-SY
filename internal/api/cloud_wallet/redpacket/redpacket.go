package redpacket

import (
	"Open_IM/internal/api/common"
	"Open_IM/pkg/agora"
	"Open_IM/pkg/base_info/redpacket_struct"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/tencent_cloud"
	"Open_IM/pkg/utils"
	"context"
	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 发送红包接口
func SendRedPacket(c *gin.Context) {
	params := redpacket_struct.SendRedPacket{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("0", "ChargeNotify", err.Error(), params)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// 获取用户ID
	/*	ok, UserID, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), params.OperationID)
		if !ok {
			errMsg := params.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
			log.NewError(params.OperationID, errMsg)
			c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
			return
		}*/
	UserID := "10018"
	if params.UserId != "" {
		UserID = params.UserId
	}
	// 复制结构体
	req := &rpc.SendRedPacketReq{}
	err := utils.CopyStructFields(req, &params)
	if err != nil {
		log.NewError(params.OperationID, "CopyStructFields failed ", err.Error(), params)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	req.UserId = UserID

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.SendRedPacket(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "调用失败 ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 抢红包接口
func ClickRedPacket(c *gin.Context) {
	params := redpacket_struct.ClickRedPacketReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("0", "ChargeNotify", err.Error(), params)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// 获取用户ID
	/*	ok, UserID, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), params.OperateID)
		if !ok {
			errMsg := params.OperateID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
			log.NewError(params.OperateID, errMsg)
			c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
			return
		}*/

	// 复制结构体
	req := &rpc.ClickRedPacketReq{}
	utils.CopyStructFields(req, &params)

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperateID)
	if etcdConn == nil {
		errMsg := req.OperateID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperateID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.ClickRedPacket(context.Background(), req)
	if err != nil {
		log.NewError(req.OperateID, "BindUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": RpcResp.CommonResp.ErrCode, "errMsg": RpcResp.CommonResp.ErrMsg})
	return
}

// 红包领取明细
func RedPacketReceiveDetail(c *gin.Context) {
	params := redpacket_struct.RedPacketReceiveDetailReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	req := &rpc.RedPacketReceiveDetailReq{
		UserId:      userId,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
		OperationID: params.OperationID,
	}

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.RedPacketReceiveDetail(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "RedPacketReceiveDetail failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 获取红包详情接口
func GetRedPacketInfo(c *gin.Context) {
	params := redpacket_struct.RedPacketInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	req := &rpc.RedPacketInfoReq{
		UserId:      userId,
		PacketId:    params.PacketId,
		OperationID: params.OperationID,
	}

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.RedPacketInfo(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "RedPacketInfo failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "timestamp": time.Now().Unix(), "data": RpcResp})
	return
}

// 禁止用户抢红包
func BanGroupClickRedPacket(c *gin.Context) {
	params := redpacket_struct.BanRedPacketReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// 获取用户ID
	/*	ok, UserID, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), params.OperateID)
		if !ok {
			errMsg := params.OperateID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
			log.NewError(params.OperateID, errMsg)
			c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
			return
		}*/
	UserID := "10000"
	if params.UserId == "" {
		UserID = "1000"
	} else {
		UserID = params.UserId
	}

	req := &rpc.ForbidGroupRedPacketReq{
		UserId:      UserID,
		GroupId:     params.GroupId,
		Forbid:      params.Forbid,
		OperationID: params.OperationID,
	}

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.ForbidGroupRedPacket(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "RedPacketInfo failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 这里是获取声网（完成）
func GetAgoraToken(c *gin.Context) {
	params := redpacket_struct.AgoraTokenReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	var role rtctokenbuilder.Role
	switch params.Role {
	case 1:
		role = rtctokenbuilder.RolePublisher
	case 2:
		role = rtctokenbuilder.RoleSubscriber
	}

	// 生成token
	result, appid, err := agora.GenerateRtcToken(userId, params.OperationID, params.Channel_name, role)
	if err != nil {
		log.NewError(params.OperationID, "RedPacketInfo failed ", err.Error(), params)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}
	resp := redpacket_struct.AgoraTokenResp{
		Token: result,
		AppID: appid,
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "msg": "获取成功", "data": resp})
}

// 翻译音频 （完成）
func TranslateVideo(c *gin.Context) {
	params := redpacket_struct.TencentMsgEscapeReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	_, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	// 消息翻译
	result, err := tencent_cloud.GetTencentCloudTranslate(params.ContentUrl, params.OperationID)
	if err != nil {
		log.NewError(params.OperationID, "RedPacketInfo failed ", err.Error(), params)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 200, "errMsg": "获取成功", "data": result})
}

// 获取版本号
func GetVersion(c *gin.Context) {
	// param in : 版本号
	// param out : 最新版本号、下载地址、更新内容、是否强制更新

	params := redpacket_struct.GetVersionReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	rpcReq := rpc.GetVersionReq{
		Version:     params.VersionCode,
		OperationID: params.OperationID,
	}

	// etcdConn
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, params.OperationID)
	if etcdConn == nil {
		errMsg := params.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(params.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.GetVersion(context.Background(), &rpcReq)
	if err != nil {
		log.NewError(params.OperationID, "GetVersion failed ", err.Error(), params)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
}
