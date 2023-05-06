package agent

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/agent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 获取推广员商城咖豆配置
func AgentShopBeanConfig(c *gin.Context) {

}

// 获取平台咖豆商城配置
func PlatformBeanShopConfig(c *gin.Context) {
	operationId := c.GetString("operationId")
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": errMsg})
		return
	}

	req := &rpc.PlatformBeanShopConfigReq{
		UserId:      "",
		OperationId: operationId,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.PlatformBeanShopConfig(c, req)
	if err != nil {
		log.NewError(operationId, "AgentAccountIncomeChart failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": RpcResp.BeanShopConfig})
	return
}

// 推广员咖豆商城配置
func AgentBeanShopConfig(c *gin.Context) {

}
