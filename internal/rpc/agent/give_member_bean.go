package agent

import (
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/agent_model"
	rocksCache "Open_IM/pkg/common/db/rocks_cache"
	"Open_IM/pkg/common/http"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/utils"
	"Open_IM/pkg/proto/agent"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

const (
	ChessApiUrl = "http://serverlocal.xingdong.sh.cn:215120" //新互娱api url
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

// 赠送下属成员咖豆
func (rpc *AgentServer) AgentGiveMemberBean(ctx context.Context, req *agent.AgentGiveMemberBeanReq) (*agent.AgentGiveMemberBeanResp, error) {
	resp := &agent.AgentGiveMemberBeanResp{CommonResp: &agent.CommonResp{Code: 0, Msg: ""}}
	log.Info(req.OperationId, fmt.Sprintf("start 推广员(%s),赠送下属成员(%d)咖豆,赠送咖豆数(%d)", req.UserId, req.ChessUserId, req.BeanNumber))

	// 加锁
	lockKey := fmt.Sprintf("AgentGiveMemberBean:%d", req.ChessUserId)
	if err := utils.Lock(ctx, lockKey); err != nil {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "操作加锁失败"
		return resp, nil
	}
	defer utils.UnLock(ctx, lockKey)

	//获取推广员信息
	info, err := imdb.GetAgentByUserId(req.UserId)
	if err != nil {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "推广员信息有误"
		return resp, nil
	}

	//是否为推广员下属成员
	agentMember, err := imdb.AgentNumberByChessUserId(req.ChessUserId)
	if err != nil || agentMember.UserId != req.UserId {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "不是推广员下属成员"
		return resp, nil
	}

	//冻结的咖豆额度
	freezeBeanBalance := rocksCache.GetAgentFreezeBeanBalance(ctx, info.AgentNumber)

	//校验推广员咖豆余额
	if info.BeanBalance < (req.BeanNumber + freezeBeanBalance) {
		log.Error(req.OperationId, fmt.Sprintf("推广员(%d),赠送下属成员(%d)咖豆,咖豆余额不足,咖豆余额(%d),赠送咖豆(%d),冻结咖豆(%d)", info.AgentNumber, req.ChessUserId, info.BeanBalance, req.BeanNumber, freezeBeanBalance))
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = "咖豆余额不足"
		return resp, nil
	}

	// 赠送新互娱用户咖豆
	err = GiveChessUserBean(req.UserId, info.AgentNumber, info.Id, req.BeanNumber, req.ChessUserId, agentMember.ChessNickname)
	if err != nil {
		resp.CommonResp.Code = 400
		resp.CommonResp.Msg = err.Error()
		log.Error(req.OperationId, fmt.Sprintf("推广员(%d),赠送下属成员(%d)咖豆,操作失败:%s", info.AgentNumber, req.ChessUserId, err.Error()))
		return resp, nil
	}

	return resp, nil
}

// 赠送互娱用户咖豆
func GiveChessUserBean(userId string, agentNumber int32, agentId, beanNumber, chessUserId int64, chessNickname string) error {
	//开启事务
	tx := db.DB.AgentMysqlDB.DefaultGormDB().Begin()

	//1、扣除推广员咖豆
	err := tx.Table("t_agent_account").Where("id = ? and bean_balance >= ?", agentId, beanNumber).UpdateColumn("bean_balance", gorm.Expr(" bean_balance - ? ", beanNumber)).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "扣除推广员咖豆失败")
	}

	//2、增加咖豆变更日志
	orderNo := GetOrderNo()
	record := &db.TAgentBeanAccountRecord{
		OrderNo:      orderNo,
		UserId:       userId,
		Type:         2,
		BusinessType: imdb.BeanAccountBusinessTypeGive,
		Describe:     fmt.Sprintf("赠送给%sID:%d %d咖豆", chessNickname, chessUserId, beanNumber),
		Amount:       0,
		Number:       beanNumber,
		GiveNumber:   0,
		Day:          time.Now().Format("2006-01-02"),
		CreatedTime:  time.Now(),
		UpdatedTime:  time.Now(),
		DB:           tx,
	}
	err = tx.Table("t_agent_bean_account_record").Create(&record).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "增加咖豆变更日志失败")
	}

	//3、调用chess api 给用户加咖豆
	if !GiveUserBeanRetry(orderNo, agentNumber, chessUserId, beanNumber, []int{0, 3}) {
		tx.Rollback()
		return errors.New("调用chess api 给用户加咖豆失败")
	}

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		log.NewError("", "GiveChessUserBean commit error:", err, "tx.Rollback().Error :", tx.Rollback().Error)
		return errors.Wrap(err, "事务提交失败")
	}
	return nil
}

// 调用chess api 给用户加咖豆
func ChessApiGiveUserBean(orderNo string, agentNumber int32, chessUserId, beanNumber int64) error {
	data := RechargeUserGoldReq{
		OrderId:     orderNo,
		Uid:         chessUserId,
		AgentNumber: agentNumber,
		Num:         beanNumber,
	}
	resp, err := http.Post(ChessApiUrl+"/v1/rechargeUserGold", data, 2)
	if err != nil {
		log.Error("", "请求chess api rechargeUserGold失败", err.Error())
		return errors.Wrap(err, "请求chess api rechargeUserGold失败")
	}

	chessApiResp := &ChessApiResp{}
	_ = json.Unmarshal(resp, &chessApiResp)
	if chessApiResp.Code != 200 {
		errMsg := fmt.Sprintf("调用chess api 给用户加咖豆失败, 订单号(%s),推广员编号(%d),互娱用户id(%d),咖豆数(%d),err:%s", orderNo, agentNumber, chessUserId, beanNumber, chessApiResp.Msg)
		log.Error("", errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// 赠送用户咖豆(重试)
func GiveUserBeanRetry(orderNo string, agentNumber int32, chessUserId, beanNumber int64, intervals []int) bool {
	var retryCh = make(chan bool)
	index := 0

	for {
		go time.AfterFunc(time.Duration(intervals[index])*time.Second, func() {
			err := ChessApiGiveUserBean(orderNo, agentNumber, chessUserId, beanNumber)
			log.Info("", "ChessApiGiveUserBean err:", err, orderNo, agentNumber, chessUserId, beanNumber)
			if err == nil {
				retryCh <- true
			} else {
				retryCh <- false
			}
		})

		if <-retryCh {
			return true
		}

		if len(intervals)-1 == index {
			log.Info("", "GiveUserBeanRetry ---次数索引-index:", index)
			return false
		}

		index++
	}
}

func GetOrderNo() string {
	// 生成一串随机数
	// 时间戳 + 6位随机数
	tim := time.Now()
	times := tim.Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(999999)
	orderNo := times + strconv.Itoa(randNum)
	return orderNo
}
