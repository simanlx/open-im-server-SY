package main

import (
	"Open_IM/internal/rpc/agent"
	"Open_IM/pkg/common/config"
	promePkg "Open_IM/pkg/common/prometheus"
	"flag"
	"fmt"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImAgentPort
	rpcPort := flag.Int("port", defaultPorts[0], "rpc listening port")
	prometheusPort := flag.Int("prometheus_port", config.Config.Prometheus.AgentPrometheusPort[0], "AgentPrometheusPort default listen port")
	flag.Parse()
	rpcServer := agent.NewRpcAgentServer(*rpcPort)
	fmt.Println("start open_im_agent server, address: ", *rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
