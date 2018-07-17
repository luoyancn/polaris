package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/dubhe/registry/etcdv3/config"
)

func main() {
	common.ReadConfig("ssl.toml", "test", logging.STD_ENABLED,
		"", true, true, config.NewEtcdConf())
}
