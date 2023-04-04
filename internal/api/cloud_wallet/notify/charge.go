package notify

import (
	"Open_IM/pkg/base_info/notify"
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

//充值回调
func ChargeNotify(c *gin.Context) {
	params := notify.ChargeNotifyReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("0", "ChargeNotify", err.Error(), params)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// 复制结构体
	req := &rpc.ChargeNotifyReq{}
	utils.CopyStructFields(req, &params)
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
	RpcResp, err := client.ChargeNotify(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 提现回调
func WithDrawNotify(c *gin.Context) {
	params := notify.WithdrawNotifyReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("0", "ChargeNotify", err.Error(), params)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	// 复制结构体
	req := &rpc.DrawNotifyReq{}
	utils.CopyStructFields(req, &params)
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
	RpcResp, err := client.WithDrawNotify(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "BindUserBankcard failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}
