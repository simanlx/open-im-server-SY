package agent_model

import "Open_IM/pkg/common/db"

// 获取推广员信息
func GetAgentByAgentNumber(agentNumber int32) (info *db.TAgentAccount, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_account").Where("agent_number = ?", agentNumber).First(&info).Error
	return
}

// 获取推广员信息
func GetAgentByChessUserId(chessUserId int64) (info *db.TAgentAccount, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_account").Where("chess_user_id = ?", chessUserId).First(&info).Error
	return
}
