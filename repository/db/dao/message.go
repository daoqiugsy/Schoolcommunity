package dao

import (
	"context"
	"gorm.io/gorm"
	"to-do-list/repository/db/model"
)

type MessageDao struct {
	*gorm.DB
}

func (d *MessageDao) CreateMessage(Message *model.Message) error {
	return d.Model(&model.Message{}).Create(Message).Error
}

func (d *MessageDao) ListMessage(sendid uint, reciverid uint) (r []*model.Message, total int64, err error) {
	err = d.Model(&model.Message{}).Where("sender_id=? and receiver_id=?", sendid, reciverid).Or("sender_id=? and receiver_id=?", reciverid, sendid).Order("created_at").Count(&total).Find(&r).Error
	//按时间顺序
	return
}

func NewMessageDao(ctx context.Context) *MessageDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &MessageDao{NewDBClient(ctx)}
}
