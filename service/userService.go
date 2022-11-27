package service

import (
	"fmt"
	"ginchat/model"
	"ginchat/util"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	Phone    string `json:"Phone" binding:"required" valid:"^1[1-9]{1}\d{9}$"`
}

func CreateUser(c *gin.Context) {
	fmt.Println("enter")

	params := CreateUserParams{}
	c.ShouldBindJSON(&params)

	fmt.Println("params", params)

	user := model.UserBasic{}

	user.UserName = params.Username
	user.PassWord = util.Md5Password(params.Password)
	user.Phone = params.Phone

	resultUser := model.FindUserByName(user.UserName)
	if resultUser.UserName == user.UserName {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名已存在",
		})
	}

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
	result, err := govalidator.ValidateStruct(params)
	if result == false {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

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

type UserLoginParams struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func GetUserInfo(c *gin.Context) {
	fmt.Println("enter")

	params := UserLoginParams{}
	c.ShouldBindJSON(&params)

	fmt.Println("params", params)

	user := model.UserBasic{}

	user.UserName = params.Username
	user.PassWord = util.Md5Password(params.Password)

	findUser := model.FindUserByName(user.UserName)

	if findUser.UserName == user.UserName && findUser.PassWord == user.PassWord {
		token := util.SignJwt(user.UserName)
		user.Identity = token
		model.UpdateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "获取 token 成功",
			"token":   token,
		})

		fmt.Println("result", util.ParseJwt(token))
		time.Sleep(2 * time.Second)
		fmt.Println("result", util.ParseJwt(token))
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名密码错误",
		})
	}
}
