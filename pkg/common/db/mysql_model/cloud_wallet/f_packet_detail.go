package cloud_wallet

import "Open_IM/pkg/common/db"

/*
CREATE TABLE `f_packet_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `packet_id` varchar(255) NOT NULL COMMENT '红包id',
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户id',
  `mer_order_id` varchar(255) DEFAULT NULL COMMENT '转账的商户id',
  `amount` int(11) NOT NULL COMMENT '领取金额分为单位',
  `receive_time` int(11) DEFAULT NULL COMMENT '领取时间',
  `created_time` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_time` int(11) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_packet_id` (`packet_id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='红包领取记录';
*/

// 红包的领取记录，提供给用户查询整年的红包领取记录
type FPacketDetail struct {
	ID          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	PacketID    string `gorm:"column:packet_id;not null" json:"packet_id"`
	UserID      string `gorm:"column:user_id" json:"user_id"`
	MerOrderID  string `gorm:"column:mer_order_id" json:"mer_order_id"`
	Amount      int64  `gorm:"column:amount;not null" json:"amount"`
	ReceiveTime int64  `gorm:"column:receive_time" json:"receive_time"`
	CreatedTime int64  `gorm:"column:created_time" json:"created_time"`
	UpdatedTime int64  `gorm:"column:updated_time" json:"updated_time"`
}

func RedPacketDetailCreateData(req *FPacketDetail) error {
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_packet_detail").Create(req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 查询用户的发送红包记录
func FPacketDetailGetByPacketID(packetID, userID string) (*FPacketDetail, error) {
	var res = FPacketDetail{}
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_packet_detail").Where("packet_id = ? and user_id= ?", packetID, userID).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return &res, nil
}

// 保存用户领取红包记录
func InsertRedPacketDetail(req *FPacketDetail) error {
	result := db.DB.MysqlDB.DefaultGormDB().Table("f_packet_detail").Save(req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
