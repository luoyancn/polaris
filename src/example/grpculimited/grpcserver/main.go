package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	"github.com/luoyancn/dubhe/logging"
)

func main() {
	logging.GetLogger("test", logging.STD_ENABLED, "", true)
	go grpclib.StartServer(
		8080, &msg.Service{HostName: "zhangjl", ListenPort: 8080}, nil, nil,
		msg.Messages_serviceDesc)
	common.Wait(grpclib.StopServer)
}
