package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
}

type Forum struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Description string
	CreatedAt   time.Time
}

type UserFollowForum struct {
	UserID     uint `gorm:"primaryKey"`
	ForumID    uint `gorm:"primaryKey"`
	FollowedAt time.Time
}

type Post struct {
	ID        uint `gorm:"primaryKey"`
	ForumID   uint
	UserID    uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostLike struct {
	PostID  uint `gorm:"primaryKey"`
	UserID  uint `gorm:"primaryKey"`
	LikedAt time.Time
}
