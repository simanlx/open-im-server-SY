package agent

import (
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/agent_model"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/proto/agent"
	"context"
)

// 推广员申请提交
func (rpc *AgentServer) AgentApply(_ context.Context, req *agent.AgentApplyReq) (*agent.AgentApplyResp, error) {
	resp := &agent.AgentApplyResp{CommonResp: &agent.CommonResp{Code: 0, Msg: ""}}

	//查询是否已申请
	info, err := imdb.ApplyInfo(req.ChessUserId)
	if info != nil {
		resp.CommonResp.Msg = "已提交申请，请勿重复提交"
		resp.CommonResp.Code = 400
		return resp, nil
	}

	//申请数据入库
	err = imdb.AgentApply(&db.TAgentApplyRecord{
		UserId:      req.UserId,
		ChessUserId: req.ChessUserId,
		Name:        req.Name,
		Mobile:      req.Mobile,
	})

	if err != nil {
		log.Error(req.OperationId, "推广员申请提交数据入库失败:%s", err.Error())
		resp.CommonResp.Msg = "申请数据保存失败"
		resp.CommonResp.Code = 400
	}

	return resp, nil
}

// 绑定推广员
func (rpc *AgentServer) BindAgentNumber(_ context.Context, req *agent.BindAgentNumberReq) (*agent.BindAgentNumberResp, error) {
	resp := &agent.BindAgentNumberResp{CommonResp: &agent.CommonResp{Code: 0, Msg: ""}}

	//查询推广员是否存在
	agentInfo, _ := imdb.GetAgentByAgentNumber(req.AgentNumber)
	if agentInfo == nil || agentInfo.OpenStatus == 0 {
		resp.CommonResp.Msg = "请输入正确的推广员ID"
		resp.CommonResp.Code = 400
		return resp, nil
	}

	//绑定推广员
	err := imdb.BindAgentNumber(&db.TAgentMember{
		UserId:        req.UserId,
		AgentNumber:   req.AgentNumber,
		ChessUserId:   req.ChessUserId,
		ChessNickname: req.ChessNickname,
	})

	if err != nil {
		log.Error(req.OperationId, "绑定推广员数据入库失败:%s", err.Error())
		resp.CommonResp.Msg = "绑定推广员失败"
		resp.CommonResp.Code = 400
	}

	return resp, nil
}

// 获取当前用户的推广员信息以及绑定关系
func (rpc *AgentServer) GetUserAgentInfo(_ context.Context, req *agent.GetUserAgentInfoReq) (*agent.GetUserAgentInfoResp, error) {
	resp := &agent.GetUserAgentInfoResp{
		IsAgent:         false,
		AgentNumber:     0,
		AgentName:       "",
		BindAgentNumber: 0,
	}

	//是否为推广员
	info, _ := imdb.GetAgentByChessUserId(req.ChessUserId)
	if info != nil {
		resp.IsAgent = true
		resp.AgentName = info.Name
		resp.AgentNumber = info.AgentNumber
	} else {
		//是否申请
		applyInfo, _ := imdb.ApplyInfo(req.ChessUserId)
		if applyInfo != nil {
			resp.IsApply = true
		}
	}

	//是否绑定推广员
	agentMember, _ := imdb.AgentNumberByChessUserId(req.ChessUserId)
	if agentMember != nil {
		resp.BindAgentNumber = agentMember.AgentNumber
	}

	return resp, nil
}

// 推广员主页信息
func (rpc *AgentServer) AgentMainInfo(_ context.Context, req *agent.AgentMainInfoReq) (*agent.AgentMainInfoResp, error) {
	resp := &agent.AgentMainInfoResp{}

	//获取推广员信息
	info, _ := imdb.GetAgentByUserId(req.UserId)
	if info == nil {
		return resp, nil
	}

	resp.AgentNumber = info.AgentNumber
	resp.AgentName = info.Name
	resp.Balance = info.Balance
	resp.BeanBalance = info.BeanBalance

	//累计收益、今日收益
	statIncome, _ := imdb.StatAgentIncomeData(req.UserId)
	resp.TodayIncome = statIncome.TodayIncome
	resp.AccumulatedIncome = statIncome.AccumulatedIncome

	//绑定的下属成员
	statMember, _ := imdb.StatAgentMemberData(req.UserId)
	resp.TodayBindUser = statMember.TodayBindUser
	resp.AccumulatedBindUser = statMember.AccumulatedBindUser
	return resp, nil
}

// 账户明细收益趋势图
func (rpc *AgentServer) AgentAccountIncomeChart(_ context.Context, req *agent.AgentAccountIncomeChartReq) (*agent.AgentAccountIncomeChartResp, error) {
	resp := &agent.AgentAccountIncomeChartResp{}

	//获取收益统计数据
	chartData, _ := imdb.AccountIncomeChart(req.UserId, req.DateType)
	if len(chartData) > 0 {
		resp.IncomeChartData = make([]*agent.IncomeChartData, 0)
		for _, v := range chartData {
			resp.IncomeChartData = append(resp.IncomeChartData, &agent.IncomeChartData{
				Date:   v.Date,
				Income: v.Income,
			})
		}
	}

	return resp, nil
}

// 账户明细详情列表
func (rpc *AgentServer) AgentAccountRecordList(_ context.Context, req *agent.AgentAccountRecordListReq) (*agent.AgentAccountRecordListResp, error) {
	resp := &agent.AgentAccountRecordListResp{Total: 0, AccountRecordList: []*agent.AccountRecordList{}}

	//搜索用户
	chessUserIds := make([]int64, 0)
	if len(req.Keyword) > 0 {
		chessUserIds, _ = imdb.FindAgentMemberByKey(req.UserId, req.Keyword)
		if len(chessUserIds) == 0 {
			return resp, nil
		}
	}

	//获取收益统计数据
	list, count, _ := imdb.AccountIncomeList(req.UserId, req.Date, req.BusinessType, req.Page, req.Size, chessUserIds)
	resp.Total = count
	if len(list) > 0 {
		for _, v := range list {
			resp.AccountRecordList = append(resp.AccountRecordList, &agent.AccountRecordList{
				BusinessType: v.BusinessType,
				Amount:       v.Amount,
				RebateAmount: v.RebateAmount,
				Describe:     v.Describe,
				CreatedTime:  v.CreatedTime.Format("2006-01-02 15:04:05"),
				Type:         v.Type,
			})
		}
	}

	return resp, nil
}
