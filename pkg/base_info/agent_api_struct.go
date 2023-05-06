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
}
