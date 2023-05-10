package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"time"
)

func (cl *CloudWalletServer) ThirdPay(ctx context.Context, in *pb.ThirdPayReq) (*pb.ThirdPayResp, error) {
	var (
		res = &pb.ThirdPayResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "领取成功",
			},
		}
	)
	// ======================= 参数校验 =======================
	if in.SendType == 2 {
		// 银行卡发送
		if in.BankcardProtocol == "" {
			res.CommonResp.ErrCode = 400
			res.CommonResp.ErrMsg = "银行卡协议不能为空"
			return res, nil
		}
	}

	// 用户是否实名
	fcount, err := rocksCache.GetUserAccountInfoFromCache(in.Userid)
	if err != nil {
		// 这里redis可能出现错误，但是都可以进行这么上报
		res.CommonResp.ErrCode = pb.CloudWalletErrCode_UserNotValidate
		res.CommonResp.ErrMsg = "您的帐号没有实名认证,请尽快去实名认证"
		return nil, err
	}

	// 校验密码
	if fcount.PaymentPassword != in.Password {
		res.CommonResp.ErrCode = 400
		res.CommonResp.ErrMsg = "支付密码错误"
		return res, nil
	}

	// 查询订单是否存在
	err, payOrder := imdb.GetThirdPayOrder(in.OrderNo)
	if err != nil {
		return nil, err
	}
	if payOrder.Id == 0 {
		res.CommonResp.ErrCode = 400
		res.CommonResp.ErrMsg = "订单不存在"
		return res, nil
	}

	// 计算具体余额
	totalAmount := cast.ToString(cast.ToFloat64(payOrder.Amount) / 100)
	merOrderId := ncount.GetMerOrderID()
	nc := &NcountPay{}
	// 发起支付
	PayRes := &PayResult{}
	if in.SendType == 1 {
		// 余额支付
		PayRes = nc.payByBalance(in.OperationID, fcount.MainAccountId, "300002428690", merOrderId, totalAmount)
	} else {
		NotifyUrl := config.Config.Ncount.Notify.ThirdPayNotifyUrl
		// 银行卡支付 ，需要注意回调接口
		PayRes = nc.payByBankCard(in.OperationID, fcount.MainAccountId, "300002428690", merOrderId, totalAmount, in.BankcardProtocol, NotifyUrl)
	}
	if err != nil {
		res.CommonResp.ErrCode = pb.CloudWalletErrCode(PayRes.ErrCode)
		res.CommonResp.ErrMsg = PayRes.ErrMsg
	}
	return res, nil
}

func (cl *CloudWalletServer) CreateThirdPayOrder(ctx context.Context, req *pb.CreateThirdPayOrderReq) (*pb.CreateThirdPayOrderResp, error) {

	var (
		resp = &pb.CreateThirdPayOrderResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "订单创建成功",
			},
		}
	)
	// 查询用户上传的merchant是否存在
	merchant, err := imdb.GetMerchant(req.MerchantId)
	if err != nil {
		log.Error(req.OperationID, "查询商户失败，err: ", err)
		resp.CommonResp.ErrCode = 400
		resp.CommonResp.ErrMsg = "网络错误"
		return resp, nil
	}
	if merchant.Id == 0 {
		resp.CommonResp.ErrCode = 400
		resp.CommonResp.ErrMsg = "商户号不存在"
		return resp, nil
	}

	// 查询订单是否存在
	err, payOrder := imdb.GetThirdPayOrder(req.MerOrderId)
	if err != nil {
		return nil, err
	}
	if payOrder.Id != 0 {
		resp.CommonResp.ErrMsg = "订单已存在"
		resp.CommonResp.ErrCode = 400
		return resp, nil
	}

	// 生成随机数5位
	random := cast.ToString(time.Now().UnixNano())
	// 生成订单号前缀： 201805061203
	orderNoPrefix := time.Now().Format("200601021504")
	// 生成订单号
	orderNo := orderNoPrefix + random
	// 创建订单
	order := db.ThirdPayOrder{
		OrderNo:        orderNo,
		MerOrderNo:     req.MerOrderId,
		MerId:          req.MerchantId,
		Amount:         int64(req.Amount),
		Status:         100,
		RecieveAccount: merchant.NcountAccount,
		AddTime:        time.Time{},
		EditTime:       time.Time{},
	}

	err = imdb.InsertThirdPayOrder(&order)
	if err != nil {
		return nil, err
	}

	resp.OrderNo = order.OrderNo
	return resp, nil
}

func (cl *CloudWalletServer) GetThirdPayOrderInfo(ctx context.Context, req *pb.GetThirdPayOrderInfoReq) (*pb.GetThirdPayOrderInfoResp, error) {
	var (
		resp = &pb.GetThirdPayOrderInfoResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "查询成功",
			},
		}
	)
	// 获取第三方的订单信息
	if req.OrderNo == "" {
		resp.CommonResp.ErrCode = 400
		resp.CommonResp.ErrMsg = "订单号不能为空"
		return resp, nil
	}

	err, payOrder := imdb.GetThirdPayOrder(req.OrderNo)
	if err != nil {
		log.Error(req.OperationID, "查询订单失败，err: ", err)
		resp.CommonResp.ErrCode = 400
		resp.CommonResp.ErrMsg = "网络错误"
		return nil, err
	}
	if payOrder.Id == 0 {
		resp.CommonResp.ErrCode = 400
		resp.CommonResp.ErrMsg = "订单不存在"
		return resp, nil
	}
	resp.OrderNo = payOrder.OrderNo
	resp.MerOrderId = payOrder.MerOrderNo
	resp.MerchantId = payOrder.MerId
	resp.Amount = int32(payOrder.Amount)
	resp.Status = payOrder.Status
	return resp, nil
}

type NcountPay struct {
	count ncount.NCounter
}

type PayResult struct {
	ErrMsg  string
	ErrCode int
}

// 余额支付
func (np *NcountPay) payByBalance(operationId, payAccountID, ReceiveAccountId, MerOrderId, totalAmount string) *PayResult {
	var (
		resp = &PayResult{
			ErrCode: 0,
			ErrMsg:  "支付成功",
		}
	)
	req := &ncount.TransferReq{
		MerOrderId: MerOrderId,
		TransferMsgCipher: ncount.TransferMsgCipher{
			PayUserId:     payAccountID,
			ReceiveUserId: ReceiveAccountId,
			TranAmount:    totalAmount, //分转元
		},
	}

	escap := time.Now()
	transferResult, err := np.count.Transfer(req)
	log.Info(operationId, "transfer req", req, "耗费时间:", time.Since(escap))
	if err != nil {
		//这里是网络层面的错误
		log.Error(operationId, "调用第三方支付出现网络错误", err)
		resp.ErrMsg = "调用第三方支付出现网络错误"
		resp.ErrCode = 400
		return resp
	}
	// 备注； 现在一般httpcode是自有逻辑实现，不用在专门判断

	//=========================成功返回=========================
	if transferResult.ResultCode != ncount.ResultCodeSuccess {
		remark := "余额发送红包失败： "
		co, _ := json.Marshal(transferResult)
		err := imdb.CreateErrorLog(remark, operationId, MerOrderId, transferResult.ErrorMsg, transferResult.ErrorCode, string(co))
		if err != nil {
			log.Error(operationId, "创建错误日志失败", err)
		}
		resp.ErrCode = 400
	}
	resp.ErrMsg = transferResult.ErrorMsg
	return resp
}

// 银行卡支付
func (np *NcountPay) payByBankCard(operationId, payAccountID, ReceiveAccountId, MerOrderId, totalAmount, BankProtocol, NotifyUrl string) *PayResult {
	var (
		resp = &PayResult{
			ErrCode: 0,
			ErrMsg:  "支付成功",
		}
	)
	//充值支付
	transferResult, err := np.count.QuickPayOrder(&ncount.QuickPayOrderReq{
		MerOrderId: MerOrderId,
		QuickPayMsgCipher: ncount.QuickPayMsgCipher{
			PayType:       "3", //绑卡协议号充值
			TranAmount:    totalAmount,
			NotifyUrl:     NotifyUrl,
			BindCardAgrNo: BankProtocol,
			ReceiveUserId: ReceiveAccountId, //收款账户
			UserId:        payAccountID,
			SubMerchantId: "2206301126073014978", // 子商户编号
		}})
	if err != nil {
		//这里是网络层面的错误
		log.Error(operationId, "调用第三方支付出现网络错误", err)
		resp.ErrMsg = "调用第三方支付出现网络错误"
		resp.ErrCode = 400
		return resp
	}
	if transferResult.ResultCode != ncount.ResultCodeSuccess {
		remark := "余额发送红包失败： "
		co, _ := json.Marshal(transferResult)
		err := imdb.CreateErrorLog(remark, operationId, MerOrderId, transferResult.ErrorMsg, transferResult.ErrorCode, string(co))
		if err != nil {
			log.Error(operationId, "创建错误日志失败", err)
		}
		resp.ErrCode = 400
	}
	resp.ErrMsg = transferResult.ErrorMsg
	return resp
}
