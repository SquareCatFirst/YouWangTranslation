package router

import (
	service "github.com/SquareCatFirst/YouWangTranslation/backend/internal/service/login"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/auth")
	{
		api.POST("/login", service.Login)
	}
}
