package types

// 创建评论
type CreateCommentReq struct {
	PostId  uint   `json:"post_id" form:"post_id"`
	Content string `json:"content" form:"content"`
}

// 获取评论
type ListCommentReq struct {
	PostId uint `json:"post_id" form:"post_id"`
}
type CommentResp struct {
	Id      uint   `json:"id"`
	UserId  uint   `json:"user_id"`
	PostId  uint   `json:"post_id"`
	Content string `json:"content"`
}
type DeleteCommentReq struct {
	Id uint `json:"id" form:"id"`
}
