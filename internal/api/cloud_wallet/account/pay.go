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
	"strconv"
	"strings"
)

// 充值
func ChargeAccount(c *gin.Context) {
	params := account.UserRechargeReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	//// 1.校验参数
	//if userid == "" || bankcardID == "" || amount == "" || OperateId == "" {
	//	return nil, errors.New("param is nil")
	//}

	req := &rpc.UserRechargeReq{
		UserId:      params.UserId,
		BankCardID:  params.BindCardAgrNo,
		Amount:      strconv.Itoa(int(params.Amount)),
		OperationID: params.OperationID,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.UserRecharge(context.Background(), req)
	if err != nil {
		log.NewError(req.OperationID, "UserRecharge failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}
