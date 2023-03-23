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
