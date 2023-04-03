package cloud_wallet

import (
	"Open_IM/pkg/common/db"
	"github.com/pkg/errors"
)

/*
CREATE TABLE `f_ncount_account` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) NOT NULL COMMENT '用户id',
  `main_account_id` varchar(20) DEFAULT NULL COMMENT '新生支付主账号id',
  `packet_account_id` varchar(20) DEFAULT NULL COMMENT '新生支付红包账户id',
  `mobile` varchar(15) NOT NULL COMMENT '手机号码',
  `realname` varchar(20) NOT NULL COMMENT '真实姓名',
  `id_card` varchar(30) NOT NULL COMMENT '身份证',
  `pay_switch` tinyint(4) DEFAULT '1' COMMENT '支付开关(0关闭、1默认开启)',
  `bod_pay_switch` tinyint(4) DEFAULT '0' COMMENT '指纹支付/人脸支付开关(0默认关闭、1开启)',
  `payment_password` varchar(32) DEFAULT NULL COMMENT '支付密码(md5加密)',
  `open_status` tinyint(4) DEFAULT '0' COMMENT '开通状态',
  `open_step` tinyint(4) DEFAULT '1' COMMENT '开通认证步骤(1身份证认证、2支付密码、3绑定银行卡或者刷脸)',
  `created_time` datetime DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='云钱包账户表';
*/

type FNcountAccount struct {
	Id              int32  `gorm:"column:id;type:int(10) unsigned;not null;primary_key;auto_increment;comment:'主键'" json:"id"`
	UserId          int32  `gorm:"column:user_id;type:varchar(64);not null;comment:'用户id'" json:"userId"`
	MainAccountId   string `gorm:"column:main_account_id;type:varchar(20);default:null;comment:'新生支付主账号id'" json:"mainAccountId"`
	PacketAccountId string `gorm:"column:packet_account_id;type:varchar(20);default:null;comment:'新生支付红包账户id'" json:"packetAccountId"`
	Mobile          string `gorm:"column:mobile;type:varchar(15);not null;comment:'手机号码'" json:"mobile"`
	RealName        string `gorm:"column:realname;type:varchar(20);not null;comment:'真实姓名'" json:"realName"`
	IdCard          string `gorm:"column:id_card;type:varchar(30);not null;comment:'身份证'" json:"idCard"`
	PaySwitch       int32  `gorm:"column:pay_switch;type:tinyint(4);default:1;comment:'支付开关(0关闭、1默认开启)'" json:"paySwitch"`
	BodPaySwitch    int32  `gorm:"column:bod_pay_switch;type:tinyint(4);default:0;comment:'指纹支付/人脸支付开关(0默认关闭、1开启)'" json:"bodPaySwitch"`
	PaymentPassword string `gorm:"column:payment_password;type:varchar(32);default:null;comment:'支付密码(md5加密)'" json:"paymentPassword"`
	OpenStatus      int32  `gorm:"column:open_status;type:tinyint(4);default:0;comment:'开通状态'" json:"openStatus"`
	OpenStep        int32  `gorm:"column:open_step;type:tinyint(4);default:1;comment:'开通认证步骤(1身份证认证、2支付密码、3绑定银行卡或者刷脸)'" json:"openStep"`
	CreatedTime     string `gorm:"column:created_time;type:datetime;default:null" json:"createdTime"`
	UpdatedTime     string `gorm:"column:updated_time;type:datetime;default:null" json:"updatedTime"`
}

func FNcountAccountGetUserAccountID(userId string) (*db.FNcountAccount, error) {
	var account *db.FNcountAccount
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").
		Select("main_account_id", "packet_account_id").Where("user_id = ?", userId).First(account).Error
	if err != nil {
		return nil, errors.Wrap(err, "FNcountAccountGetUserAccountID error")
	}
	return account, nil
}

// 获取用户账户信息
func GetNcountAccountByUserId(userID string) (info *db.FNcountAccount, err error) {
	err = db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Where("user_id = ?", userID).First(&info).Error
	if err != nil {
		return nil, err
	}
	return
}

// 创建云钱包账户
func CreateNcountAccount(info *db.FNcountAccount) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Create(&info).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新账户信息
func UpdateNcountAccountField(userId string, m map[string]interface{}) error {
	err := db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Where("user_id = ?", userId).Updates(m).Error
	return err
}

func UpdateNcountAccountInfo(info *db.FNcountAccount) error {
	return db.DB.MysqlDB.DefaultGormDB().Table("f_ncount_account").Where("user_id = ?", info.UserId).Updates(&info).Error
}
