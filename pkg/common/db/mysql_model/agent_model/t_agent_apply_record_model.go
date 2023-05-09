package agent_model

import (
	"Open_IM/pkg/common/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// 申请记录
func ApplyInfo(chessUserId int64) (info *db.TAgentApplyRecord, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_apply_record").Where("chess_user_id = ?", chessUserId).First(&info).Error
	if errors.Is(errors.Unwrap(err), gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "")
	}
	return
}

// 申请
func AgentApply(info *db.TAgentApplyRecord) error {
	info.CreatedTime = time.Now()
	info.UpdatedTime = time.Now()
	err := db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_apply_record").Create(info).Error
	return err
}
