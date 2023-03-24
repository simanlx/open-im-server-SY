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
