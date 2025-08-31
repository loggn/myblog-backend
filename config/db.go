package config

import (
	"fmt"
	"log"
	"myblog-backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=blogdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	DB = db
	fmt.Println("数据库已连接")

	db.AutoMigrate(&model.Class{}, &model.Article{})
}
