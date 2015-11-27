// flag
package main

import (
	"flag"
)

var env = flag.String("env", "dev", "环境配置")
var cfgPath = flag.String("cfgPath", "./config/config.kcfg", "配置文件路径")

func init() {
	flag.Parse()
}
