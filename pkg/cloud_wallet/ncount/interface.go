package ncount

import "github.com/pkg/errors"

/*
// 创建用户账户地址
	NewAccountURL = "https://ncount.hnapay.com/api/r010.htm"

	// 用户查询接口
	checkUserAccountInfoURL = "https://ncount.hnapay.com/api/q001.htm"

	// 绑卡接口
	bindCardURL = "https://ncount.hnapay.com/api/r007.htm"

	// 绑卡确认接口
	bindCardConfirmURL = "https://ncount.hnapay.com/api/r008.htm"

	// 个人用户解绑接口
	unbindCardURL = "https://ncount.hnapay.com/api/r009.htm"

	// 用户账户明细查询接口：转账详情
	checkUserAccountDetailURL = "https://ncount.hnapay.com/api/q004.htm"

	// 交易查询接口
	checkUserAccountTransURL = "https://ncount.hnapay.com/api/q002.htm"

	// 快捷支付下单接口
	quickPayOrderURL = "https://ncount.hnapay.com/api/t007.htm"

	// 快捷支付确认接口
	quickPayConfirmURL = "https://ncount.hnapay.com/api/t008.htm"

	// 转账接口
	transferURL = "https://ncount.hnapay.com/api/t003.htm"

	// 退款接口
	refundURL = "https://ncount.hnapay.com/api/t005.htm"

*/

// NCounter is the platform ncount interface, the interface is provided
// by the platform, and the platform implements the interface.
type NCounter interface {
	// 创建用户账户地址
	NewAccount(req *NewAccountReq) (*NewAccountResp, error)
	// 用户查询接口
	CheckUserAccountInfo(req *CheckUserAccountInfoReq) (*CheckUserAccountInfoResp, error)
	// 绑卡接口
	BindCard(req *BindCardReq) (*BindCardResp, error)
	// 绑卡确认接口
	BindCardConfirm(req *BindCardConfirmReq) (*BindCardConfirmResp, error)
	// 个人用户解绑接口
	UnbindCard(req *UnbindCardReq) (*UnbindCardResp, error)
	// 用户账户明细查询接口：转账详情
	CheckUserAccountDetail(req *CheckUserAccountDetailReq) (*CheckUserAccountDetailResp, error)
	// 交易查询接口
	CheckUserAccountTrans(req *CheckUserAccountTransReq) (*CheckUserAccountTransResp, error)
	// 快捷支付下单接口
	QuickPayOrder(req *QuickPayOrderReq) (*QuickPayOrderResp, error)
	// 快捷支付确认接口
	QuickPayConfirm(req *QuickPayConfirmReq) (*QuickPayConfirmResp, error)
	// 转账接口
	Transfer(req *TransferReq) (*TransferResp, error)
	// 退款接口
	Refund(req *RefundReq) (*RefundResp, error)
}

// 发送请求
type NewAccountReq struct {
	BaseParam *NewAccountBaseParam
	MsgCipher *NewAccountMsgCipher
}

/*
	4.1.1.1 商户请求参数列表（POST）
	version 版本号 5 目前必须为 1.0 不可 例如：1.0
	tranCode 交易代码 5 此交易只能为 R010 不可 例如：R010
	merId 商户 ID 12 新账通平台提供给商户 的唯一 ID 不可 例 如 ： 00000100000
	merOrderId 商 户 订 单 号 1 - 32 格式：数字，字母，下 划线，竖划线，中划线 不可 例 如 ： aa201612011 102
	submitTime 请 求 提 交 时间 14 格 式 ： YYYYMMDDHHMMS S 不可 例 如 ： 20161201110 233
	msgCiphert ext 报文密文 1-4000 用平台公钥 RSA 加密后 base64 的编码值 不可
	signType 签名类型 1 1：RSA 不可 例如：1
	signValue 签 名 密 文 串 将 报 文 信 息 用signType 域设 置的方 式签名后生成的字符串 不可
	merAttach 附加数据 0-80 格式：英文字母/汉字 可空
	charset 编码方式 1 1：UTF8 不可 例如：1
*/

//商户请求参数列表（POST）
type NewAccountBaseParam struct {
	Version    string `json:"version" binding:"required"`
	TranCode   string `json:"tranCode" binding:"required"`
	MerId      string `json:"merId" binding:"required"`
	MerOrderId string `json:"merOrderId" binding:"required"`
	SubmitTime string `json:"submitTime" binding:"required"`
	MsgCiphert string `json:"msgCiphert" binding:"required"`
	SignType   string `json:"signType" binding:"required"`
	SignValue  string `json:"signValue" binding:"required"`
	Charset    string `json:"charset" binding:"required"`
}

// str : version=[]tranCode=[]merId=[]merOrderId=[]submitTime=[]msgCiphertext=[]signType=[]
// signValue = version
func (n *NewAccountBaseParam) flushSignValue() (error, string) {
	if n.Version == "" {
		return errors.New("version is empty"), ""
	}
	if n.TranCode == "" {
		return errors.New("tranCode is empty"), ""
	}
	if n.MerId == "" {
		return errors.New("merId is empty"), ""
	}
	if n.MerOrderId == "" {
		return errors.New("merOrderId is empty"), ""
	}
	if n.SubmitTime == "" {
		return errors.New("submitTime is empty"), ""
	}
	if n.MsgCiphert == "" {
		return errors.New("msgCiphert is empty"), ""
	}
	if n.SignType == "" {
		return errors.New("signType is empty"), ""
	}
	var str = ""
	str += "version=[" + n.Version + "]"
	str += "tranCode=[" + n.TranCode + "]"
	str += "merId=[" + n.MerId + "]"
	str += "merOrderId=[" + n.MerOrderId + "]"
	str += "submitTime=[" + n.SubmitTime + "]"
	str += "msgCiphertext=[" + n.MsgCiphert + "]"
	str += "signType=[" + n.SignType + "]"
	return nil, str
}

/*
	merUserId 商 户 用户 唯 一标识 1-30 按照商户侧规则 同一个平台商户唯一 字符串类型 不可 例 如 ： xsqianyi1_ 148080610 02
	mobile 用户手机号 11 用户本人在运营商已 实名的手机号 不可 例如： 138000000 00
	userName 真实姓名 1-30 和身份证上的姓名保 持一致 不可 例如：张三
	certNo 身份证号 18 身份证号码 不可 例如： 110210199 008123456
*/
type NewAccountMsgCipher struct {
	MerUserId string `json:"merUserId" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	UserName  string `json:"userName" binding:"required"`
	CertNo    string `json:"certNo" binding:"required"`
}

/*
	resultCode 处理结果码 4 详情参见附录二
	resultCode 9999 errorCode 异常代码 1-10 详情参见附录一
	errorCode errorMsg 异常描述 1-200 中文、字母、数字
	userId 用户编号 12 用户编号
	signValue 签名字符串 将报文信息用
	signType 域设 置的方式签名后生成的字符 串
*/

// 调用新生接口创建账户，拿到的返回结果
type NewAccountReturnFromPlatform struct {
	ResultCode string `json:"resultCode" binding:"required"`
	ErrorCode  string `json:"errorCode" binding:"required"`
	ErrorMsg   string `json:"errorMsg" binding:"required"`
	UserId     string `json:"userId" binding:"required"`
	SignValue  string `json:"signValue" binding:"required"`
	SignType   string `json:"signtype" binding:"required"`
}

type NewAccountResp struct {
	version    string `json:"version" binding:"required"`
	merAttach  string `json:"merAttach" binding:"required"`
	charset    string `json:"charset" binding:"required"`
	resultCode string `json:"resultCode" binding:"required"`
	errorCode  string `json:"errorCode" binding:"required"`
	errorMsg   string `json:"errorMsg" binding:"required"`
	userId     string `json:"userId" binding:"required"`
	signValue  string `json:"signValue" binding:"required"`
	signType   string `json:"signtype" binding:"required"`
}

// 报文加密流程：
//	1. 将报文信息转换为 JSON 格式
//	2. 将 JSON 格式的报文信息用平台公钥 RSA 加密后 base64 的编码值
// 	3. 将加密后的报文信息放入 msgCipher 字段中

type CheckUserAccountInfoReq struct {
}

type CheckUserAccountInfoResp struct {
}

type BindCardReq struct {
}

type BindCardResp struct {
}

type BindCardConfirmReq struct {
}

type BindCardConfirmResp struct {
}

type UnbindCardReq struct {
}

type UnbindCardResp struct {
}

type CheckUserAccountDetailReq struct {
}

type CheckUserAccountDetailResp struct {
}

type CheckUserAccountTransReq struct {
}

type CheckUserAccountTransResp struct {
}

type QuickPayOrderReq struct {
}

type QuickPayOrderResp struct {
}

type QuickPayConfirmReq struct {
}

type QuickPayConfirmResp struct {
}

type TransferReq struct {
}

type TransferResp struct {
}

type RefundReq struct {
}

type RefundResp struct {
}
