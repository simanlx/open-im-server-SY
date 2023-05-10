package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"time"
)

// 第三方支付(创建一个第三方登陆)
func InsertThirdPayOrder(tp *db.ThirdPayOrder) error {
	tp.AddTime = time.Now()
	tp.EditTime = time.Now()
	result := db.DB.MysqlDB.DefaultGormDB().Table("third_pay_order").Create(tp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 修改第三方支付
func UpdateThirdPayOrder(tp *db.ThirdPayOrder, tpID int) error {
	result := db.DB.MysqlDB.DefaultGormDB().Table("third_pay_order").Where("id = ?", tpID).Updates(tp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 查询第三方支付订单
func GetThirdPayOrder(MerOrderNo string) (error, *db.ThirdPayOrder) {
	resp := &db.ThirdPayOrder{}
	result := db.DB.MysqlDB.DefaultGormDB().Table("third_pay_order").Where("order_no = ?", MerOrderNo).Find(resp)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, resp
}
