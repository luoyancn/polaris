package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/dubhe/registry/etcdv3"
	"github.com/luoyancn/dubhe/registry/etcdv3/config"
)

func main() {
	go func() {
		etcdcfg := config.NewEtcdConf()
		common.ReadConfig(
			"", "test", logging.STD_ENABLED, "", true, true, etcdcfg)
		_ = etcdv3.Register("127.0.0.1:8080")
	}()
	common.Wait(etcdv3.UnRegister)
}
