package main

import (
	docs "ginchat/docs"
	"ginchat/router"
	"ginchat/util"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	util.ConfigInit()
	util.MysqlInit()
	r := router.Route()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//test.UserTest()
}
