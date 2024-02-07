package types

import (
	"time"
)

type CreateMessageReq struct {
	ReceiverId uint   `json:"receiver_id" form:"receiver_id"`
	Content    string `json:"content" form:"content"`
}
type ListMessageReq struct {
	ReceiverId uint `json:"receiver_id" form:"receiver_id"`
}
type MessageResp struct {
	Id         uint      `json:"id"`
	SenderId   uint      `json:"sender_id"`
	ReceiverId uint      `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
