package router

import (
	// "github.com/SquareCatFirst/YouWangTranslation/backend/internal/middleware"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/middleware"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/service/login"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", login.Login)
		auth.DELETE("/logout", login.Logout)
	}

	apiUser := r.Group("/api/v1/user")
	{
		apiUser.GET("/list", user.UserList)
	}

	apiUserAccount := r.Group("/api/v1/user/account")
	apiUserAccount.Use(middleware.RequireLogin())
	{
		apiUserAccount.POST("", user.AccountChange)
	}
}
