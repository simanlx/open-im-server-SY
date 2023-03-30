package account

import (
	"github.com/gin-gonic/gin"
)

func SendRedPacket(c *gin.Context) {

	/*
			var (
				req   api.GetAllConversationsReq
				resp  api.GetAllConversationsResp
				reqPb pbUser.GetAllConversationsReq
			)
			if err := c.BindJSON(&req); err != nil {
				log.NewError(req.OperationID, utils.GetSelfFuncName(), "bind json failed", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "bind json failed " + err.Error()})
				return
			}
			log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "req: ", req)
			if err := utils.CopyStructFields(&reqPb, req); err != nil {
				log.NewDebug(req.OperationID, utils.GetSelfFuncName(), "CopyStructFields failed", err.Error())
			}
			etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImUserName, req.OperationID)
			if etcdConn == nil {
				errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
				log.NewError(req.OperationID, errMsg)
				c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
				return
			}
			client := pbUser.NewUserClient(etcdConn)
			respPb, err := client.GetAllConversations(context.Background(), &reqPb)
			if err != nil {
				log.NewError(req.OperationID, utils.GetSelfFuncName(), "SetConversation rpc failed, ", reqPb.String(), err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": "GetAllConversationMsgOpt rpc failed, " + err.Error()})
				return
			}
			resp.ErrMsg = respPb.CommonResp.ErrMsg
			resp.ErrCode = respPb.CommonResp.ErrCode
			if err := utils.CopyStructFields(&resp.Conversations, respPb.Conversations); err != nil {
				log.NewDebug(req.OperationID, utils.GetSelfFuncName(), "CopyStructFields failed, ", err.Error())
			}
			log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "resp: ", resp)
			c.JSON(http.StatusOK, resp)
		}
	*/

	//var (
	//	req   api.SendRedPacketReq  // 请求参数
	//	resp  api.SendRedPacketResp // 返回参数
	//	reqPb rpc.SendRedPacketReq  // 请求参数
	//)
	//if err := c.BindJSON(&req); err != nil {
	//	log.NewError(req.OperationID, utils.GetSelfFuncName(), "bind json failed", err.Error())
	//	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "bind json failed " + err.Error()})
	//	return
	//}
	//log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "req: ", req)
	//if err := utils.CopyStructFields(&reqPb, req); err != nil {
	//	log.NewDebug(req.OperationID, utils.GetSelfFuncName(), "CopyStructFields failed", err.Error())
	//}
	//etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, req.OperationID)
	//
	//if etcdConn == nil {
	//	errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
	//	log.NewError(req.OperationID, errMsg)
	//	c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
	//	return
	//}
	//
	//client := rpc.NewCloudWalletServiceClient(etcdConn)
	//respPb, err := client.SendRedPacket(context.Background(), &reqPb)
	//if err != nil {
	//	log.NewError(req.OperationID, utils.GetSelfFuncName(), "SendRedPacket rpc failed, ", reqPb.String(), err.Error())
	//	c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": "SendRedPacket rpc failed, " + err.Error()})
	//	return
	//}
	//resp.ErrMsg = respPb.CommonResp.ErrMsg
	//resp.ErrCode = respPb.CommonResp.ErrCode
	//if err := utils.CopyStructFields(&resp, respPb); err != nil {
	//	log.NewDebug(req.OperationID, utils.GetSelfFuncName(), "CopyStructFields failed, ", err.Error())
	//}
	//log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "resp: ", resp)
	//c.JSON(http.StatusOK, resp)
}
