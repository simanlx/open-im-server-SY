package agent

import (
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/agent_model"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/utils"
	"Open_IM/pkg/proto/agent"
	"context"
	"fmt"
)

// 互娱商城购买咖豆下单(预提交)
func (rpc *AgentServer) ChessShopPurchaseBean(ctx context.Context, req *agent.ChessShopPurchaseBeanReq) (*agent.ChessShopPurchaseBeanResp, error) {
	resp := &agent.ChessShopPurchaseBeanResp{OrderNo: "", CommonResp: &agent.CommonResp{Code: 0, Msg: ""}}

	// 加锁
	lockKey := fmt.Sprintf("ChessShopPurchaseBean:%d", req.ChessUserId)
	if err := utils.Lock(ctx, lockKey); err != nil {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "操作加锁失败"
		return resp, nil
	}
	defer utils.UnLock(ctx, lockKey)

	//获取推广员信息
	agentInfo, err := imdb.GetAgentByUserId(req.UserId)
	if err != nil || agentInfo.OpenStatus == 0 {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "推广员信息有误"
		return resp, nil
	}

	//校验咖豆配置
	configInfo, err := imdb.GetAgentBeanConfigById(req.UserId, req.ConfigId)
	if err != nil || configInfo.Status == 0 {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "咖豆配置有误"
		return resp, nil
	}

	//是否为下属成员
	agentMember, err := imdb.AgentNumberByChessUserId(req.ChessUserId)
	if err != nil || req.UserId != agentMember.UserId {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "该用户不是推广员下成员"
		return resp, nil
	}

	//冻结的咖豆额度
	freezeBeanBalance := rocksCache.GetAgentFreezeBeanBalance(ctx, agentInfo.UserId)
	//校验推广员咖豆余额 + 冻结部分
	if agentInfo.BeanBalance < (configInfo.BeanNumber + freezeBeanBalance) {
		log.Error(req.OperationId, fmt.Sprintf("推广员(%d),下属成员(%d)购买咖豆,推广员咖豆余额不足,咖豆余额(%d),冻结咖豆(%d)", agentInfo.AgentNumber, req.ChessUserId, agentInfo.BeanBalance, freezeBeanBalance))
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "推广员咖豆余额不足"
		return resp, nil
	}

	orderNo := utils.GetOrderNo() //平台订单号
	//生成订单
	err = imdb.CreatePurchaseBeanOrder(&db.TAgentBeanRechargeOrder{
		BusinessType:      imdb.RechargeOrderBusinessTypeChess,
		UserId:            req.UserId,
		ChessUserId:       req.ChessUserId,
		ChessUserNickname: agentMember.ChessNickname,
		OrderNo:           orderNo,
		ChessOrderNo:      req.ChessOrderNo,
		Number:            configInfo.BeanNumber,
		GiveNumber:        configInfo.GiveBeanNumber,
		Amount:            configInfo.Amount,
	})

	if err != nil {
		log.Error(req.OperationId, "互娱商城购买咖豆下单(预提交) 生成订单失败。互娱订单号：", req.ChessOrderNo, ",err:", err.Error())
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "生成订单失败"
		return resp, nil
	}

	//冻结推广员咖豆
	_ = rocksCache.FreezeAgentBeanBalance(ctx, agentInfo.UserId, req.ChessUserId, configInfo.BeanNumber)

	resp.OrderNo = orderNo
	resp.GiveBeanNumber = configInfo.GiveBeanNumber
	resp.BeanNumber = configInfo.BeanNumber
	resp.ConfigId = configInfo.Id
	resp.Amount = configInfo.Amount
	return resp, nil
}
