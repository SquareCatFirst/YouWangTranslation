package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 可以挂载全局中间件
	//	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 注册模块路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	RegisterUserRoutes(r)

	return r
}
