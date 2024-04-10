package model

import (
	"gorm.io/gorm"
	"strconv"
	"to-do-list/repository/cache"
)

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

// 获取帖子浏览量
func (Post *Post) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PostViewKey(Post.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}
func (Post *Post) AddView() {
	cache.RedisClient.Incr(cache.PostViewKey(Post.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Post.ID)))
}
