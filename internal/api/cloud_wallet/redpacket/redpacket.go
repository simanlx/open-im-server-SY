package redpacket

import (
	"Open_IM/pkg/base_info/redpacket_struct"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

// 发送红包接口
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

	UserID := "10018"

	// 复制结构体
	req := &rpc.ClickRedPacketReq{}
	utils.CopyStructFields(req, &params)
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
	RpcResp, err := client.ClickRedPacket(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 红包领取明细
func RedPacketReceiveDetail(c *gin.Context) {
	params := redpacket_struct.RedPacketReceiveDetailReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.RedPacketReceiveDetailReq{
		UserId:      params.UserId,
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
