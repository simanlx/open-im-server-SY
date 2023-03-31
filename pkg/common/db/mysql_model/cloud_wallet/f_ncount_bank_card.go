package cloud_wallet

import "Open_IM/pkg/common/db"

// 绑定银行卡
func CreateNcountBankCard(info *db.FNcountBankCard) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Create(&info).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新信息
func UpdateNcountBankCardField(id int32, m map[string]interface{}) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Where("id = ?", id).Updates(m).Error
	return err
}

// 获取绑定的银行卡信息
func GetNcountBankCardById(id int32) (info *db.FNcountBankCard, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Where("id = ?", id).First(&info).Error
	if err != nil {
		return nil, err
	}
	return
}

// 通过UserID 和用户的账户类型 查询到用户的银行卡信息
func GetNcountBankCardByUserIdAndType(userId string, accountType int) (info *db.FNcountBankCard, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_bank_card").Where("user_id = ? ", userId).First(&info).Error
	if err != nil {
		return nil, err
	}
	return
}
