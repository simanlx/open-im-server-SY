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
	"fmt"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

// 第三方支付
func (cl *CloudWalletServer) ThirdPay(ctx context.Context, in *pb.ThirdPayReq) (*pb.ThirdPayResp, error) {
	var (
		res = &pb.ThirdPayResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "支付成功",
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
	err, payOrder := imdb.GetThirdPayOrderNo(in.OrderNo)
	if err != nil {
		return nil, err
	}
	if payOrder.Id == 0 {
		res.CommonResp.ErrCode = 400
		res.CommonResp.ErrMsg = "订单不存在"
		return res, nil
	}

	if payOrder.Status != 100 {
		res.CommonResp.ErrCode = 400
		res.CommonResp.ErrMsg = "订单状态异常:" + strconv.Itoa(int(payOrder.Status))
		return res, nil
	}

	// 计算具体余额
	totalAmount := cast.ToString(cast.ToFloat64(payOrder.Amount) / 100)

	nc := NewNcountPay()
	// 发起支付
	PayRes := &PayResult{}
	if in.SendType == 1 {
		// 余额支付
		PayRes = nc.payByBalance(in.OperationID, fcount.MainAccountId, "300002428690", payOrder.NcountOrderNo, totalAmount)
		if PayRes.ErrCode == 0 {
			// 支付成功
			err = AddNcountTradeLog(BusinessTypeBalanceThirdPay, int32(payOrder.Amount), in.Userid, fcount.MainAccountId, PayRes.NcountOrderID, "")
			if err != nil {
				log.Error(in.OperationID, "添加交易记录失败，err: ", err)
			}
			payOrder.Status = 200 // 支付成功
			payOrder.PayTime = time.Now()
			payOrder.NcountTureNo = PayRes.NcountOrderID
			// 修改订单状态
			err := imdb.UpdateThirdPayOrder(payOrder, payOrder.Id)
			if err != nil {
				log.Error(in.OperationID, "修改订单状态失败，err: ", err)
			}
			// 订单支付成功 ： todo 通知商户
		} else {
			// 支付失败
			res.CommonResp.ErrCode = pb.CloudWalletErrCode(PayRes.ErrCode)
			res.CommonResp.ErrMsg = "新生支付：" + PayRes.ErrMsg
		}
	} else {
		res.CommonResp.ErrMsg = "支付已提交，还需要进行支付确认"
		res.CommonResp.ErrCode = 101

		NotifyUrl := config.Config.Ncount.Notify.ThirdPayNotifyUrl
		// 银行卡支付 ，需要注意回调接口
		PayRes = nc.payByBankCard(in.OperationID, fcount.MainAccountId, "300002428690", payOrder.NcountOrderNo, totalAmount, in.BankcardProtocol, NotifyUrl)
		if PayRes.ErrCode == 0 {
			// 支付成功
			err = AddNcountTradeLog(BusinessTypeBankcardThirdPay, int32(payOrder.Amount), in.Userid, fcount.MainAccountId, PayRes.NcountOrderID, "")
			if err != nil {
				log.Error(in.OperationID, "添加交易记录失败，err: ", err)
			}
			// 修改订单信息
			payOrder.NcountTureNo = PayRes.NcountOrderID
			// 修改订单状态
			err := imdb.UpdateThirdPayOrder(payOrder, payOrder.Id)
			if err != nil {
				log.Error(in.OperationID, "修改订单状态失败，err: ", err)
			}

		} else {
			// 支付失败
			res.CommonResp.ErrCode = pb.CloudWalletErrCode(PayRes.ErrCode)
			res.CommonResp.ErrMsg = PayRes.ErrMsg
		}
	}
	// 保存交易记录
	return res, nil
}

// 创建订单
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
	err, payOrder := imdb.GetThirdPayMerOrderNO(req.MerOrderId)
	if err != nil {
		return nil, err
	}
	if payOrder.Id != 0 {
		resp.CommonResp.ErrMsg = "订单已存在"
		resp.CommonResp.ErrCode = 400
		return resp, nil
	}
	fmt.Println("\n payOrder ", payOrder)

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
		NcountOrderNo:  ncount.GetMerOrderID(),
		Amount:         int64(req.Amount),
		Status:         100,
		RecieveAccount: merchant.NcountAccount,
		Remark:         req.Remark,
		NotifyUrl:      req.NotifyUrl,
		LastNotifyTime: time.Time{},
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

// 查询红包
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

	err, payOrder := imdb.GetThirdPayOrderNo(req.OrderNo)
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
	resp.Remark = payOrder.Remark
	resp.AddTime = payOrder.AddTime.Format("2006-01-02 15:04:05")
	return resp, nil
}
