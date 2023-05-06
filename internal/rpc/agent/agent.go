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

// 绑定推广员
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

// 绑定推广员
func (rpc *AgentServer) AgentMainInfo(_ context.Context, req *agent.AgentMainInfoReq) (*agent.AgentMainInfoResp, error) {
	resp := &agent.AgentMainInfoResp{
		AgentNumber:         1,
		AgentName:           "1",
		Balance:             1,
		BeanBalance:         1,
		TodayIncome:         1,
		AccumulatedIncome:   1,
		TodayBindUser:       1,
		AccumulatedBindUser: 1,
	}

	//获取推广员信息
	info, _ := imdb.GetAgentByUserId(req.UserId)
	if info == nil {
		return resp, nil
	}

	resp.AgentNumber = info.AgentNumber
	resp.AgentName = info.Name
	resp.Balance = info.Balance
	resp.BeanBalance = info.BeanBalance
	resp.AccumulatedIncome = info.AccumulatedIncome

	//今日收益=下属用户充值返利+商城出售咖豆收入

	//resp.TodayIncome = 0
	return resp, nil
}
