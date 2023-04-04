package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
	"Open_IM/pkg/common/config"
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
	"encoding/json"
	"errors"
	"fmt"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
	"time"
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
		count:           ncount.NewCounter(),
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
func (rpc *CloudWalletServer) UserNcountAccount(_ context.Context, req *cloud_wallet.UserNcountAccountReq) (*cloud_wallet.UserNcountAccountResp, error) {
	//获取用户账户信息
	accountInfo, err := imdb.GetNcountAccountByUserId(req.UserId)
	if err != nil || accountInfo.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("查询账户数据失败 %v,error:%v", req.UserId, err.Error()))
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
	if err != nil {
		return nil, errors.New(fmt.Sprintf("查询账户信息失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return nil, errors.New(fmt.Sprintf("查询账户信息失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	//绑定的银行卡列表
	bindCardsList := make([]*cloud_wallet.BindCardsList, 0)
	if len(accountResp.BindCardAgrNoList) > 0 {
		bindCards := make([]ncount.NAccountBankCard, 0)
		err = json.Unmarshal([]byte(accountResp.BindCardAgrNoList), &bindCards)
		if err == nil {
			for _, v := range bindCards {
				bindCardsList = append(bindCardsList, &cloud_wallet.BindCardsList{
					BankCode:      v.BankCode,
					CardNo:        v.CardNo,
					BindCardAgrNo: v.BindCardAgrNo,
				})
			}
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

	//组装数据
	info := &db.FNcountAccount{
		UserId:      req.UserId,
		Mobile:      req.Mobile,
		RealName:    req.RealName,
		IdCard:      req.IdCard,
		OpenStep:    1,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	//调新生支付接口-开户
	errGroup := new(errgroup.Group)
	accountIds := []string{
		fmt.Sprintf("%s%s", UserMainAccountPrefix, info.UserId),
		fmt.Sprintf("%s%s", UserPacketAccountPrefix, info.UserId),
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
			log.Info("", "accountResp", &accountResp, err)

			if err != nil {
				return errors.New(fmt.Sprintf("实名认证失败(%s)", err.Error()))
			} else {
				if accountResp.ResultCode != "0000" {
					return errors.New(fmt.Sprintf("实名认证失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
				}
			}

			//主账户
			if id == fmt.Sprintf("%s%s", UserMainAccountPrefix, info.UserId) {
				info.MainAccountId = accountResp.UserId
			} else {
				info.PacketAccountId = accountResp.UserId
			}

			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return nil, err
	}

	//实名数据入库
	err = imdb.CreateNcountAccount(info)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("数据入库失败:%s", err.Error()))
	}

	return &cloud_wallet.IdCardRealNameAuthResp{
		Step: 1,
		CommonResp: &cloud_wallet.CommonResp{
			ErrCode: 0,
			ErrMsg:  "实名认证成功",
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

	err = imdb.UpdateNcountAccountField(req.UserId, map[string]interface{}{"payment_password": secret, "open_step": 2})
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

// 绑定用户银行卡
func (rpc *CloudWalletServer) BindUserBankcard(_ context.Context, req *cloud_wallet.BindUserBankcardReq) (*cloud_wallet.BindUserBankcardResp, error) {
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
	if err != nil {
		return nil, errors.New(fmt.Sprintf("绑定银行卡失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return nil, errors.New(fmt.Sprintf("绑定银行卡失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
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
	bankCardInfo, err := imdb.GetNcountBankCardById(req.BankCardId, req.UserId)
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
	if err != nil {
		return nil, errors.New(fmt.Sprintf("绑定用户银行卡确认失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return nil, errors.New(fmt.Sprintf("绑定用户银行卡确认失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	//更新数据
	_ = imdb.BindUserBankcardConfirm(bankCardInfo.Id, req.UserId, accountResp.BindCardAgrNo, accountResp.BankCode)

	return &cloud_wallet.BindUserBankcardConfirmResp{BankCardId: bankCardInfo.Id}, err
}

// 解绑用户银行卡
func (rpc *CloudWalletServer) UnBindingUserBankcard(_ context.Context, req *cloud_wallet.UnBindingUserBankcardReq) (*cloud_wallet.UnBindingUserBankcardResp, error) {
	//获取绑定的银行卡信息
	bankCardInfo, err := imdb.GetNcountBankCardByBindCardAgrNo(req.BindCardAgrNo, req.UserId)
	if err != nil || bankCardInfo.Id <= 0 {
		return nil, errors.New(fmt.Sprintf("查询银行卡数据失败,error:%s", err.Error()))
	}

	//新生支付确定接口
	accountResp, err := ncount.NewCounter().UnbindCard(&ncount.UnBindCardReq{
		MerOrderId: ncount.GetMerOrderID(),
		UnBindCardMsgCipher: ncount.UnBindCardMsgCipher{
			OriBindCardAgrNo: bankCardInfo.BindCardAgrNo,
			UserId:           bankCardInfo.NcountUserId,
		},
	})
	fmt.Println("accountResp Println", accountResp, err, bankCardInfo.BindCardAgrNo, bankCardInfo.NcountUserId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解绑银行卡失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return nil, errors.New(fmt.Sprintf("解绑银行卡失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	//更新数据库
	_ = imdb.UnBindUserBankcard(bankCardInfo.Id, req.UserId)

	return &cloud_wallet.UnBindingUserBankcardResp{}, err
}

// 银行卡充值
func (c *CloudWalletServer) UserRecharge(_ context.Context, req *cloud_wallet.UserRechargeReq) (*cloud_wallet.UserRechargeResp, error) {
	// 获取银行卡信息
	bankCardInfo, err := imdb.GetNcountBankCardByBindCardAgrNo(req.BindCardAgrNo, req.UserId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("获取银行卡信息失败%s", err.Error()))
	}

	//充值支付
	accountResp, err := ncount.NewCounter().QuickPayOrder(&ncount.QuickPayOrderReq{
		MerOrderId: ncount.GetMerOrderID(),
		QuickPayMsgCipher: ncount.QuickPayMsgCipher{
			PayType:       "3", //绑卡协议号充值
			TranAmount:    cast.ToString(req.Amount),
			NotifyUrl:     config.Config.Ncount.Notify.RechargeNotifyUrl,
			BindCardAgrNo: bankCardInfo.BindCardAgrNo,
			ReceiveUserId: bankCardInfo.NcountUserId, //收款账户
			UserId:        bankCardInfo.NcountUserId,
			SubMerchantId: "2206301126073014978", // 子商户编号
		}})

	fmt.Println("accountResp Println", accountResp, err)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("充值失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode != "0000" {
			return nil, errors.New(fmt.Sprintf("充值失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	info := &db.FNcountTrade{
		UserID:          bankCardInfo.UserId,
		PaymentPlatform: 4,
		Type:            imdb.TradeTypeCharge,
		Amount:          req.Amount * 100, //分
		BeferAmount:     0,
		AfterAmount:     0,
		ThirdOrderNo:    accountResp.NcountOrderId,
	}

	//数据入库
	err = imdb.FNcountTradeCreateData(info)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("充值数据入库失败(%s)", err.Error()))
	}

	return &cloud_wallet.UserRechargeResp{
		OrderNo: accountResp.NcountOrderId,
	}, nil
}

// 账户充值code 确认
func (c *CloudWalletServer) UserRechargeConfirm(_ context.Context, req *cloud_wallet.UserRechargeConfirmReq) (*cloud_wallet.UserRechargeConfirmResp, error) {
	// 获取记录信息
	tradeInfo, err := imdb.GetFNcountTradeByOrderNo(req.MerOrderId, req.UserId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("获取充值记录信息失败%s", err.Error()))
	}

	//新生支付确认接口
	accountResp, err := ncount.NewCounter().QuickPayConfirm(&ncount.QuickPayConfirmReq{
		MerOrderId: ncount.GetMerOrderID(),
		QuickPayConfirmMsgCipher: ncount.QuickPayConfirmMsgCipher{
			NcountOrderId:        tradeInfo.ThirdOrderNo,
			SmsCode:              req.SmsCode,
			PaymentTerminalInfo:  "02|AA01BB",
			ReceiverTerminalInfo: "01|00001|CN|469023",
			DeviceInfo:           "192.168.0.1|E1E2E3E4E5E6|123456789012345|20000|898600MFSSYYGXXXXXXP|H1H2H3H4H5H6|AABBCC",
		},
	})

	fmt.Println("accountResp Println", accountResp, err)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("充值确认失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode == "4444" {
			return nil, errors.New(fmt.Sprintf("充值确认失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	return &cloud_wallet.UserRechargeConfirmResp{}, nil
}

// 提现
func (w *CloudWalletServer) UserWithdrawal(_ context.Context, req *cloud_wallet.DrawAccountReq) (*cloud_wallet.DrawAccountResp, error) {
	// 获取银行卡信息
	bankCardInfo, err := imdb.GetNcountBankCardByBindCardAgrNo(req.BindCardAgrNo, req.UserId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("获取银行卡信息失败%s", err.Error()))
	}

	//调用新生支付提现接口
	accountResp, err := ncount.NewCounter().Withdraw(&ncount.WithdrawReq{
		MerOrderID: ncount.GetMerOrderID(),
		MsgCipher: ncount.WithdrawMsgCipher{
			BusinessType:    "08",
			TranAmount:      req.Amount,
			UserId:          bankCardInfo.NcountUserId,
			BindCardAgrNo:   req.BindCardAgrNo,
			NotifyUrl:       config.Config.Ncount.Notify.WithdrawNotifyUrl,
			PaymentTerminal: "02|AA01BB",
			DeviceInfo:      "192.168.0.1|E1E2E3E4E5E6|123456789012345|20000|898600MFSSYYGXXXXXXP|H1H2H3H4H5H6|AABBCC",
		},
	})

	fmt.Println("accountResp Println", accountResp, err)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("提现失败(%s)", err.Error()))
	} else {
		if accountResp.ResultCode == "4444" {
			return nil, errors.New(fmt.Sprintf("提现失败 (%s,%s)", accountResp.ErrorCode, accountResp.ErrorMsg))
		}
	}

	info := &db.FNcountTrade{
		UserID:          bankCardInfo.UserId,
		PaymentPlatform: 1,
		Type:            imdb.TradeTypeWithdraw,
		Amount:          cast.ToInt32(req.Amount) * 100, //分
		BeferAmount:     0,
		AfterAmount:     0,
		ThirdOrderNo:    accountResp.NcountOrderID,
		CreatedTime:     time.Now(),
		UpdatedTime:     time.Now(),
	}

	//数据入库
	_ = imdb.FNcountTradeCreateData(info)
	return &cloud_wallet.DrawAccountResp{
		OrderNo: accountResp.NcountOrderID,
	}, nil
}
