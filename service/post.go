package service

import (
	"context"
	"sync"
	"to-do-list/pkg/ctl"
	"to-do-list/pkg/util"
	"to-do-list/repository/db/dao"
	"to-do-list/repository/db/model"
	"to-do-list/types"
)

var PostSrvIns *PostSrv
var PostSrvOnce sync.Once

type PostSrv struct {
}

func GetPostSrv() *PostSrv {
	PostSrvOnce.Do(func() {
		PostSrvIns = &PostSrv{}
	})
	return PostSrvIns
}

func (s *PostSrv) CreatePost(ctx context.Context, req *types.CreatePostReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	Post := &model.Post{
		User:    *user,
		UserID:  user.ID,
		Title:   req.Title,
		Content: req.Content,
		Likes:   0,
		Views:   0,
	}
	err = dao.NewPostDao(ctx).CreatePost(Post)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *PostSrv) ListAllPost(ctx context.Context, req *types.ListPostsReq) (resp interface{}, err error) {
	posts, total, err := dao.NewPostDao(ctx).ListAllPost(req.Limit, req.Start)
	postRespList := make([]*types.PostResp, 0)
	for _, v := range posts {
		postRespList = append(postRespList, &types.PostResp{
			ID:      v.ID,
			UserID:  v.UserID,
			Title:   v.Title,
			Content: v.Content,
			Likes:   v.Likes,
			Views:   v.Views,
		})
	}
	return ctl.RespList(postRespList, total), nil
}

// 获取单个用户的所有帖子
func (s *PostSrv) ListPost(ctx context.Context, req *types.ListonePostReq) (resp interface{}, err error) {

	tasks, total, err := dao.NewPostDao(ctx).ListPost(req.Start, req.Limit, req.UserID)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	postRespList := make([]*types.PostResp, 0)
	for _, v := range tasks {
		postRespList = append(postRespList, &types.PostResp{
			ID:      v.ID,
			UserID:  v.UserID,
			Title:   v.Title,
			Content: v.Content,
			Likes:   v.Likes,
			Views:   v.Views,
		})
	}
	return ctl.RespList(postRespList, total), nil
}

// ShowPost 展示某条任务
func (s *PostSrv) ShowPost(ctx context.Context, req *types.ShowPostReq) (resp interface{}, err error) {

	post, err := dao.NewPostDao(ctx).FindPostById(req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	respTask := &types.PostResp{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
		Likes:   post.Likes,
		Views:   int(post.View()),
	}
	post.AddView()
	return ctl.RespSuccessWithData(respTask), nil

}

func (s *PostSrv) DeletePost(ctx context.Context, req *types.DeletePostReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	err = dao.NewPostDao(ctx).DeleteTaskById(u.Id, req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return ctl.RespSuccess(), nil

}

func (s *PostSrv) UpdatePost(ctx context.Context, req *types.UpdatePostReq) (resp interface{}, err error) {

	return
}

func (s *PostSrv) SearchPost(ctx context.Context, req *types.SearchPostReq) (resp interface{}, err error) {
	tasks, err := dao.NewPostDao(ctx).SearchPost(req.Info)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	postRespList := make([]*types.PostResp, 0)
	for _, v := range tasks {
		postRespList = append(postRespList, &types.PostResp{
			ID:      v.ID,
			UserID:  v.UserID,
			Title:   v.Title,
			Content: v.Content,
			Likes:   v.Likes,
			Views:   v.Views,
		})
	}
	return ctl.RespList(postRespList, int64(len(postRespList))), nil
}
