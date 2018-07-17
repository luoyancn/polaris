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
			"ssl.toml", "test", logging.STD_ENABLED, "", true, true, etcdcfg)
		_ = etcdv3.Register("192.168.137.30:8080")
	}()
	common.Wait(etcdv3.UnRegister)
}
