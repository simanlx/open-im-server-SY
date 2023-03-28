package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/db"
	imdb "Open_IM/pkg/common/db/mysql_model/cloud_wallet"
	"Open_IM/pkg/common/log"
	promePkg "Open_IM/pkg/common/prometheus"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	"Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/utils"
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"Open_IM/pkg/common/config"

	"google.golang.org/grpc"
)

type CloudWalletServer struct {
	cloud_wallet.UnimplementedCloudWalletServiceServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func (rpc *CloudWalletServer) mustEmbedUnimplementedCloudWalletServer() {
	return
}

func NewRpcCloudWalletServer(port int) *CloudWalletServer {
	log.NewPrivateLog(constant.LogFileName)
	return &CloudWalletServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OpenImCloudWalletName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (rpc *CloudWalletServer) Run() {
	operationID := utils.OperationIDGenerator()
	log.NewInfo(operationID, "rpc auth start...")

	listenIP := ""
	if config.Config.ListenIP == "" {
		listenIP = "0.0.0.0"
	} else {
		listenIP = config.Config.ListenIP
	}
	address := listenIP + ":" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error() + rpc.rpcRegisterName)
	}
	log.NewInfo(operationID, "listen network success, ", address, listener)
	var grpcOpts []grpc.ServerOption
	if config.Config.Prometheus.Enable {
		promePkg.NewGrpcRequestCounter()
		promePkg.NewGrpcRequestFailedCounter()
		promePkg.NewGrpcRequestSuccessCounter()
		grpcOpts = append(grpcOpts, []grpc.ServerOption{
			// grpc.UnaryInterceptor(promePkg.UnaryServerInterceptorProme),
			grpc.StreamInterceptor(grpcPrometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
		}...)
	}
	srv := grpc.NewServer(grpcOpts...)
	defer srv.GracefulStop()

	//service registers with etcd
	cloud_wallet.RegisterCloudWalletServiceServer(srv, rpc)
	rpcRegisterIP := config.Config.RpcRegisterIP
	if config.Config.RpcRegisterIP == "" {
		rpcRegisterIP, err = utils.GetLocalIP()
		if err != nil {
			log.Error("", "GetLocalIP failed ", err.Error())
		}
	}
	log.NewInfo("", "rpcRegisterIP", rpcRegisterIP)

	err = getcdv3.RegisterEtcd(rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), rpcRegisterIP, rpc.rpcPort, rpc.rpcRegisterName, 10)
	if err != nil {
		log.NewError(operationID, "RegisterEtcd failed ", err.Error(),
			rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), rpcRegisterIP, rpc.rpcPort, rpc.rpcRegisterName)
		panic(utils.Wrap(err, "register auth module  rpc to etcd err"))

	}
	log.NewInfo(operationID, "RegisterAuthServer ok ", rpc.etcdSchema, strings.Join(rpc.etcdAddr, ","), rpcRegisterIP, rpc.rpcPort, rpc.rpcRegisterName)
	err = srv.Serve(listener)
	if err != nil {
		log.NewError(operationID, "Serve failed ", err.Error())
		return
	}
	log.NewInfo(operationID, "rpc auth ok")
}

// 获取用户余额
func (rpc *CloudWalletServer) UserAccountBalance(ctx context.Context, req *cloud_wallet.UserAccountBalanceReq) (*cloud_wallet.UserAccountBalanceResp, error) {
	return &cloud_wallet.UserAccountBalanceResp{
		MainBalance: 100,
	}, nil
}

// 获取云账户信息
func (rpc *CloudWalletServer) UserNcountAccount(ctx context.Context, req *cloud_wallet.UserNcountAccountReq) (*cloud_wallet.UserNcountAccountResp, error) {
	return &cloud_wallet.UserNcountAccountResp{
		Step:          2,
		RealAuth:      1,
		IdCard:        "456453",
		RealName:      "名字",
		AccountStatus: 0,
	}, nil
}

// 身份证实名认证
func (rpc *CloudWalletServer) IdCardRealNameAuth(_ context.Context, req *cloud_wallet.IdCardRealNameAuthReq) (*cloud_wallet.IdCardRealNameAuthResp, error) {
	info := &db.FNcountAccount{
		UserId:      req.UserId,
		Mobile:      req.Mobile,
		RealName:    req.RealName,
		IdCard:      req.IdCard,
		OpenStep:    1,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	//实名数据入库
	err := imdb.CreateNcountAccount(info)
	if err != nil {
		return nil, errors.New("实名认证数据入库失败")
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	//开通新生支付主账户
	go func(wg *sync.WaitGroup, info *db.FNcountAccount) {
		defer wg.Done()

		id := fmt.Sprintf("main_%d", info.UserId)
		accountResp, err := ncount.NewCounter().NewAccount(&ncount.NewAccountReq{
			OrderID: id,
			MsgCipherText: &ncount.NewAccountMsgCipherText{
				MerUserId: id,
				Mobile:    info.Mobile,
				UserName:  info.RealName,
				CertNo:    info.IdCard,
			},
		})

		fmt.Println("accountResp", accountResp)

		if err != nil {
			return
		}
		info.MainAccountId = accountResp.UserId
	}(wg, info)

	//开通新生支付钱包账户
	go func(wg *sync.WaitGroup, info *db.FNcountAccount) {
		defer wg.Done()

		id := fmt.Sprintf("packet_%d", info.UserId)
		accountResp, err := ncount.NewCounter().NewAccount(&ncount.NewAccountReq{
			OrderID: id,
			MsgCipherText: &ncount.NewAccountMsgCipherText{
				MerUserId: id,
				Mobile:    info.Mobile,
				UserName:  info.RealName,
				CertNo:    info.IdCard,
			},
		})

		fmt.Println("accountResp", accountResp)

		if err != nil {
			return
		}
		info.PacketAccountId = accountResp.UserId
	}(wg, info)

	wg.Wait()

	//更新新生支付账户id
	err = imdb.UpdateNcountAccountInfo(info)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("更新数据失败%s", err.Error()))
	}

	return &cloud_wallet.IdCardRealNameAuthResp{
		Step: 1,
		CommonResp: &cloud_wallet.CommonResp{
			ErrCode: 0,
			ErrMsg:  "",
		},
	}, nil
}

// 设置用户支付密码
func (rpc *CloudWalletServer) SetPaymentSecret(ctx context.Context, req *cloud_wallet.SetPaymentSecretReq) (*cloud_wallet.SetPaymentSecretResp, error) {
	return &cloud_wallet.SetPaymentSecretResp{
		CommonResp: &cloud_wallet.CommonResp{
			ErrCode: 1,
			ErrMsg:  "not",
		},
	}, nil
}

// 云钱包收支明细
func (rpc *CloudWalletServer) CloudWalletRecordList(ctx context.Context, req *cloud_wallet.CloudWalletRecordListReq) (*cloud_wallet.CloudWalletRecordListResp, error) {
	return &cloud_wallet.CloudWalletRecordListResp{
		Total: 100,
	}, nil
}

// 获取用户银行卡列表
func (rpc *CloudWalletServer) GetUserBankcardList(_ context.Context, req *cloud_wallet.GetUserBankcardListReq) (*cloud_wallet.GetUserBankcardListResp, error) {
	list, err := imdb.GetUserBankcardByUserId(req.UserId)
	resp := &cloud_wallet.GetUserBankcardListResp{}
	for _, v := range list {
		resp.BankCardList = append(resp.BankCardList, &cloud_wallet.BankCardList{
			BankCardID:     v.Id,
			CardOwner:      v.CardOwner,
			BankCardType:   v.BankCardType,
			BankCardNumber: v.BankCardNumber,
			CreatedTime:    v.CreatedTime.Format("2006-01-02 15:04:05"),
		})
	}

	return resp, err
}

// 绑定用户银行卡
func (rpc *CloudWalletServer) BindUserBankcard(ctx context.Context, req *cloud_wallet.BindUserBankcardReq) (*cloud_wallet.BindUserBankcardResp, error) {
	//新生支付接口预提交
	merOrderId, ncountOrderId, bindCardAgrNo := "434234424", "423424242", "42342424"

	info := &db.FNcountBankCard{
		UserId:         req.UserId,
		MerOrderId:     merOrderId,
		NcountOrderId:  ncountOrderId,
		BindCardAgrNo:  bindCardAgrNo,
		Mobile:         req.Mobile,
		CardOwner:      req.CardOwner,
		BankCardType:   req.BankCardType,
		BankCardNumber: req.GetBankCardNumber(),
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
	}

	//数据入库
	err := imdb.BindUserBankcard(info)

	return &cloud_wallet.BindUserBankcardResp{
		BankCardId: info.Id,
	}, err
}

// 绑定用户银行卡确认code
func (rpc *CloudWalletServer) BindUserBankcardConfirm(ctx context.Context, req *cloud_wallet.BindUserBankcardConfirmReq) (*cloud_wallet.BindUserBankcardConfirmResp, error) {
	//新生支付确定接口

	//更新数据
	err := imdb.BindUserBankcardConfirm(req.BankCardId, req.UserId, "")

	return &cloud_wallet.BindUserBankcardConfirmResp{BankCardId: req.BankCardId}, err
}

// 解绑用户银行卡
func (rpc *CloudWalletServer) UnBindingUserBankcard(ctx context.Context, req *cloud_wallet.UnBindingUserBankcardReq) (*cloud_wallet.UnBindingUserBankcardResp, error) {
	//新生支付解绑接口

	//更新数据库
	err := imdb.UnBindUserBankcard(req.BankCardId, req.UserId)

	return &cloud_wallet.UnBindingUserBankcardResp{}, err
}
