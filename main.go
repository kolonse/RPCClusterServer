// RPCClusterServer project RPCClusterServer.go
package main

import (
	"encoding/json"
	"github.com/kolonse/TCPServer"
	"github.com/kolonse/kolonsecfg"
	"github.com/kolonse/logs"
	"os"
	"path"
)

func main() {
	cfg := kolonsecfg.NewCfg()
	cfg.ParseFile(*cfgPath)
	cfgNode := cfg.Child(*env)
	logger := logs.NewLogger(10000)
	// 创建配置日志文件目录
	os.MkdirAll(cfgNode.Child("log.logdir").GetString(), 0777)
	filename := path.Join(cfgNode.Child("log.logdir").GetString(),
		cfgNode.Child("log.filename").GetString())
	logCfg := map[string]interface{}{
		"filename": filename,
		"maxlines": cfgNode.Child("log.maxlines").GetInt(),
		"maxsize":  cfgNode.Child("log.maxsize").GetInt(),
		"daily":    cfgNode.Child("log.daily").GetBool(),
		"maxdays":  cfgNode.Child("log.maxdays").GetInt(),
		"rotate":   cfgNode.Child("log.rotate").GetBool(),
	}
	buff, _ := json.Marshal(logCfg)
	err := logger.SetLogger("file", string(buff))
	if err != nil {
		panic(err.Error())
	}
	err = logger.SetLogger("console", "")
	if err != nil {
		panic(err.Error())
	}
	logger.Info("当前环境:%v 加载配置:\n%v", *env, cfgNode.Dump(""))
	ts := TCPServer.NewTCPServer(cfgNode.Child("serveraddr").GetString())
	ts.Register("logger", logger)
	ts.Server()
}
