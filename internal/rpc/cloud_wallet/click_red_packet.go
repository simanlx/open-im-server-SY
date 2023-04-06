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
)

// 抢红包
func (rpc *CloudWalletServer) ClickRedPacket(ctx context.Context, req *pb.ClickRedPacketReq) (*pb.ClickRedPacketResp, error) {
	// 参数校验 ：红包是否过期、 红包状态判断、红包的类型
	// 判断抢红包是否实名认证过，如果没实名认证过就不能抢红包
	// 如果是群聊红包 ： 请求进行管理
	// 如果是个人红包： 调用转账接口 ，
	// 生成红包记录日志
	// 记录用户收入日志
	// 修改红包的余额

	// 检查用户是否实名认证

	handler := handlerClickRedPacket{
		ClickRedPacketReq: req,
		OperateID:         req.OperationID,
		count:             rpc.count,
	}
	return handler.ClickRedPacket()
}

type handlerClickRedPacket struct {
	*pb.ClickRedPacketReq
	OperateID string
	count     ncount.NCounter
}

func (req *handlerClickRedPacket) ClickRedPacket() (*pb.ClickRedPacketResp, error) {
	var (
		res    = &pb.ClickRedPacketResp{}
		common = &pb.CommonResp{}
	)
	// 1. 检测红包是否过期
	// 校验红包的状态
	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		log.Error(req.OperateID, "数据库查询失败", err)
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
		amount, err = req.getRedPacketByGroup()
	} else {
		// 直接查询数据库
		amount, err = req.getRedPacketAmount()
	}
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		return res, errors.Wrap(err, "获取红包金额失败")
	}

	// 调用转账
	sendAccount, receiveAccount, err := req.getRedPacketByUser(req.UserId, req.RedPacketID)

	if err != nil {
		log.Error(req.OperateID, "获取转账信息失败", err)
		return res, errors.Wrap(err, "获取转账信息失败")
	}

	resp, err := req.transferRedPacketToUser(amount, sendAccount, receiveAccount)
	if err != nil {
		log.Error(req.OperateID, "转账失败", err)
		return res, errors.Wrap(err, "转账失败")
	}
	res.CommonResp = resp
	return res, nil
}

// 检查用户实名认证状态
func (req *handlerClickRedPacket) checkUserAuthStatus() (*pb.CommonResp, error) {
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
func (req *handlerClickRedPacket) getRedPacketByGroup() (int, error) {
	res := &pb.CommonResp{}
	amount, err := db.DB.GetRedPacket(req.RedPacketID)
	if err != nil {
		res.ErrCode = pb.CloudWalletErrCode_ServerError
		return 0, err
	}
	return amount, nil
}

// 从红包记录获取转账金额
func (req *handlerClickRedPacket) getRedPacketAmount() (int, error) {
	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		return 0, err
	}
	return int(redPacketInfo.Amount), nil
}

// 通过红包ID 倒查红包发送者的红包账户
func (req *handlerClickRedPacket) getRedPacketByUser(GrapRedPacketUserID, packetID string) (string, string, error) {
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
func (req *handlerClickRedPacket) transferRedPacketToUser(Amount int, payUserID, ReceiveUserID string) (*pb.CommonResp, error) {
	MEROrderID := ncount.GetMerOrderID()
	request := &ncount.TransferReq{
		MerOrderId: MEROrderID,
		TransferMsgCipher: ncount.TransferMsgCipher{
			PayUserId:     payUserID,
			ReceiveUserId: ReceiveUserID,
			TranAmount:    strconv.Itoa(int(Amount)),
		},
	}
	resp, err := req.count.Transfer(request)
	if err != nil {
		log.Error(req.OperateID, "发送红包转账失败", err, resp)
		return nil, errors.Wrap(err, "第三方转账失败")
	}
	var result = &pb.CommonResp{
		ErrCode: 0,
	}
	if resp.ResultCode != ncount.ResultCodeSuccess {
		log.Error(req.OperateID, "发送红包转账失败", err, resp)
		// todo
		result.ErrCode = pb.CloudWalletErrCode_ServerError
		result.ErrMsg = "很遗憾没有抢到红包，再接再厉"
	}
	return result, nil
}
