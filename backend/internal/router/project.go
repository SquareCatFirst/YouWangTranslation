package router

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/middleware"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/service/project"
	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(r *gin.Engine) {
	proj := r.Group("/api/v1/project")
	proj.Use(middleware.RequireLogin())
	{
		proj.POST("change", project.ChangeProj)
		proj.POST("", project.CreateProj)
		proj.GET("", project.GetProj)
		proj.DELETE("", project.DelProj)
		proj.POST("/chapter", project.AddChangeChapter)
		proj.DELETE("/chapter", project.DelChapter)
		proj.POST("/chapter/image", project.AddDelChapterImg)
		proj.POST("/chapter_assignments", project.AddChangeChapterAss)
		proj.DELETE("/chapter_assignments", project.DelChapterAss)
		proj.POST("/project_assignments", project.AddProjAss)
	}
}
