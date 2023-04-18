package account

import (
	"Open_IM/internal/api/common"
	utils2 "Open_IM/internal/utils"
	"Open_IM/pkg/base_info/account"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 获取账户信息
func Account(c *gin.Context) {
	params := account.AccountReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	req := &rpc.UserNcountAccountReq{UserId: userId, OperationID: params.OperationID}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, params.OperationID)
	if etcdConn == nil {
		errMsg := params.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(params.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.UserNcountAccount(context.Background(), req)
	if err != nil {
		log.NewError(params.OperationID, "UserNcountAccount failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 身份证实名认证
func IdCardRealNameAuth(c *gin.Context) {
	params := account.IdCardRealNameAuthReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	if len(params.RealName) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "实名姓名错误"})
		return
	}

	//验证身份证
	if !utils2.VerifyByIDCard(params.IdCard) {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "身份证号码错误"})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	req := &rpc.IdCardRealNameAuthReq{
		UserId:      userId,
		Mobile:      params.Mobile,
		IdCard:      params.IdCard,
		RealName:    params.RealName,
		OperationID: params.OperationID,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.IdCardRealNameAuth(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "IdCardRealNameAuth failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 设置支付密码
func SetPaymentSecret(c *gin.Context) {
	params := account.SetPaymentSecretReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//密码
	if len(params.PaymentSecret) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "支付密码参数错误"})
		return
	}

	//设置类型(1设置密码、2忘记密码smsCode设置、3修改密码)
	if params.Type == 2 && len(params.Code) != 6 {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "验证码参数错误"})
		return
	}

	if params.Type == 3 && len(params.OriginalPaymentSecret) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "原支付密码参数错误"})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	req := &rpc.SetPaymentSecretReq{
		UserId:                userId,
		PaymentSecret:         params.PaymentSecret,
		OperationID:           params.OperationID,
		Type:                  params.Type,
		Code:                  params.Code,
		OriginalPaymentSecret: params.OriginalPaymentSecret,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.SetPaymentSecret(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "SetPaymentSecret failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 校验支付密码
func CheckPaymentSecret(c *gin.Context) {
	params := account.CheckPaymentSecretReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	//6位数密码
	req := &rpc.CheckPaymentSecretReq{
		UserId:        userId,
		PaymentSecret: params.PaymentSecret,
		OperationID:   params.OperationID,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.CheckPaymentSecret(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "CheckPaymentSecret failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 云钱包明细：云钱包收支情况
func CloudWalletRecordList(c *gin.Context) {
	params := account.CloudWalletRecordListReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//解析token、获取用户id
	userId, ok := common.ParseImToken(c, params.OperationID)
	if !ok {
		return
	}

	req := &rpc.CloudWalletRecordListReq{
		UserId:      userId,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
		Page:        params.Page,
		Size:        params.Size,
		OperationID: params.OperationID,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.CloudWalletRecordList(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "CloudWalletRecordList failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}
