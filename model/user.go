package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string //存储的是密文，即加密后的密码
}

// 加密，注册时做的操作
func (user *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

// 验证密码，登录时做的操作
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
