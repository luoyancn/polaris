package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	grpconfig "github.com/luoyancn/dubhe/grpclib/config"
	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/dubhe/registry/etcdv3"
	etcdconfig "github.com/luoyancn/dubhe/registry/etcdv3/config"
)

func main() {

	grpcfg := grpconfig.NewGrpConf()
	etcdcfg := etcdconfig.NewEtcdConf()
	common.ReadConfig("ssl.toml", "test", logging.STD_ENABLED,
		"", true, true, grpcfg, etcdcfg)
	go grpclib.StartServer(
		grpconfig.GRPC_PORT, &msg.Service{HostName: "luoyan",
			ListenPort: grpconfig.GRPC_PORT},
		etcdv3.Register, etcdv3.UnRegister,
		msg.Messages_serviceDesc)
	common.Wait(grpclib.StopServer)
}
