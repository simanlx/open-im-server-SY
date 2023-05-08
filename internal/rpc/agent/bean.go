package agent

import (
	imdb "Open_IM/pkg/common/db/mysql_model/agent_model"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/proto/agent"
	"context"
)

// 获取平台咖豆商城配置
func (rpc *AgentServer) PlatformBeanShopConfig(_ context.Context, req *agent.PlatformBeanShopConfigReq) (*agent.PlatformBeanShopConfigResp, error) {
	resp := &agent.PlatformBeanShopConfigResp{}

	//获取平台咖豆redis缓存配置
	beanConfig, err := rocksCache.GetAgentPlatformBeanConfigCache()
	if err != nil {
		log.Error(req.OperationId, "获取平台咖豆商城配置缓存-GetAgentPlatformBeanConfigCache err :", err.Error())
		return resp, nil
	}

	resp.BeanShopConfig = make([]*agent.BeanShopConfig, 0)
	for _, v := range beanConfig {
		resp.BeanShopConfig = append(resp.BeanShopConfig, &agent.BeanShopConfig{
			ConfigId:       v.ConfigId,
			BeanNumber:     v.BeanNumber,
			GiveBeanNumber: v.GiveBeanNumber,
			Amount:         v.Amount,
			Status:         1,
		})
	}

	return resp, nil
}

// 推广员游戏咖豆商城配置
func (rpc *AgentServer) AgentGameBeanShopConfig(_ context.Context, req *agent.AgentGameBeanShopConfigReq) (*agent.AgentGameBeanShopConfigResp, error) {
	resp := &agent.AgentGameBeanShopConfigResp{}

	//获取推广员信息
	agentInfo, _ := imdb.GetAgentByAgentNumber(req.AgentNumber)
	if agentInfo == nil || agentInfo.OpenStatus == 0 {
		return resp, nil
	}

	//获取推广员自定义咖豆配置
	configList, _ := imdb.GetAgentDiyShopBeanOnlineConfig(agentInfo.UserId)
	if len(configList) > 0 {
		resp.BeanShopConfig = make([]*agent.BeanShopConfig, 0)
		for _, v := range configList {
			resp.BeanShopConfig = append(resp.BeanShopConfig, &agent.BeanShopConfig{
				ConfigId:       v.Id,
				BeanNumber:     v.BeanNumber,
				GiveBeanNumber: v.GiveBeanNumber,
				Amount:         v.Amount,
			})
		}
	}

	return resp, nil
}

// 推广员自定义咖豆商城配置
func (rpc *AgentServer) AgentDiyBeanShopConfig(_ context.Context, req *agent.AgentDiyBeanShopConfigReq) (*agent.AgentDiyBeanShopConfigResp, error) {
	resp := &agent.AgentDiyBeanShopConfigResp{BeanShopConfig: []*agent.BeanShopConfig{}, TodaySales: 0}

	//获取今日出售咖豆数
	resp.TodaySales = imdb.GetAgentTodaySalesNumber(req.UserId)

	//获取推广员自定义咖豆配置
	configList, _ := imdb.GetAgentDiyShopBeanConfig(req.UserId)
	if len(configList) > 0 {
		for _, v := range configList {
			resp.BeanShopConfig = append(resp.BeanShopConfig, &agent.BeanShopConfig{
				ConfigId:       v.Id,
				BeanNumber:     v.BeanNumber,
				GiveBeanNumber: v.GiveBeanNumber,
				Amount:         v.Amount,
				Status:         v.Status,
			})
		}
	}

	return resp, nil
}

// 咖豆账户明细详情列表
func (rpc *AgentServer) AgentBeanAccountRecordList(_ context.Context, req *agent.AgentBeanAccountRecordListReq) (*agent.AgentBeanAccountRecordListResp, error) {
	resp := &agent.AgentBeanAccountRecordListResp{BeanRecordList: []*agent.BeanRecordList{}, Total: 0}

	list, count, err := imdb.BeanAccountRecordList(req.UserId, req.Date, req.BusinessType, req.Page, req.Size)
	if err != nil {
		return resp, nil
	}

	resp.Total = count
	for _, v := range list {
		resp.BeanRecordList = append(resp.BeanRecordList, &agent.BeanRecordList{
			Type:         v.Type,
			BusinessType: v.BusinessType,
			Amount:       v.Amount,
			Number:       v.Number,
			GiveNumber:   v.GiveNumber,
			Describe:     v.Describe,
			CreatedTime:  v.CreatedTime.Format("2006-01-02 15:04:05"),
		})
	}

	return resp, nil
}
