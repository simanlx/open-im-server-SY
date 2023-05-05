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

// 推广员申请提交
func AgentApply(c *gin.Context) {
	params := base_info.AgentApplyReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	operationId := c.GetString("operationId")
	req := &rpc.AgentApplyReq{
		UserId:      c.GetString("userId"),
		Name:        params.Name,
		Mobile:      params.Mobile,
		ChessUserId: params.ChessUserId,
		OperationId: operationId,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.AgentApply(c, req)
	if err != nil {
		log.NewError(operationId, "AgentApply failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleAgentCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 获取当前用户的推广员信息以及绑定关系
func GetUserAgentInfo(c *gin.Context) {
	params := base_info.GetUserAgentInfo{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	operationId := c.GetString("operationId")
	req := &rpc.GetUserAgentInfoReq{
		UserId:      c.GetString("userId"),
		ChessUserId: params.ChessUserId,
		OperationId: operationId,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.GetUserAgentInfo(c, req)
	if err != nil {
		log.NewError(operationId, "GetUserAgentInfo failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 绑定推广员
func BindAgentNumber(c *gin.Context) {
	params := base_info.BindAgentNumberReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	operationId := c.GetString("operationId")
	req := &rpc.BindAgentNumberReq{
		UserId:        c.GetString("userId"),
		AgentNumber:   params.AgentNumber,
		ChessUserId:   params.ChessUserId,
		ChessNickname: params.ChessNickname,
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errMsg})
		return
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.BindAgentNumber(c, req)
	if err != nil {
		log.NewError(operationId, "BindAgentNumber failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleAgentCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}
