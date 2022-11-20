package util

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Config = viper.New()

func ConfigInit() {

	Config.SetConfigName("app") // name of config file (without extension)
	Config.AddConfigPath("./conf")
	err := Config.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

var db gorm.Dialector

func MysqlInit() {
	mysqlDns := Config.Get("mysql.dns").(string)
	db, err := gorm.Open(mysql.Open(mysqlDns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("连接 mysql 成功", db)
}
