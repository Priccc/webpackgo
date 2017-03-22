package main

import (
	"github.com/gin-gonic/gin"
	rout "{{ name }}/router"
	"flag"
)

var (
	configFile    = flag.String("config", "conf/config.json", "config file for system")
	listeningPort = flag.String("port", "8800", "listeningPort")

)

func main() {

	flag.Parse()
	if flag.Parsed() == false {
		flag.PrintDefaults()
		return
	}

	router := gin.Default()
	rout.InitRouters(router)

	router.Run(":" + *listeningPort)
}
