package agent

import (
	"Open_IM/pkg/base_info"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/agent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 获取推广员游戏商城咖豆配置
func AgentGameShopBeanConfig(c *gin.Context) {
	params := base_info.AgentGameShopBeanConfigReq{}
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

	req := &rpc.AgentGameBeanShopConfigReq{
		UserId:      "",
		AgentNumber: params.AgentNumber,
		OperationId: operationId,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.AgentGameBeanShopConfig(c, req)
	if err != nil {
		log.NewError(operationId, "AgentGameBeanShopConfig failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": RpcResp.BeanShopConfig})
	return
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

// 推广员自定义咖豆商城配置
func AgentDiyBeanShopConfig(c *gin.Context) {
	operationId := c.GetString("operationId")
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": errMsg})
		return
	}

	req := &rpc.AgentDiyBeanShopConfigReq{
		UserId:      c.GetString("userId"),
		OperationId: operationId,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.AgentDiyBeanShopConfig(c, req)
	if err != nil {
		log.NewError(operationId, "AgentDiyBeanShopConfig failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": RpcResp})
	return
}

// 咖豆账户明细详情列表
func AgentBeanAccountRecordList(c *gin.Context) {
	params := base_info.AgentBeanAccountRecordListReq{}
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

	//默认当天
	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02")
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Size == 0 || params.Size > 100 {
		params.Size = 20
	}

	req := &rpc.AgentBeanAccountRecordListReq{
		UserId:       c.GetString("userId"),
		Date:         params.Date,
		Page:         params.Page,
		Size:         params.Size,
		BusinessType: params.BusinessType,
		OperationId:  operationId,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.AgentBeanAccountRecordList(c, req)
	if err != nil {
		log.NewError(operationId, "AgentBeanAccountRecordList failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": RpcResp})
	return
}
