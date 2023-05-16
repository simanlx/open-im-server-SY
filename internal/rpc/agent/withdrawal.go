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
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

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

	orderNo := utils.GetOrderNo() //平台订单号
	commission, commissionFee := computeWithdrawalCommissionFee(req.Amount)

	//提现申请通知
	go utils.WithdrawApplyNotify(agentInfo.AgentNumber, req.Amount, agentInfo.Balance, commission, commissionFee)

	info := &db.TAgentWithdraw{
		OrderNo:         orderNo,
		NcountOrderNo:   "",
		UserId:          agentInfo.UserId,
		AgentNumber:     agentInfo.AgentNumber,
		BeforeBalance:   agentInfo.Balance,
		Balance:         req.Amount,
		NcountBalance:   0,
		TransferredTime: 0,
		Commission:      commission,
		CommissionFee:   commissionFee,
		Status:          0,
		CreatedTime:     time.Now(),
		UpdatedTime:     time.Now(),
	}

	//处理推广员余额提现逻辑
	err = BalanceWithdrawalSubmitLogic(ctx, info, req.OperationId)
	if err != nil {
		log.Error(req.OperationId, fmt.Sprintf("处理推广员余额提现逻辑失败,推广员id(%s),err:%s", req.UserId, err.Error()))
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = err.Error()
		return resp, nil
	}

	return resp, nil
}

// 处理推广员余额提现逻辑
func BalanceWithdrawalSubmitLogic(ctx context.Context, info *db.TAgentWithdraw, operationId string) error {
	//开启事务
	tx := db.DB.AgentMysqlDB.DefaultGormDB().Begin()

	// 1、写入提现记录数据
	err := tx.Table("t_agent_withdraw").Create(&info).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "写入提现记录数据失败")
	}

	// 2、冻结推广员余额
	err = tx.Table("t_agent_account").Where("user_id = ? and balance >= ?", info.UserId, info.Balance).UpdateColumn("balance", gorm.Expr(" balance - ? ", info.Balance)).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("修改推广员(%s)余额失败,余额(%d),提现余额(%d)", info.UserId, info.BeforeBalance, info.Balance))
	}

	//3、增加余额变更日志
	balanceRecord := &db.TAgentAccountRecord{
		OrderNo:           info.OrderNo,
		UserId:            info.UserId,
		Type:              2,
		BusinessType:      imdb.AccountBusinessTypeWithdraw,
		ChessUserId:       0,
		ChessUserNickname: "",
		Describe:          "提现到银行卡",
		Amount:            info.Balance,
		Day:               time.Now().Format("2006-01-02"),
		Month:             time.Now().Format("2006-01"),
		CreatedTime:       time.Now(),
		UpdatedTime:       time.Now(),
		DB:                tx,
	}
	err = tx.Table("t_agent_account_record").Create(&balanceRecord).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "增加余额变更日志失败")
	}

	//4、rpc 调用新生支付提现接口-到主账户余额
	_, err = RpcBalanceWithdrawal(ctx, info.OrderNo, info.Balance, operationId)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "新生支付提现接失败")
	}

	ncountOrderNo := ""

	//5、更新提现记录
	err = tx.Table("t_agent_withdraw").Where("id = ?", info.Id).Updates(map[string]interface{}{
		"ncount_order_no":  ncountOrderNo,
		"status":           1,
		"transferred_time": time.Now().Unix(),
		"updated_time":     time.Now(),
	}).Error
	if err != nil {
		//只记录错误日志，不回滚
		log.NewError(operationId, "订单号:", info.OrderNo, ",新生支付订单号", ncountOrderNo, ",更新提现记录 error:", err.Error())
	}

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		log.NewError(operationId, "BalanceWithdrawalSubmitLogic commit error:", err, "tx.Rollback().Error :", tx.Rollback().Error)
		return errors.Wrap(err, "事务提交失败")
	}

	return nil
}

// 计算提现手续费
func computeWithdrawalCommissionFee(amount int32) (int64, int64) {
	//提现手续费比例、千分比
	withdrawalCommission := rocksCache.GetPlatformValueConfigCache("withdrawal_commission") //获取提现手续费比例(‰)千分之几
	withdrawalCommissionDecimal, _ := decimal.NewFromString(withdrawalCommission)
	if withdrawalCommissionDecimal.IsZero() {
		return 0, 0
	}

	// (提现金额 * 提现手续费比例) / 千分比
	amountDecimal := decimal.NewFromInt32(amount)
	feeAmount, _ := amountDecimal.Mul(withdrawalCommissionDecimal).Div(decimal.NewFromInt(1000)).RoundFloor(0).Float64() //(a * b) / 1000 取整
	log.Info("", "计算提现手续费 computeWithdrawalCommissionFee,提现金额,提现手续比例,提现手续费 ", amount, withdrawalCommission, feeAmount)

	return cast.ToInt64(withdrawalCommission), cast.ToInt64(feeAmount)
}

// rpc 调用新生支付提现接口-到主账户余额
func RpcBalanceWithdrawal(ctx context.Context, orderNo string, amount int32, operationID string) (string, error) {
	return "xdfwf", nil
	//etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImCloudWalletName, operationID)
	//if etcdConn == nil {
	//	errMsg := operationID + "getcdv3.GetDefaultConn CreateThirdPayOrder == nil"
	//	log.NewError(operationID, errMsg)
	//	return "", errors.New(errMsg)
	//}
	//
	////组装数据
	//rpcReq := rpc.CreateThirdPayOrderReq{
	//	MerchantId:  config.Config.Agent.MerchantId, //商户号
	//	MerOrderId:  orderNo,
	//	NotifyUrl:   config.Config.Agent.AgentRechargeNotifyUrl,
	//	Amount:      amount,
	//	Remark:      "推广员充值咖豆",
	//	OperationID: operationID,
	//}
	//
	//client := rpc.NewCloudWalletServiceClient(etcdConn)
	//RpcResp, _ := client.CreateThirdPayOrder(ctx, &rpcReq)
	//if RpcResp.CommonResp != nil && RpcResp.CommonResp.ErrCode != 0 {
	//	log.NewError(operationID, "client.CreateThirdPayOrder 调用失败:", RpcResp.CommonResp.ErrMsg)
	//	return "", errors.New(RpcResp.CommonResp.ErrMsg)
	//}
	//
	//return RpcResp.OrderNo, nil
}
