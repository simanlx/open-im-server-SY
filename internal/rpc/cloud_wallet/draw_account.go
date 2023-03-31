package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net/url"
)

// 账户提现接口

// Path: internal/rpc/cloud_wallet/draw_account.go

func (c *CloudWalletServer) UserWithdrawal(ctx context.Context, in *pb.DrawAccountReq, opts ...grpc.CallOption) (*pb.DrawAccountResp, error) {

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
	if userid == "" || bankcardID == "" || amount == "" || OperateId == "" {
		return nil, errors.New("param is nil")
	}
	var respCommon = &pb.CommonResp{
		ErrCode: 0,
		ErrMsg:  "操作成功",
	}
	// todo 校验银行卡和用户的关系，但是在第三方已经处理完成了
	// 查询用户的第三方账号ID ，红包账户| 云钱包账户
	NccountUserID, BindCardAgrNo, err := u.getUserIDAndBindCardAgrNo(userid)
	if err != nil {
		log.Error(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, err)
		return nil, err
	}
	log.Info(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, nil)
	// 调用提现接口
	req := &ncount.QuickPayOrderReq{
		MerOrderId: ncount.GetMerOrderID(),
		QuickPayMsgCipher: ncount.QuickPayMsgCipher{
			TranAmount:    "1",
			NotifyUrl:     u.notify,
			ReceiveUserId: NccountUserID,
			BindCardAgrNo: BindCardAgrNo,
			SubMerchantId: "3", // 充值
		}}
	resp, err := u.count.QuickPayOrder(req)
	if err != nil {
		log.Error(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, err)
		return nil, err
	}
	if resp.ResultCode == "4444" {
		// 充值失败
		log.Error(OperateId, "userRecharge", "userid:"+userid+" bankcardID:"+bankcardID+" amount:"+amount+" OperateId:"+OperateId, resp)
		respCommon.ErrMsg = resp.ErrorMsg
		respCommon.ErrCode = pb.CloudWalletErrCode_ServerNcountError
	}
	return respCommon, nil
}

// 根据userid 查询第三方的用户ID 和 绑卡协议号
func (u *userWithdrawal) getUserIDAndBindCardAgrNo(userid string) (string /*NcountUserID*/, string /*BindCardAgrNo*/, error) {
	// todo 根据userid 查询第三方的用户ID 和 绑卡协议号
	result, err := imdb.GetNcountBankCardByUserIdAndType(userid, 1) // 查询用户的主钱包
	if err != nil {
		return "", "", errors.Wrap(err, "获取用户第三方信息")
	}
	if result != nil {
		return "", "", errors.New("用户未绑定银行卡")
	}

	return result.NcountOrderId, result.BindCardAgrNo, nil
}
