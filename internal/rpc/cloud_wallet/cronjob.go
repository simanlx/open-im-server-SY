package cloud_wallet

import (
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

var NotifyChannel = make(chan string, 100)

func StarCorn() {
	CornSelect()
	for i := 0; i < 50; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("Panic 通知失败，err: ", err)
				}
			}()
			for {
				MerOrderID := <-NotifyChannel // 我们平台生成的订单号
				Operation := "竞技回调merOrderID ：" + MerOrderID
				// 查询订单信息
				log.Debug(Operation, "订单号", MerOrderID)
				err, payOrder := imdb.GetThirdPayJdnMerOrderID(MerOrderID)
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						log.Error(Operation, "没有订单ID ：", MerOrderID)
					} else {
						log.Error(Operation, "查询订单失败 ：", err)
					}
					continue
				}

				// 如果请求超过5次就不再请求
				if payOrder.NotifyCount >= 5 {
					continue
				}

				// 如果超过两个小时后就不再请求
				if payOrder.AddTime.Add(time.Hour*2).Unix() < time.Now().Unix() {
					continue
				}

				// 检查上次通知时间到现在的时间间隔
				// 第一次是及时回调
				// 第二次是间隔5秒 ： 由线程休眠实现
				// 第三次是间隔30秒 ：数据库触发
				// 第四次后就是间隔5分钟
				// 第五次后就是间隔30分钟
				content := map[string]interface{}{
					"OrderID":    payOrder.OrderNo,
					"MerOrderID": payOrder.MerOrderNo,
					"Status":     payOrder.Status, // 100 是未支付，200是支付成功，300是支付失败
					"CreateTime": payOrder.AddTime.Format("2006-01-02 15:04:05"),
					"PayTime":    payOrder.PayTime.Format("2006-01-02 15:04:05"),
					"Amount":     payOrder.Amount, // 支付金额，以分为单位，整数
				}
				// 转换成为json
				body, err := json.Marshal(content)
				if err != nil {
					log.Error(Operation, "json转换失败，err: ", err)
					continue
				}
				// 发送请求
				err = HttpPost(payOrder.NotifyUrl, body)
				if err != nil {
					log.Error(Operation, "发送请求失败，err: ", err)
					// 修改订单状态
					err = imdb.UpdateThirdPayOrderCallback(0, int(payOrder.NotifyCount)+1, MerOrderID)
					if err != nil {
						log.Error(Operation, "修改订单状态失败，err: ", err)
					}
					continue
				}
				// 修改订单状态
				err = imdb.UpdateThirdPayOrderCallback(1, int(payOrder.NotifyCount)+1, MerOrderID)
				if err != nil {
					log.Error(Operation, "修改订单状态失败，err: ", err)
				}
			}

		}()
	}
}

// 每分钟查询过去2个小时的所有订单
func CornSelect() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("Panic ,定时任务：查询历史第三方支付订单失败，err: ", err)
			}
		}()
		for {
			// 每分钟查询过去2个小时所有订单
			start_time := time.Now().Add(-time.Hour * 2).Format("2006-01-02 15:04:05")
			end_time := time.Now().Format("2006-01-02 15:04:05")
			result, err := imdb.GetThirdPayOrderListByTime(start_time, end_time)
			log.Info("定时任务：查询历史第三方支付订单", fmt.Sprintf(""), time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				log.Error("查询订单失败，err: ", err)
				time.Sleep(time.Minute)
				continue
			}
			// 将订单数量写入channel通道
			for _, v := range result {

				// 1.如果是notify_count  =1 ,间隔时间为30秒
				if v.LastNotifyTime.Add(time.Second*30).Unix() > time.Now().Unix() && v.NotifyCount == 1 {
					continue
				}
				// 2.如果是notify_count  =2 ,间隔时间为5分钟
				if v.LastNotifyTime.Add(time.Minute*5).Unix() > time.Now().Unix() && v.NotifyCount == 2 {
					continue
				}
				// 3.如果是notify_count  =3 ,间隔时间为30分钟
				if v.LastNotifyTime.Add(time.Minute*30).Unix() > time.Now().Unix() && v.NotifyCount == 3 {
					continue
				}
				// 4.如果是notify_count  =4 ,间隔时间为30分钟
				if v.LastNotifyTime.Add(time.Minute*30).Unix() > time.Now().Unix() && v.NotifyCount == 4 {
					continue
				}
				// 5.如果是notify_count  =5 ,间隔时间为30分钟
				if v.LastNotifyTime.Add(time.Minute*30).Unix() > time.Now().Unix() && v.NotifyCount == 5 {
					continue
				}
				notifyThirdPay(v.NcountOrderNo)
			}
			time.Sleep(time.Minute)
		}
	}()
}

func notifyThirdPay(MerOrderID string) {
	NotifyChannel <- MerOrderID
}

func HttpPost(Url string, content []byte) error {
	body := bytes.NewBuffer(content)
	resp, err := http.Post(Url, "application/json;charset=utf-8", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Info("竞技回调-打印响应-发送请求", string(content), Url)
	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		return errors.New("返回响应HttpCode不为200," + strconv.Itoa(resp.StatusCode))
	}
	respContent, err := json.Marshal(resp.Body)
	log.Debug("竞技回调-打印响应", respContent, Url)
	return nil
}
