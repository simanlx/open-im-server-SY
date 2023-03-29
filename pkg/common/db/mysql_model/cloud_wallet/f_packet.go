package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"github.com/pkg/errors"
)

/*
CREATE TABLE `f_packet` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `packet_id` varchar(255) DEFAULT NULL COMMENT '红包ID',
  `user_id` int(11) NOT NULL COMMENT '红包发起者',
  `packet_type` tinyint(1) NOT NULL COMMENT '红包类型(1个人红包、2群红包)',
  `is_lucky` tinyint(1) DEFAULT '0' COMMENT '是否为拼手气红包',
  `exclusive_user_id` int(11) DEFAULT '0' COMMENT '专属用户id',
  `packet_title` varchar(100) NOT NULL COMMENT '红包标题',
    `amount` int(11) NOT NULL COMMENT '红包金额',
  `number` tinyint(3) NOT NULL COMMENT '红包个数',
  `expire_time` int(11) DEFAULT NULL COMMENT '红包过期时间',
  `created_time` int(11) DEFAULT NULL,
  `updated_time` int(11) DEFAULT NULL,
  `status` tinyint(1) NOT NULL COMMENT '红包状态： 1 为创建 、2 为正常、3为异常',
  `is_exclusive` tinyint(1) NOT NULL COMMENT '是否为专属红包： 0为否，1为是',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户红包表';
*/

type FPacket struct {
	ID              int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	PacketID        string `gorm:"column:packet_id;not null" json:"packet_id"`
	UserID          int64  `gorm:"column:user_id;not null" json:"user_id"`
	PacketType      int32  `gorm:"column:packet_type;not null" json:"packet_type"`
	IsLucky         int32  `gorm:"column:is_lucky;not null" json:"is_lucky"`
	ExclusiveUserID int64  `gorm:"column:exclusive_user_id;not null" json:"exclusive_user_id"`
	PacketTitle     string `gorm:"column:packet_title;not null" json:"packet_title"`
	Amount          int64  `gorm:"column:amount;not null" json:"amount"`
	Number          int32  `gorm:"column:number;not null" json:"number"`
	ExpireTime      int64  `gorm:"column:expire_time;not null" json:"expire_time"`
	CreatedTime     int64  `gorm:"column:created_time;not null" json:"created_time"`
	UpdatedTime     int64  `gorm:"column:updated_time;not null" json:"updated_time"`
	Status          int32  `gorm:"column:status;not null" json:"status"` // 0 创建未生效，1 为红包正在领取中，2为红包领取完毕，3为红包过期
	IsExclusive     int32  `gorm:"column:is_exclusive;not null" json:"is_exclusive"`
}

const (
	RedPacketStatusCreate   = iota // 0 创建未生效
	RedPacketStatusNormal          // 1 为红包正在领取中
	RedPacketStatusFinished        // 2为红包领取完毕
	RedPacketStatusExpired         // 3为红包过期
)

// 保存到红包到数据库
func RedPacketCreateData(req *FPacket) error {
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_packet").Create(req)
	if result.Error != nil {
		return errors.Wrap(result.Error, "创建红包失败")
	}
	return nil
}

// 修改红包状态
func UpdateRedPacketStatus(packetID string, status int64) error {
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_packet").Where("packet_id = ?", packetID).Update("status", status)
	if result.Error != nil {
		return errors.Wrap(result.Error, "修改红包状态失败")
	}
	return nil
}

func GetRedPacketInfo(packetID string) (*FPacket, error) {
	var fPacket FPacket
	result := db.DB.MysqlDB.DefaultGormDB().
		Table("f_packet").
		Select([]string{"is_lucky", "IsExclusive", "exclusive_user_id", "expire_time"}).
		Where("packet_id = ?", packetID).
		First(&fPacket)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "获取红包信息失败")
	}
	return &fPacket, nil
}
