package account

// AccountReq
type AccountReq struct {
	UserId      string `json:"user_id"`
	OperationID string `json:"operationID" binding:"required"`
}

// 身份证实名认证
type IdCardRealNameAuthReq struct {
	UserId      string `json:"user_id"`
	IdCard      string `json:"id_card"  binding:"required"`
	RealName    string `json:"real_name"  binding:"required"`
	OperationID string `json:"operationID"  binding:"required"`
}

// 设置支付密码
type SetPaymentSecretReq struct {
	UserId        string `json:"user_id"`
	PaymentSecret int64  `json:"payment_secret" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
}

// 绑定银行卡
type BindUserBankCardReq struct {
	UserId         string `json:"user_id"`
	CardOwner      string `json:"card_owner" binding:"required"`
	BankCardType   int32  `json:"bank_card_type" binding:"required"`
	BankCardNumber string `json:"bank_card_number" binding:"required"`
	Mobile         string `json:"mobile" binding:"required"`
	OperationID    string `json:"operationID" binding:"required"`
}

// 绑定银行卡确认-code
type BindUserBankcardConfirmReq struct {
	UserId      string `json:"user_id"`
	BankCardId  int32  `json:"bank_card_id" binding:"required"`
	Code        int32  `json:"code" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}
