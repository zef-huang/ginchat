package service

import (
	"fmt"
	"ginchat/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags 用户数据
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /userList [get]
func UserService(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println("input", pageSize)

	pageData := model.GetUserPageData(page, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"message": pageData,
	})
}

type CreateUserParams struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
	Phone    string `json:"Phone" binding:"required"`
}

func CreateUser(c *gin.Context) {
	fmt.Println("enter")

	params := CreateUserParams{}
	c.ShouldBindJSON(&params)

	fmt.Println("params", params)

	user := model.UserBasic{}

	user.UserName = params.Username
	user.PassWord = params.Password
	user.Phone = params.Phone

	model.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

type DeleteUserParams struct {
	Username string `json:"Username" binding:"required"`
}

func DeleteUser(c *gin.Context) {
	fmt.Println("enter")

	params := DeleteUserParams{}
	c.ShouldBindJSON(&params)

	fmt.Println("params", params)

	user := model.UserBasic{}

	user.UserName = params.Username

	model.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func UpdateUser(c *gin.Context) {
	fmt.Println("enter")

	params := CreateUserParams{}
	c.ShouldBindJSON(&params)

	fmt.Println("params", params)

	user := model.UserBasic{}

	user.UserName = params.Username
	user.PassWord = params.Password
	user.Phone = params.Phone

	model.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
