package model

import (
	"fmt"
	"ginchat/util"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	UserName string
	PassWord string
	Email    string
	Phone    string
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
