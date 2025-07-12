package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"particlum_backend/config"

	"gorm.io/gorm"
)

func (p PersonalInfo) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *PersonalInfo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}

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

func FindUserByEmail(email string) (*User, error) {
	var user User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
