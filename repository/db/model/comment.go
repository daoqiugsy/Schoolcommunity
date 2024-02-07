package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string //评论内容
	PostId  uint   //文章id
	UserId  uint   //用户id
}
