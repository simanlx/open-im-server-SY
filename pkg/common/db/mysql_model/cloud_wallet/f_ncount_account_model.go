package cloud_wallet

import "Open_IM/pkg/common/db"

// 获取用户账户信息
func GetNcountAccountByUserId(userID string) (info *db.FNcountAccount, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Where("user_id = ?", userID).First(info).Error
	if err != nil {
		return nil, err
	}
	return
}

// 创建云钱包账户
func CreateNcountAccount(info *db.FNcountAccount) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Create(&info).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新账户信息
func UpdateNcountAccountField(userId int32, m map[string]interface{}) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Where("user_id = ?", userId).Updates(m).Error
	return err
}

func UpdateNcountAccountInfo(info *db.FNcountAccount) error {
	return db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Where("user_id = ?", info.UserId).Updates(&info).Error
}
