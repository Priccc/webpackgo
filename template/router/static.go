package router

import (
	"github.com/gin-gonic/gin"
)

func initStatic(router *gin.Engine) {
	router.Static("/static", "./assets/static")
}
