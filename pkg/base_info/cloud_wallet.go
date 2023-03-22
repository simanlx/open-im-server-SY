package base_info

type CheckUserHaveAccountResp struct {
	HaveAccount bool `json:"have_account"`
	Step        int  `json:"step"` // 1:未实名认证 2:未绑定银行卡 3:未设置交易密码 4:已开户
}

/*
message idCardRealNameAuthReq{
  uint32 userID = 1; //用户id
  string idCard = 2; //身份证
  string realName = 3; //真实姓名
  string mobile = 4; //手机号码
}
*/
type CreateUserAccount struct {
	IdCard   string `json:"id_card"`
	RealName string `json:"real_name"`
	Mobile   string `json:"mobile"`
}

type CreateUserAccountresp struct {
	Code int    `json:"code"` // 0:成功 1:失败 11: 未实名认证 12: 未绑定银行卡 13: 未设置交易密码
	Msg  string `json:"msg"`  // 失败原因
}
