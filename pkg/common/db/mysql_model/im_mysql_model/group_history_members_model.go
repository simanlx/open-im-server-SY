package im_mysql_model

import (
	"Open_IM/pkg/common/db"
	"time"
)

// 群成员数据入库
func InsertGroupHistoryMembers(info db.GroupHistoryMembers) (err error) {
	info.LastSendMsgTime = 0
	info.CreatedTime = time.Now()
	err = db.DB.MysqlDB.DefaultGormDB().Table("group_history_members").Create(info).Error
	return
}

// 更新最后发送群消息时间
func UpGroupMembersLastSendMsgTime(groupId, userId string) (err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("group_history_members").Where("group_id = ? and user_id = ?", groupId, userId).Update("last_send_msg_time", time.Now().Unix()).Error
	return
}
