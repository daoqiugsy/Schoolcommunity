package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"to-do-list/pkg/ctl"
	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS

		// 从请求头中获取Token
		token := c.GetHeader("Authorization")

		// 如果没有Token，则返回404错误
		if token == "" {
			code = http.StatusNotFound
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}

		// 解析Token中的claims
		claims, err := util.ParseToken(token)
		log.Print(claims)

		// 如果解析出错，返回错误信息
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			// 如果Token已过期，返回错误信息
			code = e.ErrorAuthCheckTokenTimeout
		}

		// 如果出现错误，返回错误信息并终止请求处理
		if code != e.SUCCESS {
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}

		// 将解析出的用户信息放入请求上下文中，供后续处理使用
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))

		// 调用下一个中间件处理函数继续处理请求
		c.Next()
	}
}
