package account

// 账户信息
type AccountReq struct {
	UserId      string `json:"user_id"`
	OperationID string `json:"operationID" binding:"required"`
}

// 身份证实名认证
type IdCardRealNameAuthReq struct {
	UserId      string `json:"user_id"`
	IdCard      string `json:"id_card"  binding:"required"`
	Mobile      string `json:"mobile"  binding:"required"`
	RealName    string `json:"real_name"  binding:"required"`
	OperationID string `json:"operationID"  binding:"required"`
}

// 设置支付密码
type SetPaymentSecretReq struct {
	UserId        string `json:"user_id"`
	Type          int32  `json:"type" binding:"required"` //设置类型(1设置密码、2忘记密码smsCode设置)
	Code          string `json:"code"`
	PaymentSecret string `json:"payment_secret" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
}

// 校验支付密码
type CheckPaymentSecretReq struct {
	UserId        string `json:"user_id"`
	PaymentSecret string `json:"payment_secret" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
}

// 绑定银行卡
type BindUserBankCardReq struct {
	UserId            string `json:"user_id"`
	CardOwner         string `json:"card_owner" binding:"required"` //持卡人姓名
	BankCard          string `json:"bank_card" binding:"required"`  //银行卡
	Mobile            string `json:"mobile" binding:"required"`     //签约手机号码
	CardAvailableDate string `json:"cardAvailableDate"`             //信用卡有效期
	Cvv2              string `json:"cvv2"`                          //信用卡cvv2
	OperationID       string `json:"operationID" binding:"required"`
}

// 绑定银行卡确认-code
type BindUserBankcardConfirmReq struct {
	UserId      string `json:"user_id"`
	BankCardId  int32  `json:"bank_card_id" binding:"required"`
	Code        string `json:"code" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}

// 解绑银行卡
type UnBindUserBankcardReq struct {
	UserId        string `json:"user_id"`
	BindCardAgrNo string `json:"bindCardAgrNo" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
}

// 充值
type UserRechargeReq struct {
	UserId        string `json:"user_id"`
	AccountType   int32  `json:"account_type" binding:"required"` //充值账户类型
	Amount        int32  `json:"amount" binding:"required"`
	BindCardAgrNo string `json:"bindCardAgrNo" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
}

// 充值code确认
type UserRechargeConfirmReq struct {
	UserId      string `json:"user_id"`
	OrderNo     string `json:"order_no"  binding:"required"`
	Code        string `json:"code"  binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}

// 提现
type DrawAccountReq struct {
	UserId          string  `json:"user_id"`
	BindCardAgrNo   string  `json:"bindCardAgrNo"  binding:"required"`
	Amount          float32 `json:"amount"  binding:"required"`
	PaymentPassword string  `json:"payment_password"  binding:"required"`
	OperationID     string  `json:"operationID" binding:"required"`
}

// 云钱包账户明细
type CloudWalletRecordListReq struct {
	UserId      string `json:"user_id"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Page        int32  `json:"page"`
	Size        int32  `json:"size"`
	OperationID string `json:"operationID" binding:"required"`
}
