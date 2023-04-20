package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"time"
)

func CreateErrorLog(Remark, MerOrderId, ErrMsg, ErrCode, AllMsg string) error {
	log := &db.FErrorLog{
		Remark:     Remark,
		MerOrderId: MerOrderId,
		ErrMsg:     ErrMsg,
		ErrCode:    ErrCode,
		AllMsg:     AllMsg,
		CreateTime: time.Now().Unix(),
	}
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_error_log").Create(log).Error
	return err
}
