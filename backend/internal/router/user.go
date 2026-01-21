package router

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/service/login"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/auth")
	{
		api.POST("/login", login.Login)
		api.GET("/user/list", user.UserList)
	}
}
