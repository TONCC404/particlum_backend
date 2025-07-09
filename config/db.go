package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"particlumn_backend/model"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user:password@tcp(localhost:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 自动迁移
	db.AutoMigrate(&model.User{}, &model.Forum{}, &model.UserFollowForum{}, &model.Post{}, &model.PostLike{})
	DB = db
}
