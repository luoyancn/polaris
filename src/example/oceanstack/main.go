package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/logging"

	"github.com/luoyancn/ocean/config"
	"github.com/luoyancn/ocean/keystone"
	"github.com/luoyancn/ocean/nova"
)

func main() {
	cfg := config.NewRdoConf()
	common.ReadConfig("rdo.toml", "rdo", logging.FILE_ENABLED,
		"rdo", true, true, cfg)
	ctx, token, _ := keystone.Authorization()
	logging.LOG.Debugf("The auth token is: %s\n", token)
	nova.ForceDownAndComputeSrv(ctx, token, "kubernetes3", true)
	nova.EvacuateVmOnHost(ctx, token, "kubernetes3")
}
