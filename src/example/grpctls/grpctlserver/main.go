package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	"github.com/luoyancn/dubhe/grpclib/config"
	"github.com/luoyancn/dubhe/logging"
)

func main() {
	grpcfg := config.NewGrpConf()
	common.ReadConfig(
		"ssl.toml", "test", logging.STD_ENABLED, "", true, true, grpcfg)
	go grpclib.StartServer(
		config.GRPC_PORT, &msg.Service{HostName: "zhangjl",
			ListenPort: config.GRPC_PORT}, nil, nil,
		msg.Messages_serviceDesc)
	common.Wait(grpclib.StopServer)
}
