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
