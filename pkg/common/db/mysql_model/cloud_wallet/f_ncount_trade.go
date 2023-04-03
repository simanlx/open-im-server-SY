package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"github.com/pkg/errors"
	"time"
)

/*

CREATE TABLE `f_ncount_trade` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `payment_platform` tinyint(1) NOT NULL COMMENT '支付平台(1云钱包、2支付宝、3微信、4银行卡)',
  `type` tinyint(1) NOT NULL COMMENT '类型(1充值、2提现、3红包支出)',
  `amount` int(11) NOT NULL COMMENT '变更金额',
  `befer_amount` int(11) DEFAULT NULL COMMENT '变更前金额',
  `after_amount` int(11) DEFAULT NULL COMMENT '变更后金额',
  `third_order_no` varchar(100) DEFAULT NULL COMMENT '第三方订单号',
  `ncount_status` int(11) DEFAULT NULL COMMENT '第三方的回调状态（0 ：未生效，1 生效）',
  `created_time` int(11) DEFAULT NULL,
  `updated_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户账户变更表';

*/

const (
	TradeTypeCharge       = iota + 1 // 充值
	TradeTypeWithdraw                // 提现
	TradeTypeRedPacketOut            // 红包支出
	TradeTypeRedPacketIn             // 红包收入
	TradeTypeTransferOut             // 转账支出
	TradeTypeTransferIn              // 转账收入
	TradeTypeRefund                  // 退款
)

type FNcountTrade struct {
	ID              int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserID          string `gorm:"column:user_id;not null" json:"user_id"`
	PaymentPlatform int32  `gorm:"column:payment_platform;not null" json:"payment_platform"`
	Type            int32  `gorm:"column:type;not null" json:"type"`
	Amount          int64  `gorm:"column:amount;not null" json:"amount"`
	BeferAmount     int64  `gorm:"column:befer_amount;not null" json:"befer_amount"`
	AfterAmount     int64  `gorm:"column:after_amount;not null" json:"after_amount"`
	ThirdOrderNo    string `gorm:"column:third_order_no;not null" json:"third_order_no"`
	NcountStatus    int32  `gorm:"column:ncount_status;not null" json:"ncount_status"`
	CreatedTime     int64  `gorm:"column:created_time;not null" json:"created_time"`
	UpdatedTime     int64  `gorm:"column:updated_time;not null" json:"updated_time"`
}

func FNcountTradeCreateData(req *FNcountTrade) error {
	req.CreatedTime = time.Now().Unix()
	req.UpdatedTime = time.Now().Unix()
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_trade").Create(req)
	if result.Error != nil {
		return errors.Wrap(result.Error, "创建红包交易记录失败")
	}
	return nil
}

// 修改交易的状态
func FNcountTradeUpdateStatusbyThirdOrderNo(req *FNcountTrade) error {
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_trade").Where("third_order_no = ?", req.ThirdOrderNo).Update("ncount_status", req.NcountStatus)
	if result.Error != nil {
		return errors.Wrap(result.Error, "修改交易的状态失败")
	}
	return nil
}
