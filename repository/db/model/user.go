package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"type:varchar(20);not null" json:"user_name"`
	PasswordDigest string `gorm:"type:varchar(500);not null" json:"password"`
	Role           int    `gorm:"type:int;DEFAULT:0" json:"role" ` //0为用户，1为管理员
	Phone          string `gorm:"type:varchar(20);not null" json:"phone"`
}

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
