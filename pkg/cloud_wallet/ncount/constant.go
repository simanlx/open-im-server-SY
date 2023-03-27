package ncount

const (
	// 新生支付公钥
	MER_USER_ID = "300002428690"
	PUBLIC_KEY  = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC9vGvYjivDF5uPBNDXMtoAtjYQ2YPSsfareduDG6kHL/N3A05rFHA11Dbr+UON82Y4V0RFKAQeZFPWcTLjcy6ntZVI8XoYLpuVQBPsb0Ya+PwbzR8/TmUdUf91ru8APtJgqkULgPVrO1hhzZ1tQMznosNLTOqbknMnnMcwzB5yYwIDAQAB\n-----END PUBLIC KEY-----"
	PRIVATE_KEY = `MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAKPjggbqm+RZOvoWMe6To3LeBlLWS8027RGqpAIJjfLaEu1HXvXX6q3Vcww7pYlzxp4fBiTcEvZ1gjTPq+N8/KyiFWRAO9Xs68HrEN2eRa92n3Gsu3XJFSJ7OeUOwAZtXQw6XlB3iIRa9XR1ueXsx8NUoGrl4mJq1rlgEvA5KGUJAgMBAAECgYBOVadH4QmkatYaxVMWgvEELYV+QLm4nAFSiWqdIq37nyqeyZdlENA2SKkV9siX24Pa/l80bRCPRvl2frDdKlem88q6D8PfdBaPRYVr950xXRLG7AAmE7YND4O6B81pQ46je28tQ/3jzwBN54/TlJJVWWQP76m5Zo/PUD3zdxQiAQJBAN7v6YaMASuUmO5DeP2C8oAnUxABhssdRgqzhbZ73bqn7kGHYdWE3TZau52UCoy+KcYyGNxTuxQr3kWTUrj9S0ECQQC8Mb9SARqKILJGVwdGRlSAS5zgnR356/0NCTdP5vws2DeXhHV50jpaYsyXBCLyFkBXNwX+2qw0+qbuOf0of4/JAkAnWBPQiPjT5h+vPP0nUGrXrxj7pClTw1DPJqucbvPMs0JbEjdz5UTdCNo/jxbli9H3hnPYvnYvsyZBBST+PMWBAkEAqPADbgrdlydYwbn4JsaVroGx9xQzx5lnlN80Dv8sWtlRtitLBauJhH/yZpJpCGafJWuYbzo/omNrnKjjsAoquQJAGMNbUXZteQJ9B0uCbSRx0KpJkw3+Ibvf/L7VRs2HCKqXQgU1xlrKxv1kgc9jhOQvwMxTGLTUrD9NOXV2w+Kapw==`
)

const (
	// 创建用户账户地址
	NewAccountURL = "https://ncount.hnapay.com/api/r010.htm"

	// 用户查询接口
	checkUserAccountURL = "https://ncount.hnapay.com/api/q001.htm"

	// 绑卡接口
	bindCardURL = "https://ncount.hnapay.com/api/r007.htm"

	// 绑卡确认接口
	bindCardConfirmURL = "https://ncount.hnapay.com/api/r008.htm"

	// 个人用户解绑接口
	unbindCardURL = "https://ncount.hnapay.com/api/r009.htm"

	// 用户账户详情接口
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

	// 提现接口
	withdrawURL = "https://ncount.hnapay.com/api/t002.htm"
)
