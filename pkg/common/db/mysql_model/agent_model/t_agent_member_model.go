package agent_model

import (
	"Open_IM/pkg/common/db"
	"time"
)

// 绑定推广员
func BindAgentNumber(info *db.TAgentMember) error {
	info.CreatedTime = time.Now()
	info.UpdatedTime = time.Now()
	err := db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_member").Create(info).Error
	return err
}

// 推广员下属
func AgentNumberByChessUserId(chessUserId int64) (info *db.TAgentMember, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_member").Where("chess_user_id = ?", chessUserId).First(&info).Error
	return
}
