package model

import (
	"encoding/json"
	"particlum_backend/config"
)

// 插入用户到数据库
func CreateUser(user *User) error {
	personalInfoJSON, err := json.Marshal(user.PersonalInfo)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (username, email, password_hash, personal_info)
		VALUES ($1, $2, $3, $4)
	`

	result := config.DB.Exec(query, user.Username, user.Email, user.PasswordHash, personalInfoJSON)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
