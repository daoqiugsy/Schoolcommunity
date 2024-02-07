package model

import "gorm.io/gorm"

type Message struct {
	User User `gorm:"foreignKey:SenderId"`
	gorm.Model
	SenderId   uint
	ReceiverId uint
	Content    string
	IsRead     bool
}
