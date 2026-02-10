package project

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DelProj(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		util.SendJSON(c, -1, "删除项目失败：解析请求失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	if v, ok := req.Data["project_id"]; !ok || v == nil {
		util.SendJSON(c, -1, "删除项目失败：缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	type mem struct {
		user_id int64
		role    int32
	}

	if config.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Delete(&model.Project{}, util.GetInt64(req.Data, "project_id")).Error != nil {
			return tx.Error
		}
		return nil
	}) == nil {
		util.SendJSON(c, 0, "删除项目成功", []gin.H{
			{}}, 0, 0, c.FullPath(), c.Request.Method)
	} else {
		util.SendJSON(c, -1, "删除项目失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}
}
