package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
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
	return handler.ClickRedPacket(req)
}

type handlerClickRedPacket struct {
	OperateID string
	count     ncount.NCounter
}

func (h *handlerClickRedPacket) ClickRedPacket(req *pb.ClickRedPacketReq) (*pb.ClickRedPacketResp, error) {
	var (
		res    = &pb.ClickRedPacketResp{}
		common = &pb.CommonResp{}
	)
	// 1. 检测红包是否过期
	// ======================================================== 进行红包状态的校验
	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		log.Error(req.OperationID, "数据库查询失败", err)
		return res, err
	}
	if redPacketInfo.Status == imdb.RedPacketStatusCreate {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsCreate
		return res, nil
	}
	if redPacketInfo.Status == imdb.RedPacketStatusFinished {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsFinish
		return res, nil
	}
	if redPacketInfo.Status == imdb.RedPacketStatusExpired {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExpire
		return res, nil
	}

	if redPacketInfo.IsExclusive == 1 && redPacketInfo.ExclusiveUserID != req.UserId {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExclusive
		return res, nil
	}

	// 2. 检测红包的领取记录 ，如果已经完成领取就不能再领取 , 针对当前用户查询红包领取记录
	fp, err := imdb.FPacketDetailGetByPacketID(req.RedPacketID, req.UserId)
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		return res, errors.Wrap(err, "红包领取记录查询失败")
	}

	if fp.ID != 0 {
		// 代表存在领取记录
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsReceived
		return res, nil
	}
	// 如果用户没实名认证就不能进行抢红包
	//todo

	var amount int
	// 4. 判断红包的类型
	if redPacketInfo.PacketType == 1 && redPacketInfo.IsExclusive != 1 {
		// 直接从redis的set集合进行获取红包
		amount, err = h.getRedPacketAmount(req)
	} else {
		// 直接查询数据库
		amount, err = h.getRedPacketByGroup(req)
	}
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		return res, errors.Wrap(err, "获取红包金额失败")
	}

	// 调用转账
	sendAccount, receiveAccount, err := h.getRedPacketByUser(req.UserId, req.RedPacketID)

	if err != nil {
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

	res.CommonResp = resp
	return res, nil
}

// 检查用户实名认证状态
func (h *handlerClickRedPacket) checkUserAuthStatus() (*pb.CommonResp, error) {
	/*	res := &pb.CommonResp{}
		// 检查用户是否实名认证
		// 检查用户是否实名认证
		authStatus, err := req.count.GetUserAuthStatus(req.UserId)
		if err != nil {
			res.ErrCode = pb.CloudWalletErrCode_ServerError
			res.ErrMsg = "服务器错误"
			return res, errors.Wrap(err, "获取用户实名认证状态失败")
		}
		if authStatus != 1 {
			res.ErrCode = pb.CloudWalletErrCode_UserNotAuth
			res.ErrMsg = "用户未实名认证"
			return res, nil
		}
		return res, nil*/
	return nil, nil
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

// 从红包记录获取转账金额
func (h *handlerClickRedPacket) getRedPacketAmount(req *pb.ClickRedPacketReq) (int, error) {
	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
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
		return "", "", err
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
