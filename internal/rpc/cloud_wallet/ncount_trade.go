package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"errors"
	"fmt"
	"github.com/spf13/cast"
)

const (
	BusinessTypeBankcardRecharge   = 1 //银行卡充值 (增加余额)
	BusinessTypeBankcardWithdrawal = 2 //银行卡提现 (减少余额)
	BusinessTypeBankcardSendPacket = 3 //银行卡支付发送红包 (余额不变)
	BusinessTypeBalanceSendPacket  = 4 //余额支付发送红包 (减少余额)
	BusinessTypeReceivePacket      = 5 //领取红包 (增加余额)
	BusinessTypePacketExpire       = 6 //红包过期 (退还,增加余额)
)

func BusinessTypeAttr(businessType, amount, balAmount int32) (int32, int32, int32, string, error) {
	switch businessType {
	case BusinessTypeBankcardRecharge:
		return 1, 0, balAmount + amount, "银行卡充值", nil
	case BusinessTypeBankcardWithdrawal:
		return 2, 0, balAmount - amount, "提现到银行卡", nil
	case BusinessTypeBankcardSendPacket:
		return 2, 0, balAmount, "银行卡支付发送红包", nil
	case BusinessTypeBalanceSendPacket:
		return 2, 1, balAmount - amount, "余额支付发送红包", nil
	case BusinessTypeReceivePacket:
		return 1, 1, balAmount + amount, "领取红包", nil
	case BusinessTypePacketExpire:
		return 1, 1, balAmount + amount, "红包超时退回", nil
	default:
		return 0, 0, 0, "", errors.New("业务类型错误")
	}
}

// 增加账户变更日志
func AddNcountTradeLog(businessType, amount int32, userId, mainAccountId, thirdOrderNo, packetID string) (err error) {
	//获取用户余额
	accountResp, err := ncount.NewCounter().CheckUserAccountInfo(&ncount.CheckUserAccountReq{
		OrderID: ncount.GetMerOrderID(),
		UserID:  mainAccountId,
	})
	fmt.Println("accountResp Println", accountResp, err)
	if err != nil {
		return errors.New(fmt.Sprintf("查询账户信息失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return errors.New(fmt.Sprintf("查询账户信息失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	//余额变更
	balAmount := cast.ToInt32(cast.ToFloat32(accountResp.BalAmount) * 100) //用户余额
	changeType, ncountStatus, afterAmount, describe, err := BusinessTypeAttr(businessType, amount, balAmount)
	if err != nil {
		return err
	}

	//数据入库
	return imdb.FNcountTradeCreateData(&db.FNcountTrade{
		UserID:       userId,
		Type:         changeType,
		BusinessType: businessType,
		Describe:     describe,
		Amount:       amount,
		BeferAmount:  balAmount,
		AfterAmount:  afterAmount,
		ThirdOrderNo: thirdOrderNo,
		NcountStatus: ncountStatus,
		PacketID:     packetID,
	})
}
