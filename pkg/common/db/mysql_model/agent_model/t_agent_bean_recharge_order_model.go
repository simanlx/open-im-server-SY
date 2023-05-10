package agent_model

import (
	"Open_IM/pkg/common/db"
	"time"
)

const (
	RechargeOrderBusinessTypeWeb   = 1 //h5
	RechargeOrderBusinessTypeChess = 2 //互娱
)

// 创建购买咖豆订单
func CreatePurchaseBeanOrder(info *db.TAgentBeanRechargeOrder) error {
	info.CreatedTime = time.Now()
	info.UpdatedTime = time.Now()
	err := db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_bean_recharge_order").Create(info).Error
	return err
}
