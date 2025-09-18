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
	// sslmode=disable 禁用ssl加密连接（仅开发测试环境可用，生产环境应更改） TimeZone=Asia/Shanghai 将数据库会话的时区设置为Asia/Shanghai（北京时间）
	dsn := "host=mypg user=postgres password=123456 dbname=blogdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	DB = db
	fmt.Println("数据库已连接")

	db.AutoMigrate(&model.Class{}, &model.Article{})
}
