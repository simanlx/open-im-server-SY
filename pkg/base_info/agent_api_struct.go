package base_info

// 推广员申请提交
type AgentApplyReq struct {
	Name        string `json:"name"  binding:"required"`          //用户姓名
	Mobile      string `json:"mobile"  binding:"required"`        //用户手机号码
	ChessUserId int64  `json:"chess_user_id"  binding:"required"` //互娱用户id
}

type BindAgentNumberReq struct {
	AgentNumber   int32  `json:"agent_number"  binding:"required"`   //推广员编号
	ChessUserId   int64  `json:"chess_user_id"  binding:"required"`  //互娱用户id
	ChessNickname string `json:"chess_nickname"  binding:"required"` //互娱用户昵称
}

type GetUserAgentInfoReq struct {
	ChessUserId int64 `json:"chess_user_id"  binding:"required"` //互娱用户id
}

type AgentAccountIncomeChartReq struct {
	DateType int32 `json:"date_type"` //日期类型 1(7天),2(半年) 默认7天
}

type AgentAccountRecordListReq struct {
	Date         string `json:"date"`          //日期
	BusinessType int32  `json:"business_type"` //业务类型
	Keyword      string `json:"keyword"`       //搜索关键字
	Page         int32  `json:"page"`
	Size         int32  `json:"size"`
}

type AgentGameShopBeanConfigReq struct {
	AgentNumber int32 `json:"agent_number"  binding:"required"` //推广员编号
}

type AgentBeanAccountRecordListReq struct {
	Date         string `json:"date"`          //日期
	BusinessType int32  `json:"business_type"` //业务类型
	Page         int32  `json:"page"`
	Size         int32  `json:"size"`
}

type AgentBeanShopUpStatusReq struct {
	Status   int32 `json:"status"`    //状态(0下架、1上架)
	ConfigId int32 `json:"config_id"` //配置id
	IsAll    int32 `json:"is_all"`    //是否全部(1全部，0单个)
}

type AgentBeanShopUpdateReq struct {
	BeanShopConfig []*BeanShopConfig `json:"bean_shop_config"` //咖豆配置
}

type BeanShopConfig struct {
	BeanNumber     int64 `json:"bean_number"`
	GiveBeanNumber int32 `json:"give_bean_number"`
	Amount         int32 `json:"amount"`
	Status         int32 `json:"status"`
}

type AgentMemberListReq struct {
	Keyword string `json:"keyword"`  //搜索关键字
	OrderBy string `json:"order_by"` //排序
}
