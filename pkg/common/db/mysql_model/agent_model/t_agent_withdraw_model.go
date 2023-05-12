package agent_model

import (
	"Open_IM/pkg/common/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 获取提现订单by order_no
func GetWithdrawOrderByOrderNo(orderNo string) (info *db.TAgentWithdraw, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_withdraw").Where("order_no = ?", orderNo).Take(&info).Error
	if errors.Is(errors.Unwrap(err), gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "")
	}
	return
}
