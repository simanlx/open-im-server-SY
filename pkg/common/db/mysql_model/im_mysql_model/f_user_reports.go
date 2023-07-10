package im_mysql_model

import (
	"Open_IM/pkg/common/db"
	"fmt"
	"time"
)

//CREATE TABLE `f_user_reports` (
//  `id` int(11) NOT NULL AUTO_INCREMENT,
//  `user_id` varchar(255) DEFAULT '' COMMENT '用户ID',
//  `latitude` varchar(255) DEFAULT '' COMMENT '纬度',
//  `longitude` varchar(255) DEFAULT '' COMMENT '经度',
//  `speed` varchar(255) DEFAULT '' COMMENT '速度',
//  `rotate_angle` varchar(255) DEFAULT '' COMMENT '角度',
//  `battery` int(11) DEFAULT '0' COMMENT '电池电量\n',
//  `step` int(11) DEFAULT '0' COMMENT '步数\n',
//  `add_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间\n',
//  `edit_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间\n',
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

type FUserReports struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId      string    `gorm:"column:user_id;type:varchar(255);default:'';comment:'用户ID'" json:"user_id"`
	Latitude    string    `gorm:"column:latitude;type:varchar(255);default:'';comment:'纬度'" json:"latitude"`
	Longitude   string    `gorm:"column:longitude;type:varchar(255);default:'';comment:'经度'" json:"longitude"`
	Speed       string    `gorm:"column:speed;type:varchar(255);default:'';comment:'速度'" json:"speed"`
	RotateAngle string    `gorm:"column:rotate_angle;type:varchar(255);default:'';comment:'角度'" json:"rotate_angle"`
	Battery     int       `gorm:"column:battery;type:int(11);comment:'电池电量\n'" json:"battery"`
	Step        int       `gorm:"column:step;type:int(11);comment:'步数\n'" json:"step"`
	AddTime     time.Time `gorm:"column:add_time;type:datetime;default:null on update current_timestamp;comment:'添加时间\n'" json:"add_time"`
	EditTime    time.Time `gorm:"column:edit_time;type:datetime;default:null on update current_timestamp;comment:'编辑时间\n'" json:"edit_time"`
}

type FUserReportsOut struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserId      string    `gorm:"column:user_id;type:varchar(255);default:'';comment:'用户ID'" json:"user_id"`
	Latitude    string    `gorm:"column:latitude;type:varchar(255);default:'';comment:'纬度'" json:"latitude"`
	Longitude   string    `gorm:"column:longitude;type:varchar(255);default:'';comment:'经度'" json:"longitude"`
	Speed       string    `gorm:"column:speed;type:varchar(255);default:'';comment:'速度'" json:"speed"`
	RotateAngle string    `gorm:"column:rotate_angle;type:varchar(255);default:'';comment:'角度'" json:"rotate_angle"`
	Battery     int       `gorm:"column:battery;type:int(11);comment:'电池电量\n'" json:"battery"`
	Step        int       `gorm:"column:step;type:int(11);comment:'步数\n'" json:"step"`
	AddTime     time.Time `gorm:"column:add_time;type:datetime;default:null on update current_timestamp;comment:'添加时间\n'" json:"add_time"`
	EditTime    time.Time `gorm:"column:edit_time;type:datetime;default:null on update current_timestamp;comment:'编辑时间\n'" json:"edit_time"`
	AddTimeStr  string    `json:"add_time_str"`
	EditTimeStr string    `json:"edit_time_str"`
}

// 上报用户信息
func CreateUserLocation(reports *FUserReports) error {
	reports.AddTime = time.Now()
	reports.EditTime = time.Now()
	fmt.Println(reports)
	return db.DB.MysqlDB.DefaultGormDB().Table("f_user_reports").Create(&reports).Error
}

// 获取用户列表
func GetUserLocationList(usersID /*16,17,18*/ string) ([]FUserReportsOut, error) {
	var userList []FUserReportsOut

	// 需要查询每个ID的最新的一条数据
	str := fmt.Sprintf("select * from f_user_reports where user_id in (%s) and id in (select max(id) from f_user_reports group by user_id)", usersID)
	err := db.DB.MysqlDB.DefaultGormDB().Raw(str).Scan(&userList).Error
	if err != nil {
		return nil, err
	}
	// 循环处理一下时间格式
	for i := 0; i < len(userList); i++ {
		userList[i].AddTimeStr = userList[i].AddTime.Format("2006-01-02 15:04:05")
		userList[i].EditTimeStr = userList[i].EditTime.Format("2006-01-02 15:04:05")
	}
	return userList, err
}

// 转化一下time 结构
