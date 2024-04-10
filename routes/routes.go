package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"to-do-list/api"
	"to-do-list/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() // 生成了一个WSGI应用程序实例
	store := cookie.NewStore([]byte("something-very-secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())
		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 帖子操作
			authed.POST("post_create", api.CreatePostHandler())   // 创建帖子
			authed.GET("post_all_list", api.ListAllPostHandler()) // 获取所有帖子列表
			authed.GET("post_list", api.ListPostHandler())        //获取单个用户所有帖子
			authed.GET("post_show", api.ShowPostHandler())        // 获取帖子详情
			authed.GET("post_search", api.SearchPostHandler())    // 搜索帖子
			authed.DELETE("post_delete", api.DeletePostHandler()) // 删除帖子
			//评论操作
			authed.POST("comment_create", api.CreatCommentHandler())    //创建评论
			authed.GET("comment_list", api.ListCommentHandler())        //获取评论列表
			authed.DELETE("comment_delete", api.DeleteCommentHandler()) //删除评论
			//私聊操作
			authed.POST("message_creat", api.CreatMesgHandler()) //发送消息
			authed.GET("message_list", api.ListMesgHandler())
		}
	}
	return r
}
