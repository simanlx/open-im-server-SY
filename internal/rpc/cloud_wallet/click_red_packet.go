package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	imdb2 "Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/contrive_msg"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
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

	// 返回消息
	if resp.CommonResp.ErrMsg == "" {
		resp.CommonResp.ErrMsg = pb.CloudWalletErrCode_name[int32(resp.CommonResp.ErrCode)]
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

	// 1. 检测红包是否过期
	// ======================================================== 进行红包状态的校验
	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_ServerError
		log.Error(req.OperationID, "数据库查询失败", err)
		return res, err
	}
	if redPacketInfo.Status == imdb.RedPacketStatusCreate {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsCreate
		res.CommonResp.ErrMsg = "红包状态错误"
		return res, nil
	}

	if redPacketInfo.Status == imdb.RedPacketStatusFinished {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsFinish
		res.CommonResp.ErrMsg = "红包已经退还"
		return res, nil
	}
	if redPacketInfo.Status == imdb.RedPacketStatusExpired {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExpire
		res.CommonResp.ErrMsg = "红包已过期"
		return res, nil
	}

	if redPacketInfo.IsExclusive == 1 && redPacketInfo.ExclusiveUserID != req.UserId {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExclusive
		res.CommonResp.ErrMsg = "红包为专属红包"
		return res, nil
	}

	// 2. 检测红包的领取记录 ，如果已经完成领取就不能再领取 , 针对当前用户查询红包领取记录
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

	// 如果用户没实名认证就不能进行抢红包
	/*	err = h.checkUserAuthStatus(req.UserId)
		if err != nil {
			res.CommonResp.ErrCode = pb.CloudWalletErrCode_UserNotValidate
			res.CommonResp.ErrMsg = "您的帐号没有实名认证"
			return res, nil
		}*/
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
		res.CommonResp.ErrMsg = "获取红包金额失败"
		return res, errors.Wrap(err, "获取红包金额失败")
	}

	// 调用转账
	sendAccount, receiveAccount, err := h.getRedPacketByUser(req.UserId, req.RedPacketID)
	if err != nil {
		fmt.Println("获取转账信息失败", err)
		log.Error(req.OperationID, "获取转账信息失败", err)
		return res, errors.Wrap(err, "获取转账信息失败")
	}

	// 转账
	merOrderID := ncount.GetMerOrderID()
	resp, err := h.transferRedPacketToUser(amount, sendAccount, receiveAccount, merOrderID)
	if err != nil {
		log.Error(req.OperationID, "转账失败", err)
		return res, errors.Wrap(err, "转账失败")
	}

	// 5.更新红包领取记录
	err = h.RecodeRedPacket(req, amount, merOrderID)
	if err != nil {
		log.Error(req.OperationID, "更新红包领取记录失败", err)
		return res, errors.Wrap(err, "更新红包领取记录失败")
	}

	// 6.发送红包领取通知
	if err := h.sendRedPacketMsg(req, amount); err != nil {
		log.Error(req.OperationID, "发送红包领取通知失败", err)
		return res, errors.Wrap(err, "发送红包领取通知失败")
	}

	res.CommonResp = resp
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
	res := &pb.CommonResp{}
	amount, err := db.DB.GetRedPacket(req.RedPacketID)
	if err != nil {
		res.ErrCode = pb.CloudWalletErrCode_ServerError
		return 0, err
	}
	return amount, nil
}

// 发送红包领取消息
func (h *handlerClickRedPacket) sendRedPacketMsg(req *pb.ClickRedPacketReq, amount int) error {
	// 获取红包信息
	pkg, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		return err
	}
	// 获取发送用户信息
	userInfo, err := imdb2.GetUserByUserID(pkg.UserID)
	if err != nil {
		return err
	}
	// 获取到抢红包用户的信息
	recvUserInfo, err := imdb2.GetUserByUserID(req.UserId)
	if err != nil {
		return err
	}
	sendtoSenderMsg := fmt.Sprintf("你领取了%s的红包", userInfo.Nickname)
	sendtoRecvMsg := fmt.Sprintf("%s领取了你的红包", recvUserInfo.Nickname)

	// xxx 领取了你的红包
	// 你领取了xxx的红包
	contrive_msg.SendGrabPacket(pkg.UserID, pkg.RecvID, pkg.SendType, h.OperateID, sendtoSenderMsg, sendtoRecvMsg, pkg.PacketID)
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
func (h *handlerClickRedPacket) transferRedPacketToUser(Amount int, payUserID, ReceiveUserID, merOrder string) (*pb.CommonResp, error) {
	request := &ncount.TransferReq{
		MerOrderId: merOrder,
		TransferMsgCipher: ncount.TransferMsgCipher{
			PayUserId:     payUserID,
			ReceiveUserId: ReceiveUserID,
			TranAmount:    strconv.Itoa(int(Amount)),
		},
	}
	resp, err := h.count.Transfer(request)
	if err != nil {
		log.Error(h.OperateID, "发送红包转账失败", err, resp)
		return nil, errors.Wrap(err, "第三方转账失败")
	}
	var result = &pb.CommonResp{
		ErrCode: 0,
	}
	if resp.ResultCode != ncount.ResultCodeSuccess {
		log.Error(h.OperateID, "发送红包转账失败", err, resp)
		// todo
		result.ErrCode = pb.CloudWalletErrCode_ServerError
		result.ErrMsg = "很遗憾没有抢到红包，再接再厉"
	}
	return result, nil
}

// 更新红包领取记录
func (h *handlerClickRedPacket) RecodeRedPacket(req *pb.ClickRedPacketReq, amount int, merOrderID string) error {
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
	err := imdb.InsertRedPacketDetail(data)
	if err != nil {
		return err
	}
	// 2. 更新红包领取数量
	err = imdb.UpdateRedPacketRemain(req.RedPacketID)
	if err != nil {
		return err
	}
	return nil
}
