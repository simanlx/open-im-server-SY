package agent

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/agent_model"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/utils"
	"Open_IM/pkg/proto/agent"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
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
	agentInfo, err := imdb.GetAgentByAgentNumber(req.AgentNumber)
	if err != nil || agentInfo.OpenStatus == 0 {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "推广员信息有误"
		return resp, nil
	}

	//校验咖豆配置
	configInfo, err := imdb.GetAgentBeanConfigById(agentInfo.UserId, req.ConfigId)
	if err != nil || configInfo.Status == 0 {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "咖豆配置有误"
		return resp, nil
	}

	//是否为下属成员
	agentMember, err := imdb.AgentNumberByChessUserId(req.ChessUserId)
	if err != nil || agentInfo.UserId != agentMember.UserId {
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
		UserId:            agentInfo.UserId,
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

// 推广员购买咖豆
func (rpc *AgentServer) AgentPurchaseBean(ctx context.Context, req *agent.AgentPurchaseBeanReq) (*agent.AgentPurchaseBeanResp, error) {
	resp := &agent.AgentPurchaseBeanResp{CommonResp: &agent.CommonResp{Code: 0, Msg: ""}}

	configInfo, err := GetPlatformBeanConfigInfo(req.ConfigId)
	if err != nil || configInfo == nil {
		log.Error(req.OperationId, "获取平台咖豆商城配置缓存-GetAgentPlatformBeanConfigCache err :", err)
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "获取咖豆配置失败"
		return resp, nil
	}

	orderNo := utils.GetOrderNo() //平台订单号
	//生成订单
	err = imdb.CreatePurchaseBeanOrder(&db.TAgentBeanRechargeOrder{
		BusinessType:      imdb.RechargeOrderBusinessTypeWeb,
		UserId:            req.UserId,
		ChessUserId:       0,
		ChessUserNickname: "",
		OrderNo:           orderNo,
		ChessOrderNo:      "",
		Number:            configInfo.BeanNumber,
		GiveNumber:        configInfo.GiveBeanNumber,
		Amount:            configInfo.Amount,
	})

	if err != nil {
		log.Error(req.OperationId, "推广员购买咖豆下单，生成订单失败err:", err.Error())
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "生成订单失败"
		return resp, nil
	}

	//组装数据下发
	resp.OrderNo = orderNo
	resp.Amount = configInfo.Amount
	resp.NotifyUrl = config.Config.Ncount.Notify.AgentRechargeNotifyUrl

	return resp, nil
}

// 获取平台咖豆配置项
func GetPlatformBeanConfigInfo(configId int32) (*imdb.BeanShopConfig, error) {
	//获取平台咖豆redis缓存配置
	beanConfig, err := rocksCache.GetAgentPlatformBeanConfigCache()
	if err != nil || len(beanConfig) == 0 {
		return nil, errors.Wrap(err, "获取平台咖豆redis缓存配置失败")
	}

	for _, v := range beanConfig {
		if v.ConfigId == configId {
			return v, nil
		}
	}

	return nil, errors.Wrap(err, "获取平台咖豆redis缓存配置失败.")
}

// 推广员余额提现
func (rpc *AgentServer) BalanceWithdrawal(ctx context.Context, req *agent.BalanceWithdrawalReq) (*agent.BalanceWithdrawalResp, error) {
	resp := &agent.BalanceWithdrawalResp{CommonResp: &agent.CommonResp{Code: 0, Msg: ""}}

	// 加锁
	lockKey := fmt.Sprintf("BalanceWithdrawal:%s", req.UserId)
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
		resp.CommonResp.Msg = "信息有误"
		return resp, nil
	}

	//校验推广员余额
	if int64(req.Amount) > agentInfo.Balance {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "账户余额不足"
		return resp, nil
	}

	orderNo := utils.GetOrderNo()                                                 //平台订单号
	commission := rocksCache.GetPlatformValueConfigCache("withdrawal_commission") //获取提现手续费

	info := &db.TAgentWithdraw{
		OrderNo:             orderNo,
		NcountOrderNo:       "",
		UserId:              agentInfo.UserId,
		AgentNumber:         agentInfo.AgentNumber,
		BeforeBalance:       agentInfo.Balance,
		BeforeFreezeBalance: agentInfo.FreezeBalance,
		Balance:             req.Amount,
		NcountBalance:       0,
		Commission:          commission,
		CreatedTime:         time.Now(),
		UpdatedTime:         time.Now(),
	}

	//处理推广员余额提现逻辑
	err = BalanceWithdrawalSubmitLogic(info)
	if err != nil {
		log.Error("", fmt.Sprintf("处理推广员余额提现逻辑失败,推广员id(%s),err:%s", req.UserId, err.Error()))
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = err.Error()
		return resp, nil
	}

	return resp, nil
}

// 处理推广员余额提现逻辑
func BalanceWithdrawalSubmitLogic(info *db.TAgentWithdraw) error {
	//开启事务
	tx := db.DB.AgentMysqlDB.DefaultGormDB().Begin()

	// 1、写入提现记录数据
	err := tx.Table("t_agent_withdraw").Create(&info).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "写入提现记录数据失败")
	}

	// 2、冻结推广员余额
	err = tx.Table("t_agent_account").Where("user_id = ? and balance >= ?", info.UserId, info.Balance).UpdateColumns(map[string]interface{}{
		"balance":        gorm.Expr(" balance - ? ", info.Balance),
		"freeze_balance": gorm.Expr(" freeze_balance + ? ", info.Balance),
	}).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("修改推广员(%s)余额失败,余额(%d),冻结余额(%d),提现余额(%d)", info.UserId, info.BeforeBalance, info.BeforeFreezeBalance, info.Balance))
	}

	//3、调用rpc提现接口

	//notifyUrl := config.Config.Ncount.Notify.AgentWithdrawNotifyUrl //回调地址

	//调用rpc提现接口

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		log.NewError("", "BalanceWithdrawalSubmitLogic commit error:", err, "tx.Rollback().Error :", tx.Rollback().Error)
		return errors.Wrap(err, "事务提交失败")
	}

	return nil
}
