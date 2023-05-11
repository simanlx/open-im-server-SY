package agent

import (
	"Open_IM/internal/api/common"
	"Open_IM/pkg/base_info"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/agent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 互娱商城购买咖豆下单(预提交)
func ChessShopPurchaseBean(c *gin.Context) {
	params := base_info.ChessShopPurchaseBeanReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	operationId := c.GetString("operationId")
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": errMsg})
		return
	}

	req := &rpc.ChessShopPurchaseBeanReq{
		UserId:       c.GetString("userId"),
		ChessOrderNo: params.ChessOrderNo,
		ChessUserId:  params.ChessUserId,
		ConfigId:     params.ConfigId,
		OperationId:  operationId,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.ChessShopPurchaseBean(c, req)
	if err != nil {
		log.NewError(operationId, "ChessShopPurchaseBean failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleAgentCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": RpcResp})
	return
}