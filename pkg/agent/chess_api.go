package agent

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/http"
	"Open_IM/pkg/common/log"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math"
)

type RechargeUserGoldReq struct {
	OrderId     string `json:"order_id"`
	Uid         int64  `json:"uid"`
	AgentNumber int32  `json:"agent_number"`
	Num         int64  `json:"num"`
}

type ChessApiResp struct {
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
}

type AgentChessUserListResp struct {
	Code int32            `json:"code"`
	Msg  string           `json:"msg"`
	Data []*ChessUserInfo `json:"data"`
}

type ChessUserInfo struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
	Gold     int64  `json:"gold"`
}

// 调用chess api 获取推广员下属成员列表
func ChessApiAgentChessMemberList(agentNumber int32) ([]*ChessUserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/v1/getUserListInfo?agent_number=%d", config.Config.Agent.ChessApiDomain, agentNumber))
	if err != nil {
		log.Error("", "请求chess api getUserListInfo", err.Error())
		return nil, errors.Wrap(err, "请求chess api getUserListInfo失败")
	}

	chessApiResp := &AgentChessUserListResp{Data: []*ChessUserInfo{}}
	_ = json.Unmarshal(resp, &chessApiResp)
	if chessApiResp.Code != 200 {
		errMsg := fmt.Sprintf("调用chess api 接口失败, err:%s", chessApiResp.Msg)
		log.Error("", "请求chess api getUserListInfo", errMsg)
		return nil, errors.New(errMsg)
	}
	return chessApiResp.Data, nil
}

// 获取推广员下属成员列表
func GetAgentChessMemberList(agentNumber, orderBy int32) ([]*ChessUserInfo, error) {
	return ChessApiAgentChessMemberList(agentNumber)
}

//func ArraySlice(s []interface{}, offset, length uint) []interface{} {
//	if offset > uint(len(s)) {
//		panic("offset: the offset is less than the length of s")
//	}
//	end := offset + length
//	if end < uint(len(s)) {
//		return s[offset:end]	}
//	return s[offset:]
//}

func SlicePage(sliceLen, page, size int) (sliceStart, sliceEnd int) {
	if size > sliceLen {
		return 0, sliceLen
	}

	// 总页数计算
	pageCount := int(math.Ceil(float64(sliceLen) / float64(size)))
	if page > pageCount {
		return 0, 0
	}

	sliceStart = (page - 1) * size
	sliceEnd = sliceStart + size
	if sliceEnd > sliceLen {
		sliceEnd = sliceLen
	}
	return sliceStart, sliceEnd
}

// 调用chess api 给用户加咖豆
func ChessApiGiveUserBean(orderNo string, agentNumber int32, chessUserId, beanNumber int64) error {
	data := RechargeUserGoldReq{
		OrderId:     orderNo,
		Uid:         chessUserId,
		AgentNumber: agentNumber,
		Num:         beanNumber,
	}
	resp, err := http.Post(config.Config.Agent.ChessApiDomain+"/v1/rechargeUserGold", data, 2)
	if err != nil {
		log.Error("", "请求chess api rechargeUserGold失败", err.Error())
		return errors.Wrap(err, "请求chess api rechargeUserGold失败")
	}

	chessApiResp := &ChessApiResp{}
	_ = json.Unmarshal(resp, &chessApiResp)
	if chessApiResp.Code != 200 {
		errMsg := fmt.Sprintf("调用chess api 给用户加咖豆失败, 订单号(%s),推广员编号(%d),互娱用户id(%d),咖豆数(%d),err:%s", orderNo, agentNumber, chessUserId, beanNumber, chessApiResp.Msg)
		log.Error("", "请求chess api rechargeUserGold失败", errMsg)
		return errors.New(errMsg)
	}
	return nil
}
