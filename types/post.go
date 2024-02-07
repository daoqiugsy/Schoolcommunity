package types

type ShowPostReq struct {
	Id uint `json:"id" form:"id"`
}

type DeletePostReq struct {
	Id uint `json:"id" form:"id"`
}

type UpdatePostReq struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
}

type CreatePostReq struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
}

type SearchPostReq struct {
	Info string `form:"info" json:"info"`
}

// 所有用户所有帖子
type ListPostsReq struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

// 单个用户单个帖子
type ListonePostReq struct {
	UserID uint `form:"user_id" json:"user_id"`
	Limit  int  `form:"limit" json:"limit"`
	Start  int  `form:"start" json:"start"`
}

// swagger:response Resp
type PostResp struct {
	ID      uint   `json:"id" example:"1"` // 任务ID
	UserID  uint   `json:"user_id" example:"daoqiu"`
	Title   string `json:"title" example:"吃饭"`   // 题目
	Content string `json:"content" example:"睡觉"` // 内容
	Likes   int    `json:"likes" example:"12"`
	Views   int    `json:"views" example:"32"` // 浏览量
}
