package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"to-do-list/pkg/util"
	"to-do-list/service"
	"to-do-list/types"
)

// 创建评论
func CreatCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateCommentReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetCommentSrv()
			resp, err := l.CreatComment(ctx.Request.Context(), &req)
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

// 获取评论
func ListCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListCommentReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			log.Print(req)
			l := service.GetCommentSrv()
			resp, err := l.ListComment(ctx.Request.Context(), &req)
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
func DeleteCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteCommentReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetCommentSrv()
			resp, err := l.DeleteComment(ctx.Request.Context(), &req)
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
