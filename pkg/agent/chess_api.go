package agent

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ChessApiUrl = "" //新互娱api url
)

func httpPost(url string, form url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 获取互娱用户信息(咖豆、用户昵称、分页、排序)
func GetChessUserInfo(chessUserIds []int64) map[int64]map[string]interface{} {

	return map[int64]map[string]interface{}{
		10018: {"bean_number": 0, "nickname": "昵称"},
	}
}

// 赠送新互娱用户咖豆
func GiveChessUserBean(chessUserId int64) int64 {

	return 10
}
