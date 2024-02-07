package service

import (
	"context"
	"log"
	"sync"
	"to-do-list/pkg/ctl"
	"to-do-list/pkg/util"
	"to-do-list/repository/db/dao"
	"to-do-list/repository/db/model"
	"to-do-list/types"
)

var MessageSrvIns *MessageSrv
var MessageSrvOnce sync.Once

type MessageSrv struct {
}

func GetMessageSrv() *MessageSrv {
	MessageSrvOnce.Do(func() {
		MessageSrvIns = &MessageSrv{}
	})
	return MessageSrvIns
}

// 创建评论

func (s *MessageSrv) CreatMessage(ctx context.Context, req *types.CreateMessageReq) (resp interface{}, err error) {
	log.Print(req)
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
	Message := &model.Message{
		SenderId:   user.ID,
		ReceiverId: req.ReceiverId,
		Content:    req.Content,
	}
	err = dao.NewMessageDao(ctx).CreateMessage(Message)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

//

func (s *MessageSrv) ListMessage(ctx context.Context, req *types.ListMessageReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	Messages, total, err := dao.NewMessageDao(ctx).ListMessage(u.Id, req.ReceiverId)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	MessageRespList := make([]*types.MessageResp, 0)
	for _, v := range Messages {
		MessageRespList = append(MessageRespList, &types.MessageResp{
			Id:         v.ID,
			SenderId:   v.SenderId,
			ReceiverId: v.ReceiverId,
			Content:    v.Content,
			CreatedAt:  v.CreatedAt,
		})
	}

	return ctl.RespList(MessageRespList, total), nil
}
