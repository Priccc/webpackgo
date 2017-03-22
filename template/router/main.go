package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {
	initHtml(router)
	initStatic(router)
}
