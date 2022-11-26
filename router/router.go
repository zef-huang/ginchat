package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.IndexService)

	r.GET("/userList", service.UserService)
	r.POST("/createUser", service.CreateUser)
	r.POST("/deleteUser", service.DeleteUser)
	r.POST("/updateUser", service.UpdateUser)
	return r
}
