package cloud_wallet

import "Open_IM/pkg/common/db"

// 将表中blockword 所有字段查出来

func GetAllBlockword() ([]db.Blockword, error) {
	var blockword []db.Blockword
	err := db.DB.MysqlDB.DefaultGormDB().Table("blockword").Find(&blockword).Error
	return blockword, err
}
