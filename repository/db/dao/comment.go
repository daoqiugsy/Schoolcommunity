package dao

import (
	"context"
	"gorm.io/gorm"
	"to-do-list/repository/db/model"
)

type CommentDao struct {
	*gorm.DB
}

func (d *CommentDao) CreateComment(comment *model.Comment) error {
	return d.Model(&model.Comment{}).Create(comment).Error
}

func (d *CommentDao) ListComment(id uint) (r []*model.Comment, total int64, err error) {
	err = d.Model(&model.Comment{}).Where("post_id = ?", id).Count(&total).Find(&r).Error
	return
}

func (d *CommentDao) DeleteCommentById(uId, tId uint) (err error) {
	var r *model.Comment
	err = d.Model(&model.Comment{}).Where("id=? AND user_id = ?", tId, uId).First(&r).Delete(&r).Error
	return

}
func NewCommentDao(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{NewDBClient(ctx)}
}
