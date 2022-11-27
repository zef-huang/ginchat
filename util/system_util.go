package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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

var Db *gorm.DB

func MysqlInit() {
	mysqlDns := Config.Get("mysql.dns").(string)

	mysqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db, err := gorm.Open(mysql.Open(mysqlDns), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		panic("failed to connect database")
	}

	Db = db
	fmt.Println("连接 mysql 成功", Db)
}

var ctx = context.Background()
var redisDb *redis.Client

func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	redisDb = rdb
	fmt.Println("连接 redis 成功", redisDb)
}
