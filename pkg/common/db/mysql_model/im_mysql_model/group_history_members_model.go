package im_mysql_model

import (
	"Open_IM/pkg/common/db"
	"time"
)

type GroupHistoryMembers struct {
	UserId          string    `json:"user_id"`            //用户id
	FaceUrl         string    `json:"face_url"`           //头像
	Nickname        string    `json:"nickname"`           //昵称
	LastSendMsgTime int32     `json:"last_send_msg_time"` //最后发送群消息时间
	CreatedTime     time.Time `json:"created_time"`       //加群时间
}

// 群成员数据入库
func InsertGroupHistoryMembers(info *db.GroupHistoryMembers) (err error) {
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

// 获取群历史成员列表
func FindGroupMembersList(groupId string, page, size int32) (list []*GroupHistoryMembers, count int64, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("group_history_members g").Select("g.user_id,g.last_send_msg_time,g.created_time,u.name nickname,u.face_url").
		Where(" g.group_id = ? ", groupId).
		Joins("left join users u on g.user_id = u.user_id").
		Count(&count).Limit(int(size)).Offset(int(size * (page - 1))).
		Order("g.id desc").Find(&list).Error
	return
}
