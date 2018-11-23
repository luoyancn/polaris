package main

import (
	pipes "github.com/ebuchman/go-shell-pipes"

	"github.com/luoyancn/dubhe/common"
	"github.com/luoyancn/dubhe/logging"

	"github.com/luoyancn/ocean/config"
	"github.com/luoyancn/ocean/keystone"
	"github.com/luoyancn/ocean/nova"
)

func main() {

	cfg := config.NewRdoConf()
	common.ReadConfig("rdo.toml", "rdo", logging.FILE_ENABLED,
		"/tmp", true, true, cfg)
	output, err := pipes.RunString(`cat /tmp/rdo | awk '{print $1}'| grep 2012`)
	logging.LOG.Infof("Shell cmd result is %s\n", output)

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
