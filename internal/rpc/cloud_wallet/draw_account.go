package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"net/url"
)

// 账户提现接口

// Path: internal/rpc/cloud_wallet/draw_account.go

func (c *CloudWalletServer) UserWithdrawal(ctx context.Context, in *pb.DrawAccountReq) (*pb.DrawAccountResp, error) {

	return nil, nil
}

type userWithdrawal struct {
	notify string
	count  ncount.NCounter
}

func NewUserWithdrawal(count ncount.NCounter, notify string) (*userWithdrawal, error) {
	// 验证notify是不是合格的url
	_, err := url.ParseRequestURI(notify)
	if err != nil {
		return nil, err
	}

	return &userWithdrawal{
		count:  count,
		notify: notify,
	}, nil
}

// 用户提现
func (u *userWithdrawal) userWithdrawal(userid, bankcardID, amount, OperateId string) (*pb.CommonResp, error) {
	// 1.校验参数
	//if userid == "" || bankcardID == "" || amount == "" || OperateId == "" {
	//	return nil, errors.New("param is nil")
	//}
	var respCommon = &pb.CommonResp{
		ErrCode: 0,
		ErrMsg:  "操作成功",
	}
	//// todo 校验银行卡和用户的关系，但是在第三方已经处理完成了
	//// 查询用户的第三方账号ID ，红包账户| 云钱包账户
	//NccountUserID, BindCardAgrNo, err := cloud_wallet.GetNcountBankCardByBindCardAgrNo(bankcardID, userid)
	//if err != nil {
	//	log.Error(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, err)
	//	return nil, err
	//}
	//log.Info(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, nil)
	//// 调用提现接口
	//req := &ncount.QuickPayOrderReq{
	//	MerOrderId: ncount.GetMerOrderID(),
	//	QuickPayMsgCipher: ncount.QuickPayMsgCipher{
	//		TranAmount:    "1",
	//		NotifyUrl:     u.notify,
	//		ReceiveUserId: NccountUserID,
	//		BindCardAgrNo: BindCardAgrNo,
	//		SubMerchantId: "3", // 充值
	//	}}
	//resp, err := u.count.QuickPayOrder(req)
	//if err != nil {
	//	log.Error(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, err)
	//	return nil, err
	//}
	//if resp.ResultCode == "4444" {
	//	// 充值失败
	//	log.Error(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, resp)
	//	respCommon.ErrMsg = resp.ErrorMsg
	//	respCommon.ErrCode = pb.CloudWalletErrCode_ServerNcountError
	//}
	return respCommon, nil
}
