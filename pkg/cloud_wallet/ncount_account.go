package cloud_wallet

type PlatformNcountNewAccount struct {
	version    string `json:"version" binding:"required"`
	tranCode   string `json:"tranCode" binding:"required"`
	merId      string `json:"merId" binding:"required"`
	merOrderId string `json:"merOrderId" binding:"required"`
	submitTime string `json:"submitTime" binding:"required"`
	msgCiphert string `json:"msgCiphert" binding:"required"`
	signType   string `json:"signType" binding:"required"`
	signValue  string `json:"signValue" binding:"required"`
	merUserId  string `json:"merUserId" binding:"required"`
	mobile     string `json:"mobile" binding:"required"`
	userName   string `json:"userName" binding:"required"`
	certNo     string `json:"certNo" binding:"required"`
}

// 查询用户账户信息 的请求参数
type PlatformNcountCheckUserAccountInfo struct {
	version    string `json:"version" binding:"required"`
	tranCode   string `json:"tranCode" binding:"required"`
	merId      string `json:"merId" binding:"required"`
	merOrderId string `json:"merOrderId" binding:"required"`
	signType   string `json:"signType" binding:"required"`
	signValue  string `json:"signValue" binding:"required"`
}
