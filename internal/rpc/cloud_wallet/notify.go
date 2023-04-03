package cloud_wallet

import (
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	pb "Open_IM/pkg/proto/cloud_wallet"
	"context"
	"github.com/pkg/errors"
)

// 这里是处理充值和回调的地方

func (s *CloudWalletServer) ChargeNotify(ctx context.Context, req *pb.ChargeNotifyReq) (*pb.ChargeNotifyResp, error) {
	var (
		resp = &pb.ChargeNotifyResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "修改成功",
			},
		}
	)
	// 这里处理充值回调接口
	// 1.检查订单是否存在
	if req.MerOrderId == "" {
		return nil, errors.New("订单号不能为空")
	}

	//todo
	if req.ResultCode != "0000" {
		// 这里需要通知用户发送红包失败
	}
	if req.ResultCode == "0000" {
		// 需要发送code到所有群用户
	}

	f := &imdb.FNcountTrade{
		ThirdOrderNo: req.MerOrderId,
		NcountStatus: 1, // 表示修改成功
	}
	err := imdb.FNcountTradeUpdateStatusbyThirdOrderNo(f)
	if err != nil {
		log.Error("修改订单状态失败", err, req)
		return nil, err
	}
	// 2.修改订单状态
	return resp, nil
}

// 提现回调接口

func (s *CloudWalletServer) WithDrawNotify(ctx context.Context, req *pb.DrawNotifyReq) (*pb.DrawNotifyResp, error) {
	var (
		resp = &pb.DrawNotifyResp{
			CommonResp: &pb.CommonResp{
				ErrCode: 0,
				ErrMsg:  "修改成功",
			},
		}
	)
	// 这里处理充值回调接口
	// 1.检查订单是否存在
	if req.MerOrderId == "" {
		return nil, errors.New("订单号不能为空")
	}

	//todo
	if req.ResultCode != "0000" {
		// 这里需要通知用户发送红包失败
	}
	if req.ResultCode == "0000" {
		// 需要发送code到所有群用户
	}

	f := &imdb.FNcountTrade{
		ThirdOrderNo: req.MerOrderId,
		NcountStatus: 1, // 表示修改成功
	}
	err := imdb.FNcountTradeUpdateStatusbyThirdOrderNo(f)
	if err != nil {
		log.Error("修改订单状态失败", err, req)
		return nil, err
	}
	// 2.修改订单状态
	return resp, nil
}
