package main

import (
	"context"
	"fmt"
	"time"

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
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	grpclib.InitGrpcClientPool("127.0.0.1:8080", nil)
	conn := grpclib.Get()
	defer grpclib.Put(conn)
	client := msg.NewMessagesClient(conn)
	resp, err := client.Call(ctx, &msg.Request{Req: "luoyan"})
	if nil != err {
		fmt.Printf("Failed to get response from grpc server:%v\n", err)
		return
	}
	fmt.Printf("%s\n", resp.GetResp())
}
