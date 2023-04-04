package cloud_wallet

import (
	ncount "Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/db"
	commonDB "Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	pb "Open_IM/pkg/proto/cloud_wallet"
	pb2 "Open_IM/pkg/proto/push"
	"Open_IM/pkg/tools/redpacket"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// 发送红包接口
func (rpc *CloudWalletServer) SendRedPacket(ctx context.Context, in *pb.SendRedPacketReq) (*pb.SendRedPacketResp, error) {
	handler := &handlerSendRedPacket{
		OperateID: in.GetOperationID(),
		count:     rpc.count,
	}
	return handler.SendRedPacket(in)
}

type handlerSendRedPacket struct {
	OperateID string
	count     ncount.NCounter
}

// 钱包账户转账
func (h *handlerSendRedPacket) SendRedPacket(req *pb.SendRedPacketReq) (*pb.SendRedPacketResp, error) {
	// 1. 校验参数
	if err := h.validateParam(req); err != nil {
		return nil, err
	}
	// 首先生成红包ID 生成规则：红包类型+用户ID+时间戳+随机数
	// 2. 生成红包ID后，发送红包，记录红包发送记录
	redpacketID, err := h.recordRedPacket(req)
	if err != nil {
		log.Error(req.OperationID, "record red packet error", zap.Error(err))
		return nil, err
	}
	res := &pb.SendRedPacketResp{
		RedPacketID: redpacketID,
	}

	// 钱包转账,是同步的
	commonResp, err := h.walletTransfer(redpacketID, req)
	if err != nil {
		log.Error(req.OperationID, "转账失败", err)
		return nil, err
	}
	if commonResp.ErrCode != 0 {
		log.Error(req.OperationID, "wallet transfer error", zap.String("err_msg", commonResp.ErrMsg))
		return nil, errors.New(commonResp.ErrMsg)
	}
	// 处理回调内容
	err = HandleSendPacketResult(redpacketID, req.OperationID)
	if err != nil {
		log.Error(req.OperationID, "HandleSendPacketResult error", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (h *handlerSendRedPacket) validateParam(req *pb.SendRedPacketReq) error {
	if len(req.UserId) <= 0 {
		return errors.New("user_id is empty")
	}

	// 检测红包类型
	if req.PacketType != 1 && req.PacketType != 2 {
		return errors.New("red_packet_type is bad input ")
	}

	// 检测是否为幸运红包
	if req.IsLucky != 1 && req.IsLucky != 0 {
		return errors.New("is_lucky is bad input ")
	}

	// 检测是否为专属红包
	if req.IsExclusive != 1 && req.IsExclusive != 0 {
		return errors.New("is_exclusive is bad input ")
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
	// 1 检测上传的银行卡ID是否为用户自己的
	// 1. 验证用户是否在群内部
	// 2. 验证用户之间是否是好友关系
	return nil
}

// 保存红包记录
func (h *handlerSendRedPacket) recordRedPacket(in *pb.SendRedPacketReq) (string /* red packet ID */, error) {
	rand.Seed(time.Now().UnixNano())
	redID := fmt.Sprintf("%v%v%v%v", in.PacketType, in.UserId, time.Now().Unix(), rand.Intn(100000))
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
	return redID, imdb.RedPacketCreateData(redPacket)
}

// 钱包转账
func (h *handlerSendRedPacket) walletTransfer(redPacketID string, in *pb.SendRedPacketReq) (*pb.CommonResp, error) {
	// 1. 获取用户的钱包账户
	fncount, err := imdb.FNcountAccountGetUserAccountID(in.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "get user FNcountAccountGetUserAccountID by id error")
	}
	req := &ncount.TransferReq{
		MerOrderId: ncount.GetMerOrderID(),
		TransferMsgCipher: ncount.TransferMsgCipher{
			PayUserId:     fncount.MainAccountId,
			ReceiveUserId: fncount.PacketAccountId,
			TranAmount:    strconv.Itoa(int(in.Amount)),
		},
	}

	log.Info(in.OperationID, "transfer req", req)
	transferResult, err := h.count.Transfer(req)
	log.Info(in.OperationID, "transfer res", transferResult)
	if err != nil {
		return nil, errors.Wrap(err, "调用新生支付出现错误")
	}

	commonResp := &pb.CommonResp{
		ErrMsg:  "发送成功",
		ErrCode: 0,
	}
	fmt.Println(transferResult)
	if transferResult.ResultCode != ncount.ResultCodeSuccess {
		// 如果转账失败，需要给用户提示发送失败，并将红包状态修改为发送失败
		err := imdb.UpdateRedPacketStatus(redPacketID, 2 /* 发送失败 */)
		if err != nil {
			log.Error(in.OperationID, zap.Error(err))
			return nil, errors.Wrap(err, "修改红包状态失败 2 ")
			// todo 记录到死信队列中，后续监控处理 ，不需要人工介入
		}
		// 记录操作失败日志 ，提供后续给客服人员核对
		commonResp.ErrCode = 1
		commonResp.ErrMsg = transferResult.ErrorMsg
		return commonResp, nil
	}
	// 如果转账成功，需要将红包状态修改为发送成功
	err = imdb.UpdateRedPacketStatus(redPacketID, 1 /* 发送成功 */)
	if err != nil {
		// todo 记录到死信队列中，后续监控处理， 如果转账成功，但是修改红包状态失败，需要人工介入
		log.Error(in.OperationID, zap.Error(err))
		return nil, errors.Wrap(err, "修改红包状态失败 1")
	}

	payAcctAmount, err := strconv.Atoi(transferResult.PayAcctAmount)
	if err != nil {
		log.Error(in.OperationID, zap.Error(err))
		payAcctAmount = 0
	}

	// 记录用户的消费记录
	err = imdb.FNcountTradeCreateData(&db.FNcountTrade{
		UserID:          in.UserId,
		PaymentPlatform: 1,                          // 云钱包
		Type:            imdb.TradeTypeRedPacketOut, // 红包转出
		Amount:          in.Amount,
		BeferAmount:     int64(payAcctAmount) - in.Amount, // 转账前的金额
		AfterAmount:     int64(payAcctAmount),             // 转账后的金额
		ThirdOrderNo:    transferResult.MerOrderId,        // 第三方的订单号
	})
	if err != nil {
		// todo 记录到死信队列中，后续监控处理， 如果转账成功，但是记录用户的消费记录失败，需要人工介入
		log.Error(in.OperationID, zap.Error(err))
		return nil, errors.Wrap(err, "记录用户交易表失败")
	}
	// todo 需要发送红包发送成功给用户IM
	return commonResp, nil
}

// 当用户发布红包发送成功的时候，调用这个回调函数进行发布红包的后续处理
func HandleSendPacketResult(redPacketID, OperateID string) error {
	//1. 查询红包信息
	redpacketInfo, err := imdb.GetRedPacketInfo(redPacketID)
	if err != nil {
		log.Error(OperateID, "get red packet info error", zap.Error(err))
		return err
	}
	if redpacketInfo == nil {
		log.Error(OperateID, "red packet info is nil")
		return errors.New("red packet info is nil")
	}
	// 2. 生成红包
	if redpacketInfo.PacketType == 2 {
		// 群红包
		err = GroupPacket(redpacketInfo, OperateID)
	}

	// 3. 修改红包状态
	err = imdb.UpdateRedPacketStatus(redPacketID, imdb.RedPacketStatusNormal)
	if err != nil {
		log.Error(OperateID, "update red packet status error", zap.Error(err))
		return err
	}

	// 发送红包消息给用户 和发送红包的人
	_, err = sendRedMessage(redpacketInfo, OperateID)
	if err != nil {
		log.Error(OperateID, "send red packet message error", zap.Error(err))
		return err
	}
	return nil
}

// 给群发的红包
func GroupPacket(req *imdb.FPacket, redpacketID string) error {
	var err error
	if req.IsLucky == 1 {
		// 如果说是手气红包，分散放入红包池
		err = spareRedPacket(req.OperateID, redpacketID, int(req.Amount), int(req.Number))
	} else {
		// 非手气红包，平均分配
		err = spareEqualRedPacket(req.OperateID, redpacketID, int(req.Amount), int(req.Number))
	}
	if err != nil {
		log.Error(req.OperateID, zap.Error(err))
		return err
	}

	return err
}

// 将红包放入红包池
func spareRedPacket(OperateID, packetID string, amount, number int) error {
	// 将发送的红包进行计算
	result := redpacket.GetRedPacket(amount, number)
	err := commonDB.DB.SetRedPacket(packetID, result...)
	if err != nil {
		log.Error(OperateID, zap.Error(err))
		return err
	}
	return nil
}

// amount = 3 ,number =3
func spareEqualRedPacket(OperateID, packetID string, amount, number int) error {
	result := []int{}
	for i := 0; i < number; i++ {
		result = append(result, amount)
	}
	// 将发送的红包进行计算
	err := commonDB.DB.SetRedPacket(packetID, result...)
	if err != nil {
		log.Error(OperateID, zap.Error(err))
		return err
	}
	return nil
}

type paramsUserSendMsg struct {
	OperationID string `json:"operationID"`
	SendID      string `json:"sendID"`
	RecvID      string `json:"recvID"`
	GroupID     string `json:"groupID"`
	Content     struct {
		Text string `json:"text"`
	} `json:"content"`
	ContentType     int32 `json:"contentType"` // 110 红包消息
	SessionType     int32 `json:"sessionType"`
	IsOnlineOnly    bool  `json:"isOnlineOnly"`
	OfflinePushInfo Offline
}

type Offline struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Ex            string `json:"ex"`
	IOSPushSound  string `json:"iOSPushSound"`
	IOSBadgeCount bool   `json:"iOSBadgeCount"`
}

// 发送红包消息给用户
func sendRedMessage(req *imdb.FPacket, redpacketID string) (*pb2.PushMsgResp, error) {
	// 创建红包消息
	p := NewPostMessage(req)
	// http post 发送红包消息
	url := "localhost:8080/v4/openim/sendmsg"
	data, err := json.Marshal(p)
	if err != nil {
		log.Error(req.OperateID, zap.Error(err))
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Error(req.OperateID, zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(req.OperateID, zap.Error(err))
		return nil, err
	}
	log.Info(req.OperateID, zap.String("send red packet message", string(body)))
	return nil, nil
}

// 发送红包消息
func NewPostMessage(f *imdb.FPacket) *paramsUserSendMsg {
	p := &paramsUserSendMsg{
		OfflinePushInfo: Offline{},
	}
	p.OperationID = f.OperateID
	p.SendID = f.UserID
	if f.PacketType == 1 {
		// 个人红包
		p.RecvID = f.RecvID
		p.SessionType = 1 // 单聊
	} else {
		// 群红包
		p.RecvID = ""
		p.GroupID = f.RecvID
		p.SessionType = 2 // 群消息
	}
	// 离线消息
	p.OfflinePushInfo.Title = "你有新的红包"
	p.OfflinePushInfo.Desc = ""
	p.OfflinePushInfo.Ex = ""
	p.OfflinePushInfo.IOSPushSound = "default"
	return p
}
