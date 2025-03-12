package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string //存储的是密文，即加密后的密码
}
