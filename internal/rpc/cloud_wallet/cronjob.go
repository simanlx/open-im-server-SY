package cloud_wallet

import (
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var NotifyChannel = make(chan string, 100)

func StarCorn() {
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("通知失败，err: ", err)
				}
			}()

			for {
				no := <-NotifyChannel
				// 查询订单信息
				err, payOrder := imdb.GetThirdPayOrderNo(no)
				if err != nil {
					log.Error("crontab", "查询订单失败，err: ", err)
					continue
				}
				if payOrder.Id == 0 {
					continue
				}
				content := map[string]interface{}{
					"OrderID":    payOrder.OrderNo,
					"MerOrderID": payOrder.MerOrderNo,
					"Status":     payOrder.Status,
					"CreateTime": payOrder.AddTime.Format("2006-01-02 15:04:05"),
					"PayTime":    payOrder.PayTime.Format("2006-01-02 15:04:05"),
				}

				// 转换成为json
				body, err := json.Marshal(content)
				if err != nil {
					continue
				}
				// 发送请求
				err = HttpPost(payOrder.NotifyUrl, body)
				if err != nil {
					// 修改订单状态
					err = imdb.UpdateThirdPayOrderCallback(0, int(payOrder.NotifyCount)+1, payOrder.OrderNo)
					if err != nil {
						log.Error("crontab", "修改订单状态失败，err: ", err)
					}
					continue
				}
				// 修改订单状态
				err = imdb.UpdateThirdPayOrderCallback(1, int(payOrder.NotifyCount)+1, payOrder.OrderNo)
				if err != nil {
					log.Error("crontab", "修改订单状态失败，err: ", err)
				}
			}

		}()
	}
}

func HttpPost(Url string, content []byte) error {
	body := bytes.NewBuffer(content)
	resp, err := http.Post(Url, "application/json;charset=utf-8", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New("请求失败")
	}
	return nil
}
