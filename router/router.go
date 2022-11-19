package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.IndexService)
	return r
}
