package dao

import (
	"context"
	"gorm.io/gorm"
	"to-do-list/repository/db/model"
)

type PostDao struct {
	*gorm.DB
}

func (s *PostDao) CreatePost(post *model.Post) error {
	return s.Model(&model.Post{}).Create(&post).Error
}

func NewPostDao(ctx context.Context) *PostDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &PostDao{NewDBClient(ctx)}
}
func (s *PostDao) ListAllPost(start, limit int) (r []*model.Post, total int64, err error) {
	err = s.Model(&model.Post{}).Preload("User").
		Where("is_deleted=0").
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return
}
func (s *PostDao) ListPost(start, limit int, userId uint) (r []*model.Post, total int64, err error) {
	err = s.Model(&model.Post{}).Preload("User").Where("user_id = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return
}

func (s *PostDao) FindPostById(id uint) (r *model.Post, err error) {
	err = s.Model(&model.Post{}).Where("id = ?", id).First(&r).Error

	return
}

func (s *PostDao) SearchPost(info string) (r []*model.Post, err error) {
	err = s.Model(&model.Post{}).Where("title LIKE ? OR content LIKE ?",
		"%"+info+"%", "%"+info+"%").Find(&r).Error
	return

}

// PostDao 通过id删除
func (s *PostDao) DeleteTaskById(uId, tId uint) (err error) {
	var r *model.Post
	err = s.Model(&model.Post{}).Where("id=? AND user_id = ?", tId, uId).First(&r).Error
	err = s.Model(&model.Post{}).Delete(&r).Error
	return
}

//func (s *PostDao) FindPostByIdAndUserId(uId uint, tId uint) (r *model.Post, err error) {
//	err = s.Model(&model.Post{}).Where("id = ? AND user_id = ? and is_deleted=0", tId, uId).First(&r).Error
//
//	return
//}
