package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	"github.com/luoyancn/dubhe/logging"
)

func main() {
	logging.GetLogger("test", logging.STD_ENABLED, "", true)
	entity := grpclib.NewServiceDescKV(
		&msg.Service{HostName: "luoyan", ListenPort: 8080},
		msg.Messages_serviceDesc)
	go grpclib.StartServer(8080, nil, nil, entity)
	common.Wait(grpclib.StopServer)
}
