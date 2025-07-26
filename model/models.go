package model

import "time"

type PersonalInfo struct {
	Role       string   `json:"role"`
	Industry   string   `json:"industry"`
	Company    string   `json:"company"`
	Experience string   `json:"experience"`
	Goals      string   `json:"goals"`
	Interests  []string `json:"interests"`
	Bio        string   `json:"bio"`
}

type User struct {
	Username     string
	Userid       string
	Email        string
	PasswordHash string
	PersonalInfo PersonalInfo `gorm:"type:jsonb"`
	CreatedAt    time.Time
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
