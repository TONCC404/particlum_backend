package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"particlum_backend/config"
	"particlum_backend/model"
	"time"
)

// 模拟登录用户 ID
var currentUserID uint = 1

// 用户关注论坛
func FollowForum(c *gin.Context) {
	var req struct {
		ForumID uint `json:"forum_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	config.DB.Create(&model.UserFollowForum{
		UserID:     currentUserID,
		ForumID:    req.ForumID,
		FollowedAt: time.Now(),
	})
	c.JSON(http.StatusOK, gin.H{"message": "followed"})
}

// 获取用户关注的论坛
func GetFollowedForums(c *gin.Context) {
	var forums []model.Forum
	config.DB.Table("forums").
		Joins("JOIN user_follow_forums ON forums.id = user_follow_forums.forum_id").
		Where("user_follow_forums.user_id = ?", currentUserID).
		Find(&forums)

	c.JSON(http.StatusOK, forums)
}

// 获取推荐帖子（示例：7日内点赞最多）
func GetHotPosts(c *gin.Context) {
	var results []struct {
		ID        uint
		Title     string
		Content   string
		LikeCount int
		ForumName string
	}
	config.DB.Raw(`
		SELECT p.id, p.title, p.content, f.name AS forum_name,
		       (SELECT COUNT(*) FROM post_likes pl WHERE pl.post_id = p.id) AS like_count
		FROM posts p
		JOIN forums f ON f.id = p.forum_id
		WHERE p.created_at >= ?
		ORDER BY like_count DESC
		LIMIT 20
	`, time.Now().AddDate(0, 0, -7)).Scan(&results)

	c.JSON(http.StatusOK, results)
}
