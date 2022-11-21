package test

import (
	"fmt"
	"ginchat/model"
)

func UserTest() {
	fmt.Println("start test")

	pageData := model.GetUserPageData(1, 20)
	fmt.Println("pageData", pageData)
}
