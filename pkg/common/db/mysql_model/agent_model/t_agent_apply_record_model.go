package agent_model

import (
	"Open_IM/pkg/common/db"
	"time"
)

// 申请记录
func ApplyInfo(chessUserId int64) (info *db.TAgentApplyRecord, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_apply_record").Where("chess_user_id = ?", chessUserId).First(&info).Error
	return
}

// 申请
func AgentApply(info *db.TAgentApplyRecord) error {
	info.CreatedTime = time.Now()
	info.UpdatedTime = time.Now()
	err := db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_apply_record").Create(info).Error
	return err
}
