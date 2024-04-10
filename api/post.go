package api

import (
	"net/http"

	"to-do-list/consts"
	"to-do-list/pkg/util"
	"to-do-list/service"
	"to-do-list/types"

	"github.com/gin-gonic/gin"
)

// 创建任务
func CreatePostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreatePostReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetPostSrv()                              // 获取任务服务
			resp, err := l.CreatePost(ctx.Request.Context(), &req) // 调用任务服务
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

// 获取全部任务列表
func ListAllPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListPostsReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.Limit == 0 {
				req.Limit = consts.BasePageLimit
			}
			l := service.GetPostSrv()
			resp, err := l.ListAllPost(ctx.Request.Context(), &req)
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

func ListPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListonePostReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.Limit == 0 {
				req.Limit = consts.BasePageLimit
			}
			l := service.GetPostSrv()
			resp, err := l.ListPost(ctx.Request.Context(), &req)
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

func ShowPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ShowPostReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetPostSrv()
			resp, err := l.ShowPost(ctx.Request.Context(), &req)
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

func DeletePostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeletePostReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetPostSrv()
			resp, err := l.DeletePost(ctx.Request.Context(), &req)
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

// 查询任务
func SearchPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchPostReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetPostSrv()
			resp, err := l.SearchPost(ctx.Request.Context(), &req)
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
