package agent

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// chess_user_id = 100006
const (
	ChessApiUrl = "http://serverlocal.xingdong.sh.cn:21512" //新互娱api url
)

type Record struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type AgentChessUser struct {
	AgentNumber     int32            `json:"agent_number"`
	AgentMemberList []*ChessUserInfo `json:"agent_member_list"`
}

type ChessUserInfo struct {
	BeanNumber int32  `json:"bean_number"`
	Nickname   string `json:"nickname"`
}

func httpPost(url string, form url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 获取互娱用户信息(咖豆、用户昵称、分页、排序)
func GetAgentChessUserList(agentNumber int32) map[int64]map[string]interface{} {

	return map[int64]map[string]interface{}{
		10018: {"bean_number": 0, "nickname": "昵称"},
	}
}

// 获取新互娱用户资料
func GetChessUserInfo(chessUserId int64) *ChessUserInfo {
	return &ChessUserInfo{
		BeanNumber: 100,
		Nickname:   "xxx",
	}
}

// 赠送新互娱用户咖豆
func GiveChessUserBean(chessUserId int64) int64 {
	/**
	加锁lock、事务
	1、校验咖豆余额(计算冻结额度)
	2、扣除推广员咖豆、账户表更日志
	3、api 互娱加豆、异常重试一次
	*/

	return 10
}
