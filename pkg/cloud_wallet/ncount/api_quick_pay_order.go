package ncount

import (
	"github.com/pkg/errors"
)

// 快捷支付
/*
	tranAmount 支付金额 1-12 格式：数字（以元为单 位） 不可 例如： 100
	payType 支付方式 1 2:银行卡卡号 3:绑卡协议号 不可 例如：2
	cardNo 支付银行 卡卡号 0-30 payType=2 不可空 可空 例 如 ： 611888 812128
	holderName 持卡人姓 名 0-40payType=2 不可空 可空
	cardAvailableDate 信用卡有 效期 0-4 payType=2，且为 信用卡时不可空 可空 例 如 ： 0320 含 义 ： 2020 年 03 月 cvv2 信 用 卡
	CVV2 0-3 payType=2，且为 信用卡时不可空 可空 例 如 ： 318
	mobileNo 银行签约 手机号 011 payType=2 不可空 可空
	identityType 证件类型 0-2
	payType=2 不可空 暂仅支持 1 身份证 可空
	identityCode 证件号码 0-50
	payType=2 不可空 可空
	bindCardAgrNo 绑卡协议 号 30
	payType=3 不可空 可空
	notifyUrl 商户异步 通知地址 1-255 后台通知地址 不可 例 如 ： https:/ /www.x
*/
type QuickPayMsgCipher struct {
	TranAmount        string `json:"tranAmount" binding:"required"`        // 支付金额
	PayType           string `json:"payType" binding:"required"`           // 支付方式
	CardNo            string `json:"cardNo" binding:"required"`            // 支付银行卡卡号
	HolderName        string `json:"holderName" binding:"required"`        // 持卡人姓名
	CardAvailableDate string `json:"cardAvailableDate" binding:"required"` // 信用卡有效期
	Cvv2              string `json:"cvv2" binding:"required"`              // 信用卡cvv2
	MobileNo          string `json:"mobileNo" binding:"required"`          // 银行签约手机号
	IdentityType      string `json:"identityType" binding:"required"`      // 证件类型
	IdentityCode      string `json:"identityCode" binding:"required"`      // 证件号码
	BindCardAgrNo     string `json:"bindCardAgrNo" binding:"required"`     // 绑卡协议号
	NotifyUrl         string `json:"notifyUrl" binding:"required"`         // 商户异步通知地址

	/*
		orderExpireTime 订单过期 时长 0-1440 订单过期时长（单位：分 钟） 可空
		userId 用户编号 1-32 协议支付时，必填，要素 支付时，可空 可空 例 如 ： 102121
		receiveUserId 收款方 ID 1-32 消费交易时，填收款方 ID 担保交易时，填商户 ID 不可 例如： 102121
		merUserIp 商户用户 IP 0-128 商户用户签约时所在的机 器 IP 地址 可空 例 如 ： 211.12. 38.88
		riskExpand 风控扩展 信息 0-80 风控扩展信息 可空
		goodsInfo 商品信息 0-80 商品信息 可空
		subMerchantId 商户渠道 进件 ID 0-100 商户渠道进件 ID 不可
		divideFlag 是否分账 1 是否分账 0：否 （默认） 1：是 可空
		divideDetail 分账明细 信息 0-4000 分账明细 可空
		instalmentNum 分期期数 0-2 只支持 3、6、12、24 可空 例 如 ： 12
		instalmentType 商户补贴 分期手续 费方式 1 0-不贴息，1-贴息，2-全 额贴息 当分期期数 不为空 时， 此项不能为 空 例如：1
		instalmentRate 商户分期 贴息费率 6 取值为商户补贴手续费率 *100000，固定 6 位，不 为 6 位时前面补 0。如分 期手续费为 5%，商户补贴 3% ， 那 么 该 值 为 3%*100000 =003000 当分期期数 不为空 时， 此项不能为 空 例 如 ： 003000
	*/
	OrderExpireTime string `json:"orderExpireTime" binding:"required"`
	UserId          string `json:"userId" binding:"required"`
	ReceiveUserId   string `json:"receiveUserId" binding:"required"`
	MerUserIp       string `json:"merUserIp" binding:"required"`
	RiskExpand      string `json:"riskExpand" binding:"required"`
	GoodsInfo       string `json:"goodsInfo" binding:"required"`
	SubMerchantId   string `json:"subMerchantId" binding:"required"`
	DivideFlag      string `json:"divideFlag" binding:"required"`
	DivideDetail    string `json:"divideDetail" binding:"required"`
	InstalmentNum   string `json:"instalmentNum" binding:"required"`
	InstalmentType  string `json:"instalmentType" binding:"required"`
	InstalmentRate  string `json:"instalmentRate" binding:"required"`
}

func (q *QuickPayMsgCipher) Valid() error {
	if q.TranAmount == "" {
		return errors.New("支付金额不能为空")
	}
	if q.PayType == "" {
		return errors.New("支付方式不能为空")
	}
	if q.NotifyUrl == "" {
		return errors.New("异步通知地址不能为空")
	}
	if q.SubMerchantId == "" {
		return errors.New("商户渠道进件ID不能为空")
	}
	if q.ReceiveUserId == "" {
		return errors.New("收款方ID不能为空")
	}

	if q.PayType == "2" {
		if q.CardNo == "" {
			return errors.New("银行卡卡号不能为空")
		}
		if q.HolderName == "" {
			return errors.New("持卡人姓名不能为空")
		}
		if q.CardAvailableDate == "" {
			return errors.New("信用卡有效期不能为空")
		}
		if q.Cvv2 == "" {
			return errors.New("CVV2不能为空")
		}
		if q.MobileNo == "" {
			return errors.New("银行签约手机号不能为空")
		}
		if q.IdentityType == "" {
			return errors.New("证件类型不能为空")
		}
		if q.IdentityCode == "" {
			return errors.New("证件号码不能为空")
		}
	}

	if q.PayType == "3" {
		if q.BindCardAgrNo == "" {
			return errors.New("绑卡协议号不能为空")
		}
	}
	return nil
}

type QuickPayOrderReq struct {
	merOrderId        string `json:"merOrderId" binding:"required"`
	QuickPayMsgCipher QuickPayMsgCipher
}

func (q *QuickPayOrderReq) Valid() error {
	if q.merOrderId == "" {
		return errors.New("商户订单号不能为空")
	}
	return q.QuickPayMsgCipher.Valid()
}

/*
	resultCode 处理结果码 4 详情参见附录二 resultCode 9999
	errorCode 异常代码 1-10 详情参见附录一 errorCode
	errorMsg 异常描述 1-200 中文、字母、数字
	ncountOrderId 新账通订单 号 32 新账通平台交易订单号
	submitTime 商户请求时 间 同上送
	signValue 签名字符串 将报文信息用
	signType 域设 置的方式签名后生成的字符 串
*/
type QuickPayOrderResp struct {
	BaseReturnParam
	ResultCode    string `json:"resultCode"`
	ErrorCode     string `json:"errorCode"`
	ErrorMsg      string `json:"errorMsg"`
	NcountOrderId string `json:"ncountOrderId"`
	SubmitTime    string `json:"submitTime"`
	SignValue     string `json:"signValue"`
}
