package main

import (
	"github.com/luoyancn/dubhe/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	"github.com/luoyancn/dubhe/logging"
)

func main() {
	logging.GetLogger("test", logging.STD_ENABLED, "", true)
	grpclib.StartServer(
		8080, &msg.Service{HostName: "zhangjl", ListenPort: 8080},
		msg.Messages_serviceDesc)
}
