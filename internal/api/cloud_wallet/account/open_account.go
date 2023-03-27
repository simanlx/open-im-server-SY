package account

import (
	utils2 "Open_IM/internal/utils"
	"Open_IM/pkg/base_info/account"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/token_verify"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 获取账户信息
func Account(c *gin.Context) {
	operationID := utils.OperationIDGenerator()

	//获取token用户id
	ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), operationID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errInfo})
		return
	}

	req := &rpc.UserNcountAccountReq{UserId: userId}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, operationID)
	if etcdConn == nil {
		errMsg := operationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.UserNcountAccount(context.Background(), req)
	if err != nil {
		log.NewError(operationID, "IdCardRealNameAuth failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
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

	req := &rpc.IdCardRealNameAuthReq{}
	utils.CopyStructFields(req, &params)

	//获取token用户id
	//ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), params.OperationID)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errInfo})
	//	return
	//}
	userId := "cccccc"
	req.UserId = userId

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.IdCardRealNameAuth(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "IdCardRealNameAuth failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
	return
}

// 设置支付密码
func SetPaymentSecret(c *gin.Context) {
	params := account.SetPaymentSecretReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.SetPaymentSecretReq{}
	utils.CopyStructFields(req, &params)

	//获取token用户id
	ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), params.OperationID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errInfo})
		return
	}
	req.UserId = userId

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.SetPaymentSecret(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "IdCardRealNameAuth failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
	return
}

// 获取账户余额
func UserAccountBalance(c *gin.Context) {
	operationID := utils.OperationIDGenerator()

	//调新生支付接口获取用户余额
	//获取token用户id
	ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), operationID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errInfo})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"balance": 99, "user_id": userId})
	return
}

// 云钱包明细：云钱包收支情况
func CloudWalletRecordList(c *gin.Context) {
	operationID := utils.OperationIDGenerator()

	//获取token用户id
	ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), operationID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errInfo})
		return
	}

	req := &rpc.CloudWalletRecordListReq{
		UserId:   userId,
		Date:     "2023-04",
		PageNum:  0,
		PageSize: 0,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.CloudWalletRecordList(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "CloudWalletRecordList failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
	return
}
