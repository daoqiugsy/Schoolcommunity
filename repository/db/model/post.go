package model

import "gorm.io/gorm"

type Post struct {
	User User `gorm:"foreignKey:UserID"`
	gorm.Model
	UserID  uint
	Title   string
	Content string
	Sid     string `gorm:"DEFAULT:0"`
	Likes   int    `gorm:"DEFAULT:0"`
	Views   int    `gorm:"DEFAULT:0"`
}
