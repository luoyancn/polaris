package main

import (
	"context"
	"time"

	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/fake/mail"
	"github.com/luoyancn/dubhe/fake/msg"
	"github.com/luoyancn/dubhe/grpclib"
	"github.com/luoyancn/dubhe/grpclib/config"
	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/dubhe/registry/etcdv3"
	etcdconfig "github.com/luoyancn/dubhe/registry/etcdv3/config"
)

func main() {
	grpcfg := config.NewGrpConf()
	etcdcfg := etcdconfig.NewEtcdConf()
	common.ReadConfig(
		"ssl.toml", "test", logging.STD_ENABLED,
		"", true, true, grpcfg, etcdcfg)
	grpclib.InitGrpcClientPool("", etcdv3.NewResolver)
	conn := grpclib.Get()
	defer grpclib.Put(conn)
	client := msg.NewMessagesClient(conn)
	//for {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	resp, err := client.Call(ctx, &msg.Request{Req: "luoyan"})
	if nil != err {
		logging.LOG.Errorf(
			"Failed to get response from grpc server:%v\n", err)
		return
	}
	//time.Sleep(5 * time.Microsecond)
	logging.LOG.Infof("The response of grpc server is %s\n", resp.GetResp())
	//}

	mail_receiver := mail.NewMailClient(conn)
	reply, err := mail_receiver.Call(ctx, &mail.Sender{Content: "hello"})
	if nil != err {
		logging.LOG.Errorf(
			"Failed to get response from grpc server:%v\n", err)
		return
	}
	logging.LOG.Infof("The response of grpc server is %s\n", reply.GetReply())

}
