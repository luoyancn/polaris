package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	"github.com/luoyancn/dubhe/grpclib/config"
	"github.com/luoyancn/dubhe/logging"
)

func main() {
	grpcfg := config.NewGrpConf()
	common.ReadConfig(
		"ssl.toml", "test", logging.STD_ENABLED, "", true, true, grpcfg)
	entity := grpclib.NewServiceDescKV(
		&msg.Service{HostName: "luoyan", ListenPort: config.GRPC_PORT},
		msg.Messages_serviceDesc)
	go grpclib.StartServer(config.GRPC_PORT, nil, nil, entity)
	common.Wait(grpclib.StopServer)
}
