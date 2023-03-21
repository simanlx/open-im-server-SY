package cloud_wallet

const (
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
)
