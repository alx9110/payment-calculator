package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	BotId    string `json:"botid" gorm:"unique"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	BotId    int16  `json:"bot_id"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	ID       uint    `json:"-"`
	Email    float32 `json:"hot_value"`
	Password string  `json:"password"`
}
