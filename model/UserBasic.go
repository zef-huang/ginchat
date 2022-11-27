package model

import (
	"fmt"
	"ginchat/util"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	UserName string `json:"UserName" binding:"required" `
	PassWord string `json:"PassWord" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Phone    string `json:"Phone" binding:"required"`
}

func (table *UserBasic) tableName() string {
	return "user_basic"
}

func GetUserPageData(page int, pageSize int) []*UserBasic {
	var users []*UserBasic
	util.Db.Table("user_basics").Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)
	fmt.Println(users)

	return users
}

func CreateUser(user UserBasic) {
	fmt.Println(user)

	util.Db.Table("user_basics").Create(&user)
}

func DeleteUser(user UserBasic) {
	fmt.Println(user)

	util.Db.Table("user_basics").Delete(&user)
}

func UpdateUser(user UserBasic) {
	fmt.Println("UpdateUser", user)

	cond := UserBasic{UserName: user.UserName}

	util.Db.Table("user_basics").Where(&cond).Updates(&user)
}

func FindUserByName(username string) UserBasic {
	fmt.Println("username", username)

	cond := UserBasic{UserName: username}

	user := UserBasic{}
	util.Db.Table("user_basics").Where(&cond).Find(&user)

	return user
}