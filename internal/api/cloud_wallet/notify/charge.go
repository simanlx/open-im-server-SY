package notify

import (
	"Open_IM/pkg/base_info/notify"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	rpc "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 充值回调
func ChargeNotify(c *gin.Context) {
	params := notify.ChargeNotifyReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("0", "ChargeNotify", err.Error(), params)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.ChargeNotifyReq{
		MerOrderId:     params.MerOrderId,
		ResultCode:     params.ResultCode,
		ErrorCode:      params.ErrorCode,
		ErrorMsg:       params.ErrorMsg,
		NcountOrderId:  params.NcountOrderId,
		TranAmount:     params.TranAmount,
		SubmitTime:     params.SubmitTime,
		TranFinishTime: params.TranFinishTime,
		FeeAmount:      params.FeeAmount,
	}

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, "0000")
	if etcdConn == nil {
		errMsg := "0000" + "getcdv3.GetDefaultConn == nil"
		log.NewError("0000", errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.ChargeNotify(context.Background(), req)
	if err != nil {
		log.NewError("0000", "ChargeNotify failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}

// 提现回调
func WithDrawNotify(c *gin.Context) {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println("WithDrawNotify Body", string(data))
	//log.Error("0", "WithDrawNotify Body", string(data))
	params := notify.WithdrawNotifyReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("0", "WithDrawNotify", err.Error(), params)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	req := &rpc.DrawNotifyReq{
		MerOrderId:     params.MerOrderId,
		ResultCode:     params.ResultCode,
		ErrorCode:      params.ErrorCode,
		ErrorMsg:       params.ErrorMsg,
		NcountOrderId:  params.NcountOrderId,
		TranFinishDate: params.TranFinishDate,
		ServiceAmount:  params.ServiceAmount,
		PayAcctAmount:  params.PayAcctAmount,
	}

	//调用rpc
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, "0000")
	if etcdConn == nil {
		errMsg := "0000" + "getcdv3.GetDefaultConn == nil"
		log.NewError("0000", errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}
	// 创建rpc连接
	client := rpc.NewCloudWalletServiceClient(etcdConn)
	RpcResp, err := client.WithDrawNotify(context.Background(), req)
	if err != nil {
		log.NewError("0000", "WithDrawNotify failed ", err.Error(), req.String())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errCode": 200, "data": RpcResp})
	return
}
