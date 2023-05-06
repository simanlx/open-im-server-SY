package agent

import (
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/proto/agent"
	"context"
	"fmt"
)

// 获取平台咖豆商城配置
func (rpc *AgentServer) PlatformBeanShopConfig(_ context.Context, req *agent.PlatformBeanShopConfigReq) (*agent.PlatformBeanShopConfigResp, error) {
	resp := &agent.PlatformBeanShopConfigResp{}

	//获取redis缓存配置
	beanConfig, err := rocksCache.GetAgentPlatformBeanConfigCache()
	fmt.Println("--GetAgentPlatformBeanConfigCache", beanConfig, err)
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

// 推广员咖豆商城配置
func (rpc *AgentServer) AgentBeanShopConfig(_ context.Context, req *agent.AgentBeanShopConfigReq) (*agent.AgentBeanShopConfigResp, error) {
	resp := &agent.AgentBeanShopConfigResp{}

	return resp, nil
}
