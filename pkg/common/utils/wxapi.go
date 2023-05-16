package utils

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/http"
	"Open_IM/pkg/common/log"
	"encoding/json"
	"fmt"
)

type QyApi struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type QyApiResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// 推广员提现告警
func WithdrawNotifyWarn(orderNo string, agentNumber, balance, amount int32) {
	content := fmt.Sprintf("推广系统业务告警\n> 业务类型 : <font color=\"warning\">推广员提现金额异常</font>\n> 推广员编号 : <font color=\"comment\">%d</font>\n> 订单号 : <font color=\"comment\">%s</font>\n> 提现金额 : <font color=\"comment\">%d</font>\n> 到账金额 : <font color=\"comment\">%d</font> ", agentNumber, orderNo, balance, amount)

	//内容格式 - markdown
	qyApi := &QyApi{}
	qyApi.Msgtype = "markdown"
	qyApi.Markdown.Content = content

	WxApi(qyApi, content)
}

// 推广员赠送咖豆告警
func AgentGiveMemberBeanWarn(agentNumber int32, chessUserId, beanNumber int64, errMsg string) {
	content := fmt.Sprintf("推广系统业务告警\n> 业务类型 : <font color=\"warning\">推广员赠送咖豆</font>\n> 推广员编号 : <font color=\"comment\">%d</font>\n> 互娱用户ID : <font color=\"comment\">%d</font>\n> 赠送咖豆数 : <font color=\"comment\">%d</font>\n> 错误原因 : <font color=\"comment\">%s</font> ", agentNumber, chessUserId, beanNumber, errMsg)

	//内容格式 - markdown
	qyApi := &QyApi{}
	qyApi.Msgtype = "markdown"
	qyApi.Markdown.Content = content

	WxApi(qyApi, content)
}

// 请求企业微信api
func WxApi(data interface{}, content string) {
	key := config.Config.Agent.AgentRechargeNotifyUrl
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, data, 3)
	if err != nil {
		log.Error("", fmt.Sprintf("通知内容(%s) 、api请求失败：%s", content, err.Error()))
		return
	}

	wxApiResp := &QyApiResp{}
	_ = json.Unmarshal(resp, &wxApiResp)
	if wxApiResp.Errcode != 0 {
		log.Error("", fmt.Sprintf("通知内容(%s) 、失败：%s", content, wxApiResp.Errmsg))
	}
	return
}
