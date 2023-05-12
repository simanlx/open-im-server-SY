package agent

import (
	"Open_IM/internal/api/common"
	"Open_IM/pkg/base_info"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/agent"
	"Open_IM/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 推广员成员购买咖豆回调(推广员商城) - 互娱回调
func ChessPurchaseBeanNotify(c *gin.Context) {
	params := base_info.ChessPurchaseBeanNotifyReq{}
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

	req := &rpc.ChessPurchaseBeanNotifyReq{
		OrderNo:       params.OrderNo,
		NcountOrderNo: params.NcountOrderNo,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.ChessPurchaseBeanNotify(c, req)
	if err != nil {
		log.NewError(operationId, "ChessPurchaseBeanNotify failed ", err.Error(), req.String())
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

// 推广员成员购买咖豆回调(平台商城) - 互娱回调
func PlatformPurchaseBeanNotify(c *gin.Context) {
	params := base_info.PlatformPurchaseBeanNotifyReq{}
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

	req := &rpc.PlatformPurchaseBeanNotifyReq{
		ChessOrderNo:   params.ChessOrderNo,
		NcountOrderNo:  params.NcountOrderNo,
		AgentNumber:    params.AgentNumber,
		ChessUserId:    params.ChessUserId,
		BeanNumber:     params.BeanNumber,
		GiveBeanNumber: params.GiveBeanNumber,
		Amount:         params.Amount,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.PlatformPurchaseBeanNotify(c, req)
	if err != nil {
		log.NewError(operationId, "PlatformPurchaseBeanNotify failed ", err.Error(), req.String())
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

// 推广员充值咖豆 - 新生支付回调
func RechargeNotify(c *gin.Context) {
	params := base_info.NcountNotifyReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	operationId := c.GetString("operationId")
	log.Info(operationId, "推广员充值咖豆-新生支付回调:", utils.JsonFormat(params))

	if params.Status != 200 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "支付状态非200错误"})
		return
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": errMsg})
		return
	}

	req := &rpc.RechargeNotifyReq{
		OrderNo:       params.MerOrderId,
		NcountOrderNo: params.OrderId,
		PayTime:       params.PayTime,
		Amount:        params.Amount,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.RechargeNotify(c, req)
	if err != nil {
		log.NewError(operationId, "RechargeNotify failed ", err.Error(), req.String())
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

// 推广员提现余额 - 新生支付回调
func WithdrawNotify(c *gin.Context) {
	params := base_info.NcountNotifyReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	operationId := c.GetString("operationId")
	log.Info(operationId, "推广员提现余额-新生支付回调:", utils.JsonFormat(params))

	if params.Status != 200 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "支付状态非200错误"})
		return
	}

	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImAgentName, operationId)
	if etcdConn == nil {
		errMsg := operationId + "getcdv3.GetDefaultConn == nil"
		log.NewError(operationId, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": errMsg})
		return
	}

	req := &rpc.WithdrawNotifyReq{
		OrderNo:       params.MerOrderId,
		NcountOrderNo: params.OrderId,
		PayTime:       params.PayTime,
		Amount:        params.Amount,
	}

	client := rpc.NewAgentSystemServiceClient(etcdConn)
	RpcResp, err := client.WithdrawNotify(c, req)
	if err != nil {
		log.NewError(operationId, "WithdrawNotify failed ", err.Error(), req.String())
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	// handle rpc err
	if common.HandleAgentCommonRespErr(RpcResp.CommonResp, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": ""})
	return
}
