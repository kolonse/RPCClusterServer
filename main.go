// RPCClusterServer project RPCClusterServer.go
package main

import (
	"github.com/kolonse/kolonsecfg"
	"github.com/kolonse/logs"
)

func main() {
	cfg := kolonsecfg.NewCfg()
	cfg.ParseFile(*cfgPath)
	logger := logs.NewLogger(10000)
	//	{
	//		"filename":"logs/beego.log",
	//		"maxlines":10000,
	//		"maxsize":1<<30,
	//		"daily":true,
	//		"maxdays":15,
	//		"rotate":true
	//	}
	err := logger.SetLogger("file", "")
	if err != nil {
		panic(err.Error())
	}
	logger.Info("%v", cfg.Dump())
}
