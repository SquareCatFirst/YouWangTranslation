package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		uid := sess.Get("user_id")
		if uid == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "未登录",
			})
			return
		}
		// 放到上下文，后续 handler 直接用
		c.Set("user_id", uid)
		c.Next()
	}
}
