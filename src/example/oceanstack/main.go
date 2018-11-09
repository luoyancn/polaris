package main

import (
	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/merak"

	"github.com/luoyancn/ocean/config"
	"github.com/luoyancn/ocean/keystone"
	"github.com/luoyancn/ocean/nova"
)

func main() {

	merak.Exec([]string{"ps", "-efl", "|", "grep", "hugo"})
	cfg := config.NewRdoConf()
	common.ReadConfig("rdo.toml", "rdo", logging.FILE_ENABLED,
		"/tmp", true, true, cfg)
	ctx, token, err := keystone.Authorization()
	if nil != err {
		return
	}
	logging.LOG.Debugf("The auth token is: %s\n", token)
	err = nova.ForceDownAndComputeSrv(ctx, token, "kubernetes3", true)
	if nil != err {
		return
	}
	nova.EvacuateVmOnHost(ctx, token, "kubernetes3")
}
