package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
	router.Static("/static", "assets/static")

	router.LoadHTMLGlob("assets/*.html")

	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":" + *listeningPort)
}
