package cloud_wallet

import (
	ncount "Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/config"
	commonDB "Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	imdb2 "Open_IM/pkg/common/db/mysql_model/im_mysql_model"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/contrive_msg"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/tools/redpacket"
	"context"
	"database/sql"
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

// 发送红包
func (h *handlerSendRedPacket) SendRedPacket(req *pb.SendRedPacketReq) (*pb.SendRedPacketResp, error) {
	var (
		result = &pb.SendRedPacketResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "发送成功",
			},
		}
	)
	// 1. 校验参数
	if err := h.validateParam(req); err != nil {
		return nil, err
	}

	// ========================================= 验证发送用户的信息=========================================
	userAC, err := imdb.GetNcountAccountByUserId(req.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			result.CommonResp.ErrMsg = "账户信息错误"
			result.CommonResp.ErrCode = 400
			return result, nil
		}
		return nil, errors.New("网络错误：当前账户信息错误")
	}

	if req.Password != userAC.PaymentPassword {
		result.CommonResp.ErrMsg = "支付密码错误"
		result.CommonResp.ErrCode = 400
		return result, nil
	}

	// todo 暂时这么处理
	if checkGroupValidate := h.checkGroupPacketState(req); checkGroupValidate != "" {
		return nil, errors.New(checkGroupValidate)
	}

	// ========================================= 查看发送用户ID是否存在 =========================================
	if req.PacketType == 1 {
		// 单聊红包
		user, err := imdb2.GetUserByUserID(req.RecvID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.CommonResp.ErrMsg = "您发送红包的用户不存在"
				result.CommonResp.ErrCode = 400
				return result, nil
			}
			return nil, errors.New("查询用户信息失败")
		}
		if user.UserID == "" {
			result.CommonResp.ErrMsg = "您发送红包的用户不存在"
			result.CommonResp.ErrCode = 400
			return result, nil
		}
	} else {
		// 群聊红包
		group, err := imdb2.GetGroupInfoByGroupID(req.RecvID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.CommonResp.ErrMsg = "您发送的红包群不存在"
				result.CommonResp.ErrCode = 400
				return result, nil
			}
			return nil, errors.New("查询群信息失败")
		}
		if group.GroupID == "" {
			result.CommonResp.ErrMsg = "您发送的红包群不存在"
			result.CommonResp.ErrCode = 400
			return result, nil
		}
	}

	// ========================================= 创建红包记录 =========================================
	redpacketID, err := h.recordRedPacket(req, userAC.PacketAccountId)
	if err != nil {
		log.Error(req.OperationID, "record red packet error", zap.Error(err))
		return nil, err
	}
	res := &pb.SendRedPacketResp{
		RedPacketID: redpacketID,
	}

	// 3. 判断支付类型
	if req.SendType == 1 {
		transferMssgae, err := h.walletTransfer(redpacketID, req)
		if err != nil {
			log.Error(req.OperationID, "转账失败", err)
			return nil, err
		}

		if transferMssgae != "" {
			result.CommonResp.ErrMsg = transferMssgae
			result.CommonResp.ErrCode = 400
			return result, nil
		}

		// 回调处理红包
		err = HandleSendPacketResult(redpacketID, req.OperationID)
		if err != nil {
			log.Error(req.OperationID, "HandleSendPacketResult error", zap.Error(err))
			return nil, err
		}
	} else {
		// 走银行卡转账
		transferMsg, err := h.bankTransfer(redpacketID, req)
		// 这里是调用银行卡转账接口
		if err != nil {
			log.Error(req.OperationID, "bankTransfer error", zap.Error(err))
			return nil, err
		}
		if transferMsg != "" {
			result.CommonResp.ErrMsg = transferMsg
			result.CommonResp.ErrCode = 400
			return result, nil
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
func (h *handlerSendRedPacket) recordRedPacket(in *pb.SendRedPacketReq, packetID /*发红包的用户ID*/ string) (string /* red packet ID */, error) {
	rand.Seed(time.Now().UnixNano())
	redID := fmt.Sprintf("%v%v%v%v", in.PacketType, in.UserId, time.Now().Unix(), rand.Intn(100000))
	redPacket := &imdb.FPacket{
		PacketID:             redID,
		UserID:               in.UserId,
		UserRedpacketAccount: packetID,
		PacketType:           in.PacketType,
		IsLucky:              in.IsLucky,
		ExclusiveUserID:      in.ExclusiveUserID,
		PacketTitle:          in.PacketTitle,
		Amount:               in.Amount,
		Number:               in.Number,
		MerOrderID:           h.merOrderID,
		OperateID:            h.OperateID,
		SendType:             in.SendType,
		BindCardAgrNo:        in.BindCardAgrNo,
		RecvID:               in.RecvID, // 接收ID
		Remain:               int64(in.Number),
		RemainAmout:          in.Amount,
		ExpireTime:           time.Now().Unix() + 60*60*24,
		CreatedTime:          time.Now().Unix(),
		UpdatedTime:          time.Now().Unix(),
		Status:               0, // 红包被创建，但是还未掉第三方的内容
		IsExclusive:          in.IsExclusive,
	}
	return redID, imdb.RedPacketCreateData(redPacket)
}

// 走用户的钱包转账
func (h *handlerSendRedPacket) walletTransfer(redPacketID string, in *pb.SendRedPacketReq) (string, error) {
	var res string
	// 1. 获取用户的钱包账户
	fncount, err := imdb.FNcountAccountGetUserAccountID(in.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			res = "用户没有实名注册"
			return res, nil
		}
		return "", errors.Wrap(err, "查询用户信息错误")
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
	if err != nil {
		log.Error(in.OperationID, "调用新生支付出现错误", transferResult)
		res = "第三方支付出现错误 ，操作：" + h.OperateID
		return res, nil
	}

	// ========================================================下面是成功返回  =========================================
	if transferResult.ResultCode != ncount.ResultCodeSuccess {
		// 如果转账失败，需要给用户提示发送失败，并将红包状态修改为发送失败
		err := imdb.UpdateRedPacketStatus(redPacketID, 2 /* 发送失败 */)
		if err != nil {
			log.Error(in.OperationID, zap.Error(err))
			// todo 记录到死信队列中，后续监控处理 ，不需要人工介入
		}
		// 记录操作失败日志 ，提供后续给客服人员核对
		res = transferResult.ErrorMsg + "，操作：" + h.OperateID
		return res, nil
	}
	// 如果转账成功，需要将红包状态修改为发送成功
	err = imdb.UpdateRedPacketStatus(redPacketID, 1 /* 发送成功 */)
	if err != nil {
		// todo 记录到死信队列中，后续监控处理， 如果转账成功，但是修改红包状态失败，需要人工介入
		log.Error(in.OperationID, zap.Error(err))
		return "", errors.Wrap(err, "修改红包状态失败 1")
	}

	//增加账户变更日志
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(in.OperationID, "增加账户变更日志失败:Panic", zap.Any("err", err))
			}
		}()
		err = AddNcountTradeLog(BusinessTypeBalanceSendPacket, int32(in.Amount), in.UserId, fncount.MainAccountId, transferResult.MerOrderId, redPacketID)
		if err != nil {
			log.Error(in.OperationID, "增加账户变更日志失败", zap.Error(err))
		}
	}()

	return res, nil
}

// 银行卡转账
func (h *handlerSendRedPacket) bankTransfer(redPacketID string, in *pb.SendRedPacketReq) (string, error) {
	//银行卡充值到红包账户
	err := BankCardRechargePacketAccount(in.UserId, in.BindCardAgrNo, int32(in.Amount), redPacketID)
	if err != nil {
		return "", err
	}

	// 如果转账成功，需要将红包状态修改为发送成功
	err = imdb.UpdateRedPacketStatus(redPacketID, 1 /* 发送成功 */)
	if err != nil {
		// todo 记录到死信队列中，后续监控处理， 如果转账成功，但是修改红包状态失败，需要人工介入
		log.Error(in.OperationID, zap.Error(err))
		return "", errors.Wrap(err, "修改红包状态失败 1")
	}
	return "", nil
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
	if redpacketInfo.PacketType == 2 && redpacketInfo.Number > 1 {
		// 群红包
		err = GroupPacket(redpacketInfo, redPacketID)
		if err != nil {
			return err
		}
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

	// 发送红包消息
	return contrive_msg.SendSendRedPacket(freq, int(redpacketInfo.PacketType))
}

// 给群发的红包
func GroupPacket(req *imdb.FPacket, redpacketID string) error {

	// 在群红包这里，
	var err error
	if req.IsLucky == 1 {
		// 如果说是手气红包，分散放入红包池
		err = spareRedPacket(redpacketID, int(req.Amount), int(req.Number))
	} else {
		// 凭手气红包
		err = spareEqualRedPacket(redpacketID, int(req.Amount), int(req.Number))
	}
	if err != nil {
		log.Error(req.OperateID, zap.Error(err))
		return err
	}
	return err
}

// 将红包放入红包池
func spareRedPacket(packetID string, amount, number int) error {
	defer func() {
		if err := recover(); err != nil {
			log.Error("spareRedPacket panic", zap.Any("err", err))
		}
	}()
	// 将发送的红包进行计算
	result := redpacket.GetRedPacket(amount, number)
	err := commonDB.DB.SetRedPacket(packetID, result)
	if err != nil {
		return err
	}
	return nil
}

// amount = 3 ,number =3
func spareEqualRedPacket(packetID string, amount, number int) error {
	result := []int{}
	for i := 0; i < number; i++ {
		result = append(result, amount)
	}
	// 将发送的红包进行计算
	err := commonDB.DB.SetRedPacket(packetID, result)
	if err != nil {
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
	if err != nil {
		return nil, errors.New("红包信息不存在")
	}

	//补充发红包人的用户信息
	nickname, faceUrl := "", ""
	userInfo, err := im_mysql_model.GetUserByUserID(redPacketInfo.UserID)
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

// 禁止群抢红包
func (rpc *CloudWalletServer) ForbidGroupRedPacket(ctx context.Context, req *pb.ForbidGroupRedPacketReq) (*pb.ForbidGroupRedPacketResp, error) {
	var (
		result = &pb.ForbidGroupRedPacketResp{
			CommonResp: &pb.CommonResp{
				ErrMsg:  "禁止群抢红包成功",
				ErrCode: 0,
			},
		}
	)
	// 查看用户是否为群主
	group, err := imdb2.GetGroupInfoByGroupID(req.GroupId)
	if (err != nil && errors.Is(err, sql.ErrNoRows)) || group.GroupID == "" {
		result.CommonResp.ErrCode = 400
		result.CommonResp.ErrMsg = "群信息不存在"
		return result, nil
	}

	// 如果存在群，且用户不是群主
	if group.CreatorUserID != req.UserId {
		result.CommonResp.ErrCode = 400
		result.CommonResp.ErrMsg = "您不是群主"
		return result, nil
	}

	// 禁止抢红包
	err = imdb2.UpdateGroupIsAllowRedPacket(req.GroupId, req.Forbid)
	if err != nil {
		log.Error(req.OperationID, "禁止群抢红包失败", err)
		return nil, err
	}

	// 如果ok 删除
	err = rocksCache.DelGroupInfoFromCache(req.GroupId)
	if err != nil {
		log.Error(req.OperationID, "删除群缓存", err)
		return nil, err
	}

	return result, nil
}
