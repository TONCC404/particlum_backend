package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"particlumn_backend/model"
)

var DB *gorm.DB

func InitDB() {
	log.Println("InitDB....")
	if err := godotenv.Load(); err != nil {
		panic("failed to load .env file: " + err.Error())
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to PostgreSQL: " + err.Error())
	}
	if err := db.AutoMigrate(&model.User{}, &model.Forum{}, &model.UserFollowForum{}, &model.Post{}, &model.PostLike{}); err != nil {
		panic("auto migration failed: " + err.Error())
	}

	DB = db
}
