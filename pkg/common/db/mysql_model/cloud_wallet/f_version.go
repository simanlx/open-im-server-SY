package cloud_wallet

import "Open_IM/pkg/common/db"

// 获取APP的版本
func GetFVersion(versionCode string) (db.FVersion, error) {
	var fversion db.FVersion
	// 获取最新的版本
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_version").Where("version_code = ?", versionCode).First(&fversion)
	return fversion, result.Error
}

// 获取最新红包信息
func GetLastedFVersion() (db.FVersion, error) {
	var fversion db.FVersion
	// 获取最新的版本
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_version").Order("id desc").First(&fversion)
	return fversion, result.Error
}
