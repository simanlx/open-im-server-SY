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
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"strconv"
	"strings"
	"time"

	"Open_IM/pkg/common/config"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

const (
	UserMainAccountPrefix   = "main_"   //主账户前缀 完整账户id + 用户id
	UserPacketAccountPrefix = "packet_" //红包账户前缀 完整账户id + 用户id
)

type CloudWalletServer struct {
	cloud_wallet.UnimplementedCloudWalletServiceServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string

	// 依赖钱包服务
	count ncount.NCounter
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

// 获取云账户信息
func (rpc *CloudWalletServer) UserNcountAccount(ctx context.Context, req *cloud_wallet.UserNcountAccountReq) (*cloud_wallet.UserNcountAccountResp, error) {
	//获取用户账户信息
	accountInfo, err := imdb.GetNcountAccountByUserId(req.UserId)
	if err != nil || accountInfo.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("查询账户数据失败 %d,error:%s", req.UserId, err.Error()))
	}

	operationID := utils.OperationIDGenerator()
	log.Info(operationID, "accountInfo", accountInfo, err)
	fmt.Println("accountInfo Println", accountInfo, err)

	accountResp, err := ncount.NewCounter().CheckUserAccountInfo(&ncount.CheckUserAccountReq{
		OrderID: ncount.GetMerOrderID(),
		UserID:  accountInfo.MainAccountId,
	})

	log.Info(operationID, "accountResp", &accountResp, err)
	fmt.Println("accountResp Println", accountResp, err)
	if err != nil || accountResp.ResultCode != "0000" {
		return nil, errors.New(fmt.Sprintf("查询账户信息失败,code:"))
	}

	//绑定的银行卡列表
	bindCardsList := make([]*cloud_wallet.BindCardsList, 0)
	if len(accountResp.BindCards) > 0 {
		for _, v := range accountResp.BindCards {
			bindCardsList = append(bindCardsList, &cloud_wallet.BindCardsList{
				IndCardAgrNo: v.IndCardAgrNo,
				BankCode:     v.BankCode,
				CardNo:       v.CardNo,
			})
		}
	}

	return &cloud_wallet.UserNcountAccountResp{
		Step:             accountInfo.OpenStep,
		IdCard:           accountInfo.IdCard,
		RealName:         accountInfo.RealName,
		AccountStatus:    accountInfo.OpenStatus,
		BalAmount:        accountResp.BalAmount,
		AvailableBalance: accountResp.AvailableBalance,
		BindCardsList:    bindCardsList,
	}, nil
}

// 身份证实名认证
func (rpc *CloudWalletServer) IdCardRealNameAuth(_ context.Context, req *cloud_wallet.IdCardRealNameAuthReq) (*cloud_wallet.IdCardRealNameAuthResp, error) {
	//获取用户账户信息
	accountInfo, err := imdb.GetNcountAccountByUserId(req.UserId)
	if accountInfo != nil && accountInfo.Id > 0 {
		return nil, errors.New("已实名认证,请勿重复操作")
	}

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
	err = imdb.CreateNcountAccount(info)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("数据入库失败:%s", err.Error()))
	}

	errGroup := new(errgroup.Group)
	accountIds := []string{
		fmt.Sprintf("%s%d", UserMainAccountPrefix, info.UserId),
		fmt.Sprintf("%s%d", UserPacketAccountPrefix, info.UserId),
	}
	for _, account := range accountIds {
		id := account
		errGroup.Go(func() error {
			accountResp, err := ncount.NewCounter().NewAccount(&ncount.NewAccountReq{
				OrderID: ncount.GetMerOrderID(),
				MsgCipherText: &ncount.NewAccountMsgCipherText{
					MerUserId: id,
					Mobile:    info.Mobile,
					UserName:  info.RealName,
					CertNo:    info.IdCard,
				},
			})

			fmt.Println("accountResp", id, accountResp)

			if err != nil || accountResp.ResultCode != "0000" {
				return errors.New(fmt.Sprintf("认证失败,code:%s ,msg :%s", accountResp))
			}
			//主账户
			if id == fmt.Sprintf("%s%d", UserMainAccountPrefix, info.UserId) {
				info.MainAccountId = accountResp.UserId
			} else {
				info.PacketAccountId = accountResp.UserId
			}

			return nil
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := errGroup.Wait(); err != nil {
		fmt.Println("errGroup.Wait", err.Error())

		//return &cloud_wallet.IdCardRealNameAuthResp{
		//	Step: 0,
		//	CommonResp: &cloud_wallet.CommonResp{
		//		ErrCode: 1,
		//		ErrMsg:  err.Error(),
		//	},
		//}, nil
	}

	//更新新生支付账户id
	if len(info.MainAccountId) > 0 || len(info.PacketAccountId) > 0 {
		err = imdb.UpdateNcountAccountInfo(info)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("更新数据失败:%s", err.Error()))
		}
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
func (rpc *CloudWalletServer) SetPaymentSecret(_ context.Context, req *cloud_wallet.SetPaymentSecretReq) (*cloud_wallet.SetPaymentSecretResp, error) {
	//获取用户账户信息
	accountInfo, err := imdb.GetNcountAccountByUserId(req.UserId)
	if err != nil || accountInfo.Id <= 0 {
		return nil, errors.New("账户信息不存在")
	}

	//md5 加密密码
	m5 := md5.New()
	m5.Write([]byte(req.PaymentSecret))
	md5Data := m5.Sum([]byte(""))
	secret := hex.EncodeToString(md5Data)

	err = imdb.UpdateNcountAccountField(req.UserId, map[string]interface{}{"payment_password": secret})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("保存数据失败,err:%s", err.Error()))
	}

	return &cloud_wallet.SetPaymentSecretResp{
		Step: 2,
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
			BankCardNumber: v.BankCardNumber,
			CreatedTime:    v.CreatedTime.Format("2006-01-02 15:04:05"),
		})
	}

	return resp, err
}

// 绑定用户银行卡
func (rpc *CloudWalletServer) BindUserBankcard(ctx context.Context, req *cloud_wallet.BindUserBankcardReq) (*cloud_wallet.BindUserBankcardResp, error) {
	//获取用户账户信息
	accountInfo, err := imdb.GetNcountAccountByUserId(req.UserId)
	if err != nil || accountInfo.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("查询账户数据失败 %s,error:%s", req.UserId, err.Error()))
	}

	merOrderId := ncount.GetMerOrderID()
	accountResp, err := ncount.NewCounter().BindCard(&ncount.BindCardReq{
		MerOrderId: merOrderId,
		BindCardMsgCipherText: ncount.BindCardMsgCipherText{
			CardNo:       req.BankCardNumber,
			HolderName:   req.CardOwner,
			MobileNo:     req.Mobile,
			IdentityType: "1",
			IdentityCode: accountInfo.IdCard,
			UserId:       accountInfo.MainAccountId,
		},
	})

	log.Info(merOrderId, "accountResp", &accountResp, err, ncount.BindCardReq{
		MerOrderId: merOrderId,
		BindCardMsgCipherText: ncount.BindCardMsgCipherText{
			CardNo:       req.BankCardNumber,
			HolderName:   req.CardOwner,
			MobileNo:     req.Mobile,
			IdentityType: "1",
			IdentityCode: accountInfo.IdCard,
			UserId:       accountInfo.MainAccountId,
		},
	})
	fmt.Println("accountResp Println", accountResp, err)
	if err != nil || accountResp.ResultCode != "0000" {
		return nil, errors.New(fmt.Sprintf("绑定银行卡失败,code:"))
	}

	info := &db.FNcountBankCard{
		UserId:            req.UserId,
		MerOrderId:        merOrderId,
		NcountOrderId:     accountResp.NcountOrderId,
		NcountUserId:      accountInfo.MainAccountId,
		Mobile:            req.Mobile,
		CardOwner:         req.CardOwner,
		BankCardNumber:    req.BankCardNumber,
		Cvv2:              req.Cvv2,
		CardAvailableDate: req.CardAvailableDate,
		CreatedTime:       time.Now(),
		UpdatedTime:       time.Now(),
	}

	//数据入库
	_ = imdb.BindUserBankcard(info)

	return &cloud_wallet.BindUserBankcardResp{
		BankCardId: info.Id,
	}, nil
}

// 绑定用户银行卡确认code
func (rpc *CloudWalletServer) BindUserBankcardConfirm(_ context.Context, req *cloud_wallet.BindUserBankcardConfirmReq) (*cloud_wallet.BindUserBankcardConfirmResp, error) {
	//获取绑定的银行卡信息
	bankCardInfo, err := imdb.GetNcountBankCardById(req.BankCardId)
	if err != nil || bankCardInfo.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("查询银行卡数据失败,error:%s", err.Error()))
	}

	//已绑定
	if bankCardInfo.IsBind == 1 {
		return &cloud_wallet.BindUserBankcardConfirmResp{BankCardId: bankCardInfo.Id}, err
	}

	//新生支付确定接口
	accountResp, err := ncount.NewCounter().BindCardConfirm(&ncount.BindCardConfirmReq{
		MerOrderId: ncount.GetMerOrderID(),
		BindCardConfirmMsgCipherText: ncount.BindCardConfirmMsgCipherText{
			NcountOrderId: bankCardInfo.NcountOrderId,
			SmsCode:       req.SmsCode,
			MerUserIp:     req.MerUserIp,
		},
	})

	fmt.Println("accountResp Println", accountResp, err)
	if err != nil || accountResp.ResultCode != "0000" {
		return nil, errors.New(fmt.Sprintf("绑定用户银行卡确认,code:"))
	}

	//更新数据
	_ = imdb.UpdateNcountBankCardField(bankCardInfo.Id, map[string]interface{}{"bind_card_agr_no": accountResp.BindCardAgrNo, "is_bind": 1, "bank_code": accountResp.BankCode})

	return &cloud_wallet.BindUserBankcardConfirmResp{BankCardId: bankCardInfo.Id}, err
}

// 解绑用户银行卡
func (rpc *CloudWalletServer) UnBindingUserBankcard(_ context.Context, req *cloud_wallet.UnBindingUserBankcardReq) (*cloud_wallet.UnBindingUserBankcardResp, error) {
	//获取绑定的银行卡信息
	bankCardInfo, err := imdb.GetNcountBankCardById(req.BankCardId)
	if err != nil || bankCardInfo.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("查询银行卡数据失败,error:%s", err.Error()))
	}

	//新生支付确定接口
	accountResp, err := ncount.NewCounter().UnbindCard(&ncount.UnBindCardReq{
		MerOrderId: ncount.GetMerOrderID(),
		UnBindCardMsgCipher: ncount.UnBindCardMsgCipher{
			OriBindCardAgrN: bankCardInfo.BindCardAgrNo,
			UserId:          bankCardInfo.NcountUserId,
		},
	})

	fmt.Println("accountResp Println", accountResp, err)
	if err != nil || accountResp.ResultCode != "0000" {
		return nil, errors.New(fmt.Sprintf("解绑银行卡失败"))
	}

	//更新数据库
	_ = imdb.UnBindUserBankcard(req.BankCardId, req.UserId)

	return &cloud_wallet.UnBindingUserBankcardResp{}, err
}
