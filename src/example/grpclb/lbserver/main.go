package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/grpclib"
	grpconfig "github.com/luoyancn/dubhe/grpclib/config"
	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/dubhe/registry/etcdv3"
	etcdconfig "github.com/luoyancn/dubhe/registry/etcdv3/config"
	"github.com/luoyancn/fake/mail"
	"github.com/luoyancn/fake/msg"
)

func main() {

	grpcfg := grpconfig.NewGrpConf()
	etcdcfg := etcdconfig.NewEtcdConf()
	common.ReadConfig("ssl.toml", "test", logging.STD_ENABLED,
		"", true, true, grpcfg, etcdcfg)
	entity := grpclib.NewServiceDescKV(
		&msg.Service{HostName: "luoyan", ListenPort: grpconfig.GRPC_PORT},
		msg.Messages_serviceDesc)
	mail_entity := grpclib.NewServiceDescKV(
		&mail.Email{HostName: "zhangjl", ListenPort: grpconfig.GRPC_PORT},
		mail.Mail_serviceDesc)
	go grpclib.StartServer(
		grpconfig.GRPC_PORT,
		etcdv3.Register, etcdv3.UnRegister,
		entity, mail_entity)
	common.Wait(grpclib.StopServer)
}
