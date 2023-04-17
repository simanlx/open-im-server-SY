package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	imdb2 "Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/contrive_msg"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"time"
)

// 抢红包
func (rpc *CloudWalletServer) ClickRedPacket(ctx context.Context, req *pb.ClickRedPacketReq) (*pb.ClickRedPacketResp, error) {
	handler := &handlerClickRedPacket{
		OperateID: req.OperationID,
		count:     rpc.count,
	}

	resp, err := handler.ClickRedPacket(req)
	if err != nil {
		log.Error(req.OperationID, "抢红包失败", err)
		return nil, err
	}

	return resp, nil
}

type handlerClickRedPacket struct {
	OperateID string
	count     ncount.NCounter
}

func (h *handlerClickRedPacket) ClickRedPacket(req *pb.ClickRedPacketReq) (*pb.ClickRedPacketResp, error) {
	var (
		res = &pb.ClickRedPacketResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "领取成功",
			},
		}
	)

	// 如果用户没实名认证就不能进行抢红包
	if err := h.checkUserAuthStatus(req.UserId); err != nil {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_UserNotValidate
		res.CommonResp.ErrMsg = "您的帐号没有实名认证,请尽快去实名认证"
		return res, nil
	}

	// 检测红包领取记录
	fp, err := imdb.FPacketDetailGetByPacketID(req.RedPacketID, req.UserId)
	if err != nil {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_ServerError
		res.CommonResp.ErrMsg = "存在领取记录"
		return res, errors.Wrap(err, "红包领取记录查询失败")
	}
	if fp.ID != 0 {
		// 代表存在领取记录
		res.CommonResp.ErrMsg = "你已经领取过该红包"
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsReceived
		return res, nil
	}

	// 1. 检测红包的状态
	// ======================================================== 进行红包状态的校验========================================================

	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		log.Error(h.OperateID, "获取红包信息失败", err)
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_ServerError
		res.CommonResp.ErrMsg = "网络错误"
		return res, err
	}

	// 3. 群红包是否设置了禁止抢红包
	if redPacketInfo.PacketType == 2 { // 群红包
		groupInfo, err := rocksCache.GetGroupInfoFromCache(redPacketInfo.RecvID)
		if err != nil || groupInfo.GroupID == "" {
			log.Info(req.OperationID, "获取群信息失败", err)
			res.CommonResp.ErrCode = pb.CloudWalletErrCode_ServerError
			res.CommonResp.ErrMsg = "获取群信息失败"
			return res, errors.Wrap(err, "获取群信息失败")
		}

		if groupInfo.BanClickPacket == 1 {
			res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsBan
			res.CommonResp.ErrMsg = "该群禁止抢红包"
			return res, nil
		}
	}

	if redPacketInfo.Status == imdb.RedPacketStatusCreate {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsCreate
		res.CommonResp.ErrMsg = "红包状态错误"
		return res, nil
	}

	if redPacketInfo.Status == imdb.RedPacketStatusFinished {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsFinish
		res.CommonResp.ErrMsg = "红包已经被抢完"
		return res, nil
	}
	if redPacketInfo.Status == imdb.RedPacketStatusExpired || redPacketInfo.ExpireTime < time.Now().Unix() {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExpire
		res.CommonResp.ErrMsg = "红包已过期"
		return res, nil
	}

	if redPacketInfo.IsExclusive == 1 && redPacketInfo.ExclusiveUserID != req.UserId {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExclusive
		res.CommonResp.ErrMsg = "红包为专属红包"
		return res, nil
	}

	if redPacketInfo.Remain <= 0 {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsFinish
		res.CommonResp.ErrMsg = "红包已经被抢完"
		return res, nil
	}

	// ===========================================================红包状态如果是OK的，进行抢红包操作==================================================
	// 5. 根据红包的
	var amount int
	// 4. 判断红包的类型
	if (redPacketInfo.PacketType == 1 && redPacketInfo.IsExclusive != 1) || redPacketInfo.Number == 1 {
		// 直接查询数据库
		amount, err = h.getRedPacketAmount(req.RedPacketID)
	} else {
		// 通过查询redis
		amount, err = h.getRedPacketByGroup(req)
	}
	if err != nil {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_ServerError
		res.CommonResp.ErrMsg = "服务器错误：获取红包金额失败"
		return res, nil
	}

	if amount == 0 {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsFinish
		res.CommonResp.ErrMsg = "红包已经被抢完"
		return res, nil
	}

	// 获取发送用户的账户和接受用户的账户 ： 查询预备数据
	sendAccount, receiveAccount, err := h.getRedPacketByUser(req.UserId, req.RedPacketID)
	if err != nil {
		log.Error(req.OperationID, "网络错误：获取用户转账信息失败", err)
		res.CommonResp.ErrMsg = "网络错误：获取用户转账信息失败,操作ID:" + req.OperationID
		return res, nil
	}
	amount1 := cast.ToString(cast.ToFloat64(amount) / 100)
	merOrderID := ncount.GetMerOrderID()

	// ========================================================第三方转账操作========================================================

	// todo 这里可能出现问题，没有进行补偿操作
	respNcount, err := h.transferRedPacketToUser(amount1, sendAccount, receiveAccount, merOrderID)
	if err != nil {
		log.Error(req.OperationID, "转账失败", err, respNcount)
		return res, errors.Wrap(err, "转账失败")
	}

	if respNcount.ResultCode != "0000" {
		// 所有新生支付的错误都进行暴露，这样处理不太规范，但是前期方便核对错误
		log.Error(req.OperationID, "转账失败", err, respNcount)
		res.CommonResp.ErrMsg = respNcount.ErrorMsg
		return res, nil
	}

	// 5.更新红包领取记录 ，todo 这里可以重复保存 如果保存失败应该需要去重试机制
	islastOne, err := h.RecodeRedPacket(req, amount, merOrderID)
	if err != nil {
		log.Error(req.OperationID, "更新红包领取记录失败", err)
		return res, errors.Wrap(err, "更新红包领取记录失败")
	}

	// ========================================================发送抢红包结果操作========================================================

	// 6.todo 这里可以重复保存 如果保存失败应该需要去重试机制
	if redPacketInfo.PacketType == 1 {
		SendRedPacketMsg(redPacketInfo, req.OperationID, req.UserId)
	} else {
		SendRedPacketMsg(redPacketInfo, req.OperationID)
	}

	// 7. 如果这个红包是最后一个红包 需要更新运气王的信息、修改红包的状态
	if islastOne && redPacketInfo.PacketType == 2 {
		// 修改用户运气王的信息
		err = imdb.UpdateLuckyKing(req.RedPacketID, req.UserId)
		if err != nil {
			log.Error(req.OperationID, "更新红包信息失败", err)
		}
	}

	// ========================================================添加交易记录操作========================================================
	// 8.添加交易记录
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(req.OperationID, "Panic: 添加交易记录失败", err)
			}
		}()
		err = AddNcountTradeLog(BusinessTypeReceivePacket, int32(amount), req.UserId, "", respNcount.NcountOrderId, redPacketInfo.PacketID)
		if err != nil {
			log.Error(req.OperationID, "添加交易记录失败", err)
		}
	}()

	res.CommonResp.ErrCode = 0
	res.CommonResp.ErrMsg = "领取成功"
	return res, nil
}

// 检查用户实名认证状态
func (h *handlerClickRedPacket) checkUserAuthStatus(userID string) error {
	// 判断用户是否实名注册
	_, err := imdb.GetNcountAccountByUserId(userID)
	if err != nil {
		return errors.Wrap(err, "查询用户实名认证状态失败")
	}
	return nil
}

// 如果红包是群聊红包， 直接redis的set集合进行获取红包
func (h *handlerClickRedPacket) getRedPacketByGroup(req *pb.ClickRedPacketReq) (int, error) {
	amount, err := db.DB.GetRedPacket(req.RedPacketID)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

// 发送红包领取消息
func SendRedPacketMsg(redpacketInfo *imdb.FPacket, operationID string, clickUserID ...string) error {

	// 这里判断群红包是单人红包还是群聊红包 ： 单人红包需要推送两条领取消息，群聊红包需要推送一条领取消息
	if redpacketInfo.PacketType == 1 {
		// 获取发送用户信息
		userInfo, err := imdb2.GetUserByUserID(redpacketInfo.UserID)
		if err != nil {
			return err
		}
		// 获取到抢红包用户的信息
		recvUserInfo, err := imdb2.GetUserByUserID(redpacketInfo.RecvID)
		if err != nil {
			return err
		}
		// 发送给发送者的消息
		err = contrive_msg.RedPacketGrabPushToUser(operationID, redpacketInfo.UserID, redpacketInfo.UserID, redpacketInfo.PacketID, userInfo.Nickname, recvUserInfo.Nickname, redpacketInfo.RecvID)
		// 发送给接受者的消息
		if err != nil {
			return err
		}
		err = contrive_msg.RedPacketGrabPushToUser(operationID, redpacketInfo.RecvID, redpacketInfo.UserID, redpacketInfo.PacketID, userInfo.Nickname, recvUserInfo.Nickname, redpacketInfo.UserID)
		if err != nil {
			return err
		}
	} else {
		if len(clickUserID) == 0 {
			return errors.New("群聊红包需要传入点击用户ID")
		}
		clickID := clickUserID[0]
		recvUserInfo, err := imdb2.GetUserByUserID(clickID)
		if err != nil {
			return err
		}
		userInfo, err := imdb2.GetUserByUserID(redpacketInfo.UserID)
		if err != nil {
			return err
		}
		// 这里是群聊红包逻辑
		err = contrive_msg.RedPacketGrabPushToGroup(operationID, redpacketInfo.UserID, recvUserInfo.UserID, redpacketInfo.PacketID, userInfo.Nickname, recvUserInfo.Nickname, redpacketInfo.RecvID)
		if err != nil {
			return err
		}
	}
	return nil
}

// 从红包记录获取转账金额
func (h *handlerClickRedPacket) getRedPacketAmount(redID string) (int, error) {
	redPacketInfo, err := imdb.GetRedPacketInfo(redID)
	if err != nil {
		return 0, err
	}
	return int(redPacketInfo.Amount), nil
}

// 通过红包ID 倒查红包发送者的红包账户
func (h *handlerClickRedPacket) getRedPacketByUser(GrapRedPacketUserID, packetID string) (string, string, error) {
	//  获取到发红包的用户ID
	senderAccount, err := imdb.SelectRedPacketSenderRedPacketAccountByPacketID(packetID)
	if err != nil {
		return "", "", errors.Wrap(err, "查询红包发送用户失败")
	}
	recieveAccount, err := imdb.SelectUserMainAccountByUserID(GrapRedPacketUserID)
	if err != nil {
		return "", "", err
	}

	return senderAccount, recieveAccount, nil
}

// 发起转账 红包账户对具体用户进行转账，并调用红包消息发送
func (h *handlerClickRedPacket) transferRedPacketToUser(Amount string, payUserID, ReceiveUserID, merOrder string) (*ncount.TransferResp, error) {
	request := &ncount.TransferReq{
		MerOrderId: merOrder,
		TransferMsgCipher: ncount.TransferMsgCipher{
			PayUserId:     payUserID,
			ReceiveUserId: ReceiveUserID,
			TranAmount:    Amount,
		},
	}
	return h.count.Transfer(request)

}

// 更新红包领取记录
func (h *handlerClickRedPacket) RecodeRedPacket(req *pb.ClickRedPacketReq, amount int, merOrderID string) (bool, error) {
	// 保存红包领取记录
	data := &imdb.FPacketDetail{
		PacketID:    req.RedPacketID,
		UserID:      req.UserId,
		MerOrderID:  merOrderID,
		Amount:      int64(amount),
		ReceiveTime: time.Now().Unix(),
		CreatedTime: time.Now().Unix(),
		UpdatedTime: time.Now().Unix(),
	}
	// 1. 更新红包领取记录
	return imdb.SaveRedPacketDetail(data)
}
