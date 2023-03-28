package cloud_wallet

import (
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/proto/cloud_wallet"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

func (rpc *CloudWalletServer) SendRedPacket(ctx context.Context, in *cloud_wallet.SendRedPacketReq, opts ...grpc.CallOption) (*cloud_wallet.SendRedPacketResp, error) {
	handler := &handlerSendRedPacket{SendRedPacketReq: in}
	handler.SendRedPacket()
	return nil, nil
}

/*
	 string userId = 1; //用户id
  int32 PacketType = 2; //红包类型(1个人红包、2群红包)
  int32 IsLucky = 3; //是否为拼手气红包
  int32 IsExclusive = 4; //是否为专属红包(0不是、1是)
  int32 ExclusiveUserID = 5; //专属红包接收者 和isExclusive
  string PacketTitle = 6; //红包标题
  float Amount = 7; //红包金额 单位：分
  int32 Number = 8; //红包个数

  // 通过哪种方式发送红包
  int32 SendType = 9; //发送方式(1钱包余额、2银行卡)
  int64 BankCardID = 10 ;//银行卡id
  string operationID = 11; //链路跟踪id
}
*/

type handlerSendRedPacket struct {
	*cloud_wallet.SendRedPacketReq
}

func (req *handlerSendRedPacket) SendRedPacket() (*cloud_wallet.SendRedPacketResp, error) {
	// 1. 校验参数
	if err := req.validateParam(); err != nil {
		return nil, err
	}

	// 首先生成红包ID 生成规则：红包类型+用户ID+时间戳+随机数
	// 2. 生成红包ID后，发送红包，记录红包发送记录
	if err := req.recordRedPacket(); err != nil {
		return nil, err
	}

	// 3. 判断红包发送方式， 是钱包转账 还是银行卡转账
	if req.SendType == 1 {
		// 钱包转账
		if err := req.walletTransfer(); err != nil {
			return nil, err
		}
	} else {
		// 银行卡转账
		if err := req.bankTransfer(); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (req *handlerSendRedPacket) validateParam() error {
	if req.UserId == 0 {
		return errors.New("user_id is empty")
	}

	// 检测红包类型
	if req.PacketType != 1 && req.PacketType != 2 {
		return errors.New("red_packet_type is bad input ")
	}

	// 检测是否为幸运红包
	if req.IsLucky != 1 && req.IsLucky != 0 {
		return errors.New(fmt.Sprintf("is_lucky is bad input ", req.IsLucky))
	}

	// 检测是否为专属红包
	if req.IsExclusive != 1 && req.IsExclusive != 0 {
		return errors.New(fmt.Sprintf("is_exclusive is bad input ", req.IsExclusive))
	}

	// 专属红包必须要专属用户id
	if req.IsExclusive == 1 && req.ExclusiveUserID == 0 {
		return errors.New("exclusive red packet must be exclusive user id")
	}
	// 红包必须要标题
	if req.PacketTitle == "" {
		return errors.New("red_packet_title is empty")
	}

	// 红包金额必须大于0
	if req.Amount <= 0 {
		return errors.New(fmt.Sprintf("red_packet_amount is bad input , %v", req.Amount/100))
	}

	// 红包个数必须大于0
	if req.Number <= 0 {
		return errors.New("red_packet_number is bad input")
	}

	if req.IsExclusive == 1 && req.PacketType != 2 {
		return errors.New("exclusive red packet must be group red packet")
	}

	// 检测发送方式
	if req.SendType != 1 && req.SendType != 2 {
		return errors.New("send_type is bad input ")
	}

	if req.SendType == 2 && req.BankCardID == 0 {
		return errors.New("bank_card_id is empty")
	}

	return nil
}

func (req *handlerSendRedPacket) validateMore() error {
	// 1. 验证用户是否在群内部
	// 2. 验证用户之间是否是好友关系
	return nil
}

// 保存红包记录
func (in *handlerSendRedPacket) recordRedPacket() error {
	rand.Seed(time.Now().UnixNano())
	redID := fmt.Sprintf("%d%d%d%d", in.PacketType, in.UserId, time.Now().Unix(), rand.Intn(100000))
	redPacket := &imdb.FPacket{
		PacketID:        redID,
		UserID:          in.UserId,
		PacketType:      in.PacketType,
		IsLucky:         in.IsLucky,
		ExclusiveUserID: int64(in.ExclusiveUserID),
		PacketTitle:     in.PacketTitle,
		Amount:          in.Amount,
		Number:          in.Number,
		ExpireTime:      time.Now().Unix() + 60*60*24,
		CreatedTime:     time.Now().Unix(),
		UpdatedTime:     time.Now().Unix(),
		Status:          0, // 红包被创建，但是还未掉第三方的内容
		IsExclusive:     in.IsExclusive,
	}
	return imdb.RedPacketCreateData(redPacket)
}

// 钱包转账
func (in *handlerSendRedPacket) walletTransfer() error {
	// 1. 构造转账接口参数
	// 2. 调用转账接口
	// 3. 开启事务 ： 1. 更新红包的状态 2. 更新用户的状态

	return nil
}

// 银行卡转账
func (in *handlerSendRedPacket) bankTransfer() error {

	// 1. 构造转账接口参数
	// 2. 调用转账接口
	// 3. 开启事务 ： 1. 更新红包的状态 2. 更新用户的状态
	return nil
}
