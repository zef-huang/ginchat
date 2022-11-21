package test

import (
	"ginchat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func GormTest() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/ginchat"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	user := &model.UserBasic{
		UserName: "hzf",
	}

	// Migrate the schema
	db.AutoMigrate(user)

	// Create
	db.Create(user)
}
