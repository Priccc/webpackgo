package main

import (
	"github.com/gin-gonic/gin"
	"{{ name }}/common"
	rout "{{ name }}/router"
	"flag"
)

var (
	configFile    = flag.String("config", "conf/config.json", "config file for system")
	config     common.IConfigInfo

)

func main() {

	flag.Parse()
	if flag.Parsed() == false {
		flag.PrintDefaults()
		return
	}

	conf, err := common.LoadConfig(*configFile)
	if err != nil {
		panic("配置文件加载错误: " + err.Error())
	}

	config = conf

	router := gin.Default()
	rout.InitRouters(router)

	router.Run(":" + conf.GetListeningPort())
}
