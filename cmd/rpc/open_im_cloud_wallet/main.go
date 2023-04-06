package main

import (
	"Open_IM/internal/rpc/cloud_wallet"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	promePkg "Open_IM/pkg/common/prometheus"
	"flag"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImCloudWalletPort
	log.NewPrivateLog("open_im_cloud_wallet")
	rpcPort := flag.Int("port", defaultPorts[0], "rpc listening port")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.CloudWalletPrometheusPort[0], "CloudWalletPrometheusPort default listen port")
	flag.Parse()
	rpcServer := cloud_wallet.NewRpcCloudWalletServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
