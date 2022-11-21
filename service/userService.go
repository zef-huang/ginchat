package service

import (
	"fmt"
	"ginchat/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserService(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	fmt.Println("input", pageSize)

	pageData := model.GetUserPageData(page, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"message": pageData,
	})
}
