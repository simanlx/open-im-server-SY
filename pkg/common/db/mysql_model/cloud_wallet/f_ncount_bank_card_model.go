package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"time"
)

// 获取银行卡列表
func GetUserBankcardByUserId(userID string) (list []*db.FNcountBankCard, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Where("user_id = ? and is_bind = ?", userID, 1).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

// 绑定用户银行卡
func BindUserBankcard(info *db.FNcountBankCard) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Create(info).Error
	if err != nil {
		return err
	}
	return nil
}

// 绑定用户银行卡确认
func BindUserBankcardConfirm(bankcardId int32, userId string, ncountOrderId string) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Where("id = ? and user_id = ?", bankcardId, userId).Updates(map[string]interface{}{
		"ncount_order_id": ncountOrderId, "is_bind": 1, "updated_time": time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// 解绑用户银行卡
func UnBindUserBankcard(bankcardId int32, userId string) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Where("id = ? and user_id = ?", bankcardId, userId).Updates(map[string]interface{}{
		"is_delete": 1, "is_bind": 0, "updated_time": time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}
