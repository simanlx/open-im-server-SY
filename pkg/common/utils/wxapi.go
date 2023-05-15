package utils

import (
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

// 推广员赠送咖豆
func AgentGiveMemberBean(agentNumber int32, chessUserId, beanNumber int64, errMsg string) {
	content := fmt.Sprintf("推广系统业务告警\n> 业务类型 : <font color=\"warning\">推广员赠送咖豆</font>\n> 推广员编号 : <font color=\"comment\">%d</font>\n> 互娱用户ID : <font color=\"comment\">%d</font>\n> 赠送咖豆数 : <font color=\"comment\">%d</font>\n> 错误原因 : <font color=\"comment\">%s</font> ", agentNumber, chessUserId, beanNumber, errMsg)

	qyApi := &QyApi{}
	qyApi.Msgtype = "markdown"
	qyApi.Markdown.Content = content

	key := "b862ac73-b394-424f-b089-25bdb460cf0d"
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, qyApi, 3)
	if err != nil {
		errMsg := fmt.Sprintf("推广员编号(%d) 、调用wxapi AgentGiveMemberBean失败：%s", agentNumber, err.Error())
		log.Error("", "请求wxapi AgentGiveMemberBean失败", errMsg)
		return
	}

	wxApiResp := &QyApiResp{}
	_ = json.Unmarshal(resp, &wxApiResp)
	if wxApiResp.Errcode != 0 {
		errMsg := fmt.Sprintf("推广员编号(%d) 、调用wxapi AgentGiveMemberBean失败：%s", agentNumber, wxApiResp.Errmsg)
		log.Error("", "请求wxapi AgentGiveMemberBean失败", errMsg)
	}
	return
}
