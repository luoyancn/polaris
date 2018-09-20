package main

import (
	"errors"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/luoyancn/dubhe/logging"
	"github.com/luoyancn/merak"
)

var once sync.Once
var url string
var fromfile string
var proxy string = ""
var paiallel int8 = 1

var rootcmd = &cobra.Command{
	Short:   "Tools for Donwload file from web",
	Long:    ` The commands aims to download file from static web server`,
	Run:     download,
	Args:    cobra.NoArgs,
	PreRunE: mutually,
}

func init() {
	rootcmd.PersistentFlags().StringVarP(
		&fromfile, "recordefile", "f", "",
		"The files to record all download links, required")
	rootcmd.PersistentFlags().StringVarP(
		&url, "url", "u", "", "The single download link, required")
	rootcmd.PersistentFlags().StringVarP(
		&proxy, "proxy", "p", "",
		"Using http proxy to download, link http://localhost:1234")
	rootcmd.PersistentFlags().Int8VarP(
		&paiallel, "paiallel", "n", 1,
		"The number of download at the same time, default is 1. Only used with params recordefile")
}

func mutually(cmd *cobra.Command, args []string) error {
	_url := cmd.Flags().Lookup("url").Value.String()
	_file := cmd.Flags().Lookup("recordefile").Value.String()
	if "" == _url && "" == _file {
		return errors.New("Flag url or recordefile must be provide\n")
	}
	if "" != _url && "" != _file {
		return errors.New("Only one of url or recordefile could be recognised at the same time\n")
	}
	return nil
}

func download(cmd *cobra.Command, args []string) {
	if "" != url {
		merak.Download(url, proxy)
	} else {
		urls, err := merak.ReadLines(fromfile)
		if nil != err {
			logging.LOG.Errorf("ERROR:%v\n", err)
			return
		}
		downloads := make(chan struct{}, paiallel)
		defer close(downloads)
		for _, url := range urls {
			go func(url string) {
				logging.LOG.Infof("Download file from %s\n", url)
				err := merak.Download(url, proxy)
				downloads <- struct{}{}
				if nil != err {
					logging.LOG.Infof("%s Download failed\n", url)
				} else {
					logging.LOG.Infof("%s Download completed\n", url)
				}
			}(url)
		}

		for _ = range urls {
			<-downloads
		}
	}
}

func main() {
	//logging.GetLogger("", logging.STD_ENABLED, "", false)
	if err := rootcmd.Execute(); nil != err {
		os.Exit(1)
	}
}
