package agent_model

import (
	"Open_IM/pkg/common/db"
	"fmt"
	"time"
)

type SAgentMemberData struct {
	TodayBindUser       int64 `json:"today_bind_user"`
	AccumulatedBindUser int64 `json:"accumulated_bind_user"`
}

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

// 统计推广员下属数据
func StatAgentMemberData(userId string) (data *SAgentMemberData, err error) {
	today := time.Now().Format("2006-01-02")
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_member").
		Select("count(1) accumulated_bind_user,sum(if(`day`=?,1,0)) today_bind_user", today).
		Where("user_id = ?", userId).Scan(&data).Error
	return
}

// 条件获取代理商下属成员ids
func FindAgentMemberIds(userId, keyword string) (chessUserIds []int64, err error) {
	model := db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_member").Where("user_id = ?", userId)

	if len(keyword) > 0 {
		model = model.Where("chess_user_id = ? or chess_nickname like ?", keyword, fmt.Sprintf("%%%s%%", keyword))
	}
	err = model.Pluck("chess_user_id", &chessUserIds).Error
	return
}

// 获取推广员下属成员列表
func FindAgentMemberList(userId, keyword, orderBy string, page, size int32) (list []*db.TAgentMember, count int64, err error) {
	model := db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_member").Where("user_id = ?", userId)

	if len(keyword) > 0 {
		model = model.Where("chess_user_id = ? or chess_nickname like ?", keyword, fmt.Sprintf("%%%s%%", keyword))
	}

	model = model.Count(&count).Limit(int(size)).Offset(int(size * (page - 1)))

	//排序
	if len(orderBy) > 0 {
		model = model.Order(orderBy)
	} else {
		model = model.Order("id desc")
	}

	err = model.Find(&list).Error
	return
}
