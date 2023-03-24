package cloud_wallet

import (
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/log"
	promePkg "Open_IM/pkg/common/prometheus"
	"Open_IM/pkg/grpc-etcdv3/getcdv3"
	"Open_IM/pkg/proto/cloud_wallet"
	"Open_IM/pkg/utils"
	"context"
	"net"
	"strconv"
	"strings"

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

// 获取用户余额
//func (rpc *CloudWalletServer) UserAccountBalance(ctx context.Context, req *cloud_wallet.SetPaymentSecretReq) (*cloud_wallet.UserAccountBalanceResp, error) {
//	return &cloud_wallet.UserAccountBalanceResp{
//		CommonResp: &cloud_wallet.CommonResp{
//			ErrCode: 1,
//			ErrMsg:  "not",
//		},
//	}, nil
//}

func (rpc *CloudWalletServer) UserAccountBalance(ctx context.Context, req *cloud_wallet.UserAccountBalanceReq) (*cloud_wallet.UserAccountBalanceResp, error) {
	return &cloud_wallet.UserAccountBalanceResp{
		MainBalance: 100,
	}, nil
}

func (rpc *CloudWalletServer) mustEmbedUnimplementedCloudWalletServer() {
	//TODO implement me
	return
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
func (rpc *CloudWalletServer) IdCardRealNameAuth(ctx context.Context, req *cloud_wallet.IdCardRealNameAuthReq) (*cloud_wallet.IdCardRealNameAuthResp, error) {
	return &cloud_wallet.IdCardRealNameAuthResp{
		Step: 1,
		CommonResp: &cloud_wallet.CommonResp{
			ErrCode: 1,
			ErrMsg:  "not",
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
