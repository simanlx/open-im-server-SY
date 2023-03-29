package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"github.com/pkg/errors"
)

// 抢红包接口
func (rpc *CloudWalletServer) ClickRedPacket(ctx context.Context, req *pb.ClickRedPacketReq) (*pb.ClickRedPacketResp, error) {
	// 参数校验 ：红包是否过期、 红包状态判断、红包的类型
	// 判断抢红包是否实名认证过，如果没实名认证过就不能抢红包
	// 如果是群聊红包 ： 请求进行管理
	// 如果是个人红包： 调用转账接口 ，
	// 生成红包记录日志
	// 记录用户收入日志
	// 修改红包的余额

	// 检查用户是否实名认证

}

type handlerClickRedPacket struct {
	*pb.ClickRedPacketReq
	OperateID string
	count     ncount.NCounter
}

func (req *handlerClickRedPacket) ClickRedPacket() (*pb.ClickRedPacketResp, error) {

	var (
		res    = &pb.ClickRedPacketResp{}
		common = &pb.CommonResp{}
	)
	// 1. 检测红包是否过期
	// 校验红包的状态
	redPacketInfo, err := imdb.GetRedPacketInfo(req.RedPacketID)
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		common.ErrMsg = "服务器错误"
		return res, errors.Wrap(err, "获取红包信息失败")
	}
	if redPacketInfo.Status == imdb.RedPacketStatusCreate {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsCreate
		common.ErrMsg = "红包还未创建"
		return res, nil
	}
	if redPacketInfo.Status == imdb.RedPacketStatusFinished {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsFinish
		common.ErrMsg = "红包已经抢完"
		return res, nil
	}
	if redPacketInfo.Status == imdb.RedPacketStatusExpired {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExpire
		common.ErrMsg = "红包过期"
		return res, nil
	}

	if redPacketInfo.IsExclusive == 1 && redPacketInfo.ExclusiveUserID != req.UserId {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsExclusive
		common.ErrMsg = "红包是专属红包"
		return res, nil
	}

	// 2. 检测红包的领取记录 ，如果已经完成领取就不能再领取
	fp, err := imdb.FPacketDetailGetByPacketID(req.RedPacketID, req.UserId)
	if err != nil {
		common.ErrCode = pb.CloudWalletErrCode_ServerError
		common.ErrMsg = "服务器错误"
		return res, errors.Wrap(err, "获取红包领取记录失败")
	}

	if fp.ID != 0 {
		common.ErrCode = pb.CloudWalletErrCode_PacketStatusIsReceived
		common.ErrMsg = "红包已经领取"
		return res, nil
	}
	// 3. 检查用户是否实名认证 todo

	// 4. 判断红包的类型
	if redPacketInfo.PacketType == 1 {
		// 群聊红包
		amount, err := req.getRedPacketByGroup()
		if err != nil {
			common.ErrCode = pb.CloudWalletErrCode_ServerError
			common.ErrMsg = "服务器错误"
			return res, errors.Wrap(err, "获取红包失败")
		}
		return res, nil
	} else {
		// 个人红包 ,发起转账

	}
}

// 检查用户实名认证状态
func (req *handlerClickRedPacket) checkUserAuthStatus() (*pb.CommonResp, error) {
	res := &pb.CommonResp{}
	// 检查用户是否实名认证
	// 检查用户是否实名认证
	authStatus, err := req.count.GetUserAuthStatus(req.UserId)
	if err != nil {
		res.ErrCode = pb.CloudWalletErrCode_ServerError
		res.ErrMsg = "服务器错误"
		return res, errors.Wrap(err, "获取用户实名认证状态失败")
	}
	if authStatus != 1 {
		res.ErrCode = pb.CloudWalletErrCode_UserNotAuth
		res.ErrMsg = "用户未实名认证"
		return res, nil
	}
	return res, nil
}

// 如果红包是群聊红包， 直接redis的set集合进行获取红包
func (req *handlerClickRedPacket) getRedPacketByGroup() (int, error) {
	res := &pb.CommonResp{}
	amount, err := db.DB.GetRedPacket(req.RedPacketID)
	if err != nil {
		res.ErrCode = pb.CloudWalletErrCode_ServerError
		return 0, err
	}
	if amount > 0 {

	}
}

// // 调用转账接口  从红包账户 转账到 抢红包的用户的账户
func (req *handlerClickRedPacket) getRedPacketByUser(UserID int) (*pb.CommonResp, error) {
	//  获取到发红包的用户ID

	// 调用转正接口
}
