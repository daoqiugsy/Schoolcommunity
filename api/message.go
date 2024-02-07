package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"to-do-list/pkg/util"
	"to-do-list/service"
	"to-do-list/types"
)

// 创建私聊消息
func CreatMesgHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateMessageReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetMessageSrv()
			resp, err := l.CreatMessage(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}

}

// 获取私聊消息
func ListMesgHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListMessageReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			log.Print(req)
			l := service.GetMessageSrv()
			resp, err := l.ListMessage(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

// 删除评论
