package main

import (
	"ginchat/router"
	"ginchat/util"
)

func main() {
	util.ConfigInit()
	util.MysqlInit()
	r := router.Route()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//test.UserTest()
}
