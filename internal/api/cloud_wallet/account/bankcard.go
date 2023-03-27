package account

import (
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

// 获取用户银行卡列表
func GetUserBankCardList(c *gin.Context) {
	userId, _ := c.Get("userID")
	req := &rpc.GetUserBankcardListReq{
		UserId:      userId.(string),
		OperationID: "",
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.GetUserBankcardList(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "GetUserBankcardList failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
	return
}

// 绑定银行卡
func BindUserBankCard(c *gin.Context) {
	userId, _ := c.Get("userID")

	params := account.BindUserBankCardReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.BindUserBankcardReq{
		UserId:         userId.(string),
		CardOwner:      params.CardOwner,
		BankCardType:   params.BankCardType,
		BankCardNumber: params.BankCardNumber,
		Mobile:         params.Mobile,
		OperationID:    "",
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.BindUserBankcard(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
	return
}

// 绑定银行卡code确认
func BindUserBankcardConfirm(c *gin.Context) {
	userId, _ := c.Get("userID")

	params := account.BindUserBankcardConfirmReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.BindUserBankcardConfirmReq{
		UserId:      userId.(string),
		BankCardId:  params.BankCardId,
		MerOrderId:  "",
		SmsCode:     params.Code,
		OperationID: "",
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.BindUserBankcardConfirm(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcardConfirm failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RpcResp)
	return

}
