package cloud_wallet

import (
	ncount "Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/config"
	commonDB "Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	imdb2 "Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/contrive_msg"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/tools/redpacket"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

// 发送红包接口
func (rpc *CloudWalletServer) SendRedPacket(ctx context.Context, in *pb.SendRedPacketReq) (*pb.SendRedPacketResp, error) {
	handler := &handlerSendRedPacket{
		OperateID:  in.GetOperationID(),
		merOrderID: ncount.GetMerOrderID(),
		count:      rpc.count,
	}
	return handler.SendRedPacket(in)
}

type handlerSendRedPacket struct {
	OperateID  string
	merOrderID string
	count      ncount.NCounter
}

// 钱包账户转账
func (h *handlerSendRedPacket) SendRedPacket(req *pb.SendRedPacketReq) (*pb.SendRedPacketResp, error) {
	// 1. 校验参数
	if err := h.validateParam(req); err != nil {
		return nil, err
	}
	// 首先生成红包ID 生成规则：红包类型+用户ID+时间戳+随机数
	// 2. 生成红包ID后，发送红包，记录红包发送记录

	// todo 暂时这么处理
	if checkGroupValidate := h.checkGroupPacketState(req); checkGroupValidate != "" {
		return nil, errors.New(checkGroupValidate)
	}

	redpacketID, err := h.recordRedPacket(req)
	if err != nil {
		log.Error(req.OperationID, "record red packet error", zap.Error(err))
		return nil, err
	}
	res := &pb.SendRedPacketResp{
		RedPacketID: redpacketID,
	}

	// 3. 判断支付类型
	if req.SendType == 1 {
		// 钱包转账,是同步的
		commonResp, err := h.walletTransfer(redpacketID, req)
		if err != nil {
			log.Error(req.OperationID, "转账失败", err)
			return nil, err
		}
		if commonResp.ErrCode != 0 {
			log.Error(req.OperationID, "钱包转账失败", zap.String("err_msg", commonResp.ErrMsg))
			return nil, errors.New(commonResp.ErrMsg)
		}
		// 记录转账信息

		// 回调处理红包
		err = HandleSendPacketResult(redpacketID, req.OperationID)
		if err != nil {
			log.Error(req.OperationID, "HandleSendPacketResult error", zap.Error(err))
			return nil, err
		}
	} else {
		// 这里是调用银行卡转账接口
		if err != nil {
			log.Error(req.OperationID, "BankCardRechargePacketAccount error", zap.Error(err))
			return nil, err
		}
	}
	return res, nil
}

func (h *handlerSendRedPacket) validateParam(req *pb.SendRedPacketReq) error {
	if len(req.UserId) <= 0 {
		return errors.New("user_id 不能为空")
	}

	// 检测红包类型
	if req.PacketType != 1 && req.PacketType != 2 {
		return errors.New("red_packet_type 错误输入 ")
	}

	// 检测是否为幸运红包
	if req.IsLucky != 1 && req.IsLucky != 0 {
		return errors.New("is_lucky 错误输入 ")
	}

	// 检测是否为专属红包
	if req.IsExclusive != 1 && req.IsExclusive != 0 {
		return errors.New("is_exclusive 错误输入 ")
	}

	// 专属红包必须要专属用户id
	if req.IsExclusive == 1 && req.ExclusiveUserID == "" {
		return errors.New("是专属红包就必须存在ExclusiveUserID")
	}
	// 红包必须要标题
	if req.PacketTitle == "" {
		return errors.New("red_packet_title 红包title不能为空")
	}

	// 红包金额必须大于0
	if req.Amount <= 0 {
		return errors.New(fmt.Sprintf("red_packet_amount 红包金额必须为大于0 , %v", req.Amount/100))
	}

	// 红包个数必须大于0
	if req.Number <= 0 {
		return errors.New("red_packet_number 个数必须大于0")
	}

	if req.IsExclusive == 1 && req.PacketType != 2 {
		return errors.New("IsExclusive 属性红包必须是PacketType = 2 ")
	}

	// 检测发送方式
	if req.SendType != 1 && req.SendType != 2 {
		return errors.New("send_type 发送方式输入错误 ")
	}

	if req.SendType == 2 && req.BankCardID == 0 {
		return errors.New("SendType = 2 时，BankCardID && BindCardAgrNo 不能为空	")
	}

	if req.RecvID == "" {
		return errors.New("RecvID 不能为空")
	}

	return nil
}

func (req *handlerSendRedPacket) validateMore() error {
	// 1 检测上传的银行卡ID是否为用户自己的
	// 1. 验证用户是否在群内部
	// 2. 验证用户之间是否是好友关系

	// 检测用户是否在群里
	return nil
}

// 验证业务上的逻辑错误
func (h *handlerSendRedPacket) checkGroupPacketState(req *pb.SendRedPacketReq) string {
	// 1.用户是否在群里
	if req.PacketType == 2 {
		ok := imdb2.IsExistGroupMember(req.RecvID, req.UserId)
		if !ok {
			return "用户不在群里"
		}
	}
	return ""
}

// 创建红包信息
func (h *handlerSendRedPacket) recordRedPacket(in *pb.SendRedPacketReq) (string /* red packet ID */, error) {
	rand.Seed(time.Now().UnixNano())
	redID := fmt.Sprintf("%v%v%v%v", in.PacketType, in.UserId, time.Now().Unix(), rand.Intn(100000))
	redPacket := &imdb.FPacket{
		PacketID:        redID,
		UserID:          in.UserId,
		PacketType:      in.PacketType,
		IsLucky:         in.IsLucky,
		ExclusiveUserID: in.ExclusiveUserID,
		PacketTitle:     in.PacketTitle,
		Amount:          in.Amount,
		Number:          in.Number,
		MerOrderID:      h.merOrderID,
		OperateID:       h.OperateID,
		SendType:        in.SendType,
		BindCardAgrNo:   in.BindCardAgrNo,
		RecvID:          in.RecvID, // 接收ID
		Remain:          int64(in.Number),
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
		return nil, errors.Wrap(err, "用户没有实名注册")
	}
	req := &ncount.TransferReq{
		MerOrderId: h.merOrderID,
		TransferMsgCipher: ncount.TransferMsgCipher{
			PayUserId:     fncount.MainAccountId,
			ReceiveUserId: fncount.PacketAccountId,
			TranAmount:    cast.ToString(cast.ToFloat64(in.Amount) / 100), //分转元
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

	//增加账户变更日志
	err = AddNcountTradeLog(BusinessTypeBalanceSendPacket, int32(in.Amount), in.UserId, fncount.MainAccountId, transferResult.MerOrderId, redPacketID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("增加账户变更日志失败(%s)", err.Error()))
	}

	return commonResp, nil
}

// 银行卡转账
func (h *handlerSendRedPacket) bankTransfer(redPacketID string, in *pb.SendRedPacketReq) (*pb.CommonResp, error) {
	//银行卡充值到红包账户
	err := BankCardRechargePacketAccount(in.UserId, in.BindCardAgrNo, int32(in.Amount), redPacketID)
	if err != nil {
		return nil, err
	}

	commonResp := &pb.CommonResp{
		ErrMsg:  "发送成功",
		ErrCode: 0,
	}

	// 如果转账成功，需要将红包状态修改为发送成功
	err = imdb.UpdateRedPacketStatus(redPacketID, 1 /* 发送成功 */)
	if err != nil {
		// todo 记录到死信队列中，后续监控处理， 如果转账成功，但是修改红包状态失败，需要人工介入
		log.Error(in.OperationID, zap.Error(err))
		return nil, errors.Wrap(err, "修改红包状态失败 1")
	}
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

	// todo 发送红包消息
	freq := &contrive_msg.FPacket{
		PacketID:        redPacketID,
		UserID:          redpacketInfo.UserID,
		PacketType:      redpacketInfo.PacketType,
		IsLucky:         redpacketInfo.IsLucky,
		ExclusiveUserID: redpacketInfo.ExclusiveUserID,
		PacketTitle:     redpacketInfo.PacketTitle,
		Amount:          redpacketInfo.Amount,
		Number:          redpacketInfo.Number,
		ExpireTime:      redpacketInfo.ExpireTime,
		MerOrderID:      redpacketInfo.MerOrderID,
		OperateID:       redpacketInfo.OperateID,
		RecvID:          redpacketInfo.RecvID,
		CreatedTime:     redpacketInfo.CreatedTime,
		UpdatedTime:     redpacketInfo.UpdatedTime,
		IsExclusive:     redpacketInfo.IsExclusive,
	}
	contrive_msg.SendSendRedPacket(freq, int(redpacketInfo.PacketType))
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

// 银行卡充值到红包账户
func BankCardRechargePacketAccount(userId, bindCardAgrNo string, amount int32, packetID string) error {
	//获取用户账户信息
	accountInfo, err := imdb.GetNcountAccountByUserId(userId)
	if err != nil || accountInfo.Id <= 0 {
		return errors.New("账户信息不存在")
	}

	//充值支付
	accountResp, err := ncount.NewCounter().QuickPayOrder(&ncount.QuickPayOrderReq{
		MerOrderId: ncount.GetMerOrderID(),
		QuickPayMsgCipher: ncount.QuickPayMsgCipher{
			PayType:       "3", //绑卡协议号充值
			TranAmount:    cast.ToString(cast.ToFloat64(amount) / 100),
			NotifyUrl:     config.Config.Ncount.Notify.RechargeNotifyUrl,
			BindCardAgrNo: bindCardAgrNo,
			ReceiveUserId: accountInfo.PacketAccountId, //收款账户
			UserId:        accountInfo.MainAccountId,
			SubMerchantId: "2206301126073014978", // 子商户编号
		}})
	if err != nil {
		return errors.New(fmt.Sprintf("充值失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return errors.New(fmt.Sprintf("充值失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	//增加账户变更日志
	err = AddNcountTradeLog(BusinessTypeBankcardSendPacket, amount, userId, accountInfo.MainAccountId, accountResp.NcountOrderId, packetID)
	if err != nil {
		return errors.New(fmt.Sprintf("增加账户变更日志失败(%s)", err.Error()))
	}

	return nil
}

// 红包领取明细
func (rpc *CloudWalletServer) RedPacketReceiveDetail(_ context.Context, req *pb.RedPacketReceiveDetailReq) (*pb.RedPacketReceiveDetailResp, error) {
	//查询时间转换
	sTime, _ := time.ParseInLocation("2006-01-02", req.StartTime, time.Local)
	eTime, _ := time.ParseInLocation("2006-01-02", req.EndTime, time.Local)

	//获取列表数据
	list, _ := imdb.FindReceiveRedPacketList(req.UserId, sTime.Unix(), eTime.Unix()+86399)

	receiveList := make([]*pb.RedPacketReceiveDetail, 0)
	for _, v := range list {
		receiveList = append(receiveList, &pb.RedPacketReceiveDetail{
			PacketId:    v.PacketId,
			Amount:      v.Amount,
			PacketTitle: v.PacketTitle,
			ReceiveTime: time.Unix(v.ReceiveTime, 0).Format("2006-01-02 15:04:05"),
			PacketType:  v.PacketType,
			IsLucky:     v.IsLucky,
		})
	}

	return &pb.RedPacketReceiveDetailResp{
		RedPacketReceiveDetail: receiveList,
	}, nil
}

// 红包详情
func (rpc *CloudWalletServer) RedPacketInfo(_ context.Context, req *pb.RedPacketInfoReq) (*pb.RedPacketInfoResp, error) {
	//获取红包记录
	redPacketInfo, err := imdb.GetRedPacketInfo(req.PacketId)
	if err != nil || redPacketInfo.UserID != req.UserId {
		return nil, errors.New("红包信息不存在")
	}

	//补充发红包人的用户信息
	nickname, faceUrl := "", ""
	userInfo, err := im_mysql_model.GetUserByUserID(req.UserId)
	if err == nil {
		nickname = userInfo.Nickname
		faceUrl = userInfo.FaceURL
	}

	info := &pb.RedPacketInfoResp{
		UserId:          redPacketInfo.UserID,
		PacketType:      redPacketInfo.PacketType,
		IsLucky:         redPacketInfo.IsLucky,
		IsExclusive:     redPacketInfo.IsExclusive,
		ExclusiveUserID: redPacketInfo.ExclusiveUserID,
		PacketTitle:     redPacketInfo.PacketTitle,
		Amount:          redPacketInfo.Amount,
		Number:          redPacketInfo.Number,
		ExpireTime:      redPacketInfo.ExpireTime,
		Remain:          redPacketInfo.Remain,
		Nickname:        nickname,
		FaceUrl:         faceUrl,
		ReceiveDetail:   make([]*pb.ReceiveDetail, 0),
	}

	//获取当前红包领取记录
	receiveList, _ := imdb.ReceiveListByPacketId(req.PacketId)
	for _, v := range receiveList {
		info.ReceiveDetail = append(info.ReceiveDetail, &pb.ReceiveDetail{
			UserId:      v.UserId,
			Amount:      v.Amount,
			Nickname:    v.Nickname,
			FaceUrl:     v.FaceUrl,
			ReceiveTime: time.Unix(v.ReceiveTime, 0).Format("01月02日 15:04"),
		})
	}

	return info, nil
}
