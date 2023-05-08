package agent_model

import "Open_IM/pkg/common/db"

// 获取推广员自定义商城咖豆配置
func GetAgentDiyShopBeanConfig(userId string) (data []*db.TAgentBeanShopConfig, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_bean_shop_config").Where("user_id = ?", userId).Order("id asc").Find(&data).Error
	return
}

// 获取推广员自定义商城上架咖豆配置
func GetAgentDiyShopBeanOnlineConfig(userId string) (data []*db.TAgentBeanShopConfig, err error) {
	err = db.DB.AgentMysqlDB.DefaultGormDB().Table("t_agent_bean_shop_config").Where("user_id = ? and status = ?", userId, 1).Order("id asc").Find(&data).Error
	return
}
