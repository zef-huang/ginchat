package model

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	UserName string
	PassWord string
	Email    string
	Phone    string
}
