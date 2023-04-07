package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"github.com/pkg/errors"
	"time"
)

const (
	TradeTypeCharge       = iota + 1 // 充值
	TradeTypeWithdraw                // 提现
	TradeTypeRedPacketOut            // 红包支出
	TradeTypeRedPacketIn             // 红包收入
	TradeTypeTransferOut             // 转账支出
	TradeTypeTransferIn              // 转账收入
	TradeTypeRefund                  // 退款
)

func FNcountTradeCreateData(req *db.FNcountTrade) error {
	req.CreatedTime = time.Now()
	req.UpdatedTime = time.Now()
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_trade").Create(req)
	if result.Error != nil {
		return errors.Wrap(result.Error, "创建交易记录失败")
	}
	return nil
}

// 修改交易的状态
func FNcountTradeUpdateStatusbyThirdOrderNo(req *db.FNcountTrade) error {
	// 查询这个红包
	res := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_trade").Where("third_order_no = ?", req.ThirdOrderNo).First(&req)
	if res.Error != nil {
		return errors.Wrap(res.Error, "查询交易记录失败")
	}
	// 修改红包状态
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_trade").Where("third_order_no = ?", req.ThirdOrderNo).Update("ncount_status", req.NcountStatus)
	if result.Error != nil {
		return errors.Wrap(result.Error, "修改交易的状态失败")
	}

	return nil
}

// 获取充值记录信息
func GetFNcountTradeByOrderNo(orderNo, userId string) (info *db.FNcountTrade, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_trade").Where("third_order_no = ? and user_id = ?", orderNo, userId).First(&info).Error
	if err != nil {
		return nil, err
	}
	return
}
