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

var CommentSrvIns *CommentSrv
var CommentSrvOnce sync.Once

type CommentSrv struct {
}

func GetCommentSrv() *CommentSrv {
	CommentSrvOnce.Do(func() {
		CommentSrvIns = &CommentSrv{}
	})
	return CommentSrvIns
}

// 创建评论

func (s *CommentSrv) CreatComment(ctx context.Context, req *types.CreateCommentReq) (resp interface{}, err error) {
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
	Comment := &model.Comment{
		UserId:  user.ID,
		PostId:  req.PostId,
		Content: req.Content,
	}
	err = dao.NewCommentDao(ctx).CreateComment(Comment)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

// 删除评论
func (s *CommentSrv) DeleteComment(ctx context.Context, req *types.DeleteCommentReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	err = dao.NewCommentDao(ctx).DeleteCommentById(u.Id, req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return ctl.RespSuccess(), nil

}

func (s *CommentSrv) ListComment(ctx context.Context, req *types.ListCommentReq) (resp interface{}, err error) {
	comments, total, err := dao.NewCommentDao(ctx).ListComment(req.PostId)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	commentRespList := make([]*types.CommentResp, 0)
	for _, v := range comments {
		commentRespList = append(commentRespList, &types.CommentResp{
			Id:      v.ID,
			UserId:  v.UserId,
			PostId:  v.PostId,
			Content: v.Content,
		})
	}

	return ctl.RespList(commentRespList, total), nil
}
