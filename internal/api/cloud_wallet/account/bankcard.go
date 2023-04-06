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

// 绑定银行卡
func BindUserBankCard(c *gin.Context) {
	params := account.BindUserBankCardReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.BindUserBankcardReq{
		UserId:            params.UserId,
		CardOwner:         params.CardOwner,
		BankCardNumber:    params.BankCard,
		Mobile:            params.Mobile,
		CardAvailableDate: params.CardAvailableDate,
		Cvv2:              params.Cvv2,
		OperationID:       params.OperationID,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.BindUserBankcard(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 绑定银行卡code确认
func BindUserBankcardConfirm(c *gin.Context) {
	params := account.BindUserBankcardConfirmReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.BindUserBankcardConfirmReq{
		UserId:      params.UserId,
		BankCardId:  params.BankCardId,
		SmsCode:     params.Code,
		MerUserIp:   c.ClientIP(),
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
	RpcResp, err := client.BindUserBankcardConfirm(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcardConfirm failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 解绑银行卡
func UnBindUserBankcard(c *gin.Context) {
	params := account.UnBindUserBankcardReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.UnBindingUserBankcardReq{
		UserId:        params.UserId,
		BindCardAgrNo: params.BindCardAgrNo,
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
	RpcResp, err := client.UnBindingUserBankcard(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "UnBindingUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}
