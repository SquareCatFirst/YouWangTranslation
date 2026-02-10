package project

import (
	"github.com/gin-gonic/gin"
)

func AddChangeChapter(c *gin.Context) {
	// req, ok := util.ParseRequest(c)
	// if !ok {
	// 	util.SendJSON(c, -1, "操作章节失败：解析请求失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// if req.Data["action"] == nil {
	// 	util.SendJSON(c, -1, "操作章节失败: 缺少操作类型", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// } else if req.Data["action"] == "add" {
	// 	if req.Data["project_id"] == nil {
	// 		util.SendJSON(c, -1, "添加章节失败: 缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 		return
	// 	}
	// 	projID := util.GetInt64(req.Data, "project_id")
	// 	if config.DB.Where("id = ?", projID).First(&model.Project{}).Error != nil {
	// 		if config.DB.Where("id = ?", projID).First(&model.Project{}).Error == gorm.ErrRecordNotFound {
	// 			util.SendJSON(c, -1, "添加章节失败: 项目不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 			return
	// 		}
	// 	}

	// 	row := model.ProjectChapter{
	// 		ProjectID:   projID,
	// 		ChapterName: "未命名章节" + util.GenSN(),
	// 	}

	// 	if config.DB.Transaction(func(tx *gorm.DB) error {
	// 		if tx.Create(&row).Error != nil {
	// 			return tx.Error
	// 		}
	// 		return nil
	// 	}) == nil {
	// 		util.SendJSON(c, 0, "添加章节成功", []interface{}{
	// 			"chapter_id": row.ID,
	//   			"chapter_name": row.ChapterName,
	//   			"is_visible": row.IsVisible,
	//   			"order_index": row.OrderIndex,
	//   			"project_chapters": []
	// 		}, 1, 1, c.FullPath(), c.Request.Method)
	// 	} else {
	// 		util.SendJSON(c, -1, "添加章节失败：添加数据失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 		return
	// 	}
	// } else if req.Data["action"] == "change" {

	// } else {
	// 	util.SendJSON(c, -1, "操作章节失败: 操作类型不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// var projID int64
	// if req.Data["project_id"] == "" {
	// 	util.SendJSON(c, -1, "操作章节失败: 缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// } else {
	// 	projID = util.GetInt64(req.Data, "project_id")
	// }

	// var role int
	// if req.Data["role"] == "" {
	// 	util.SendJSON(c, -1, "添加权限失败: 缺少添加角色", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// } else if req.Data["role"] == "admin" {
	// 	role = 0
	// } else if req.Data["role"] == "source" {
	// 	role = 1
	// } else if req.Data["role"] == "supervisor" {
	// 	role = 5
	// } else if req.Data["role"] == "translator" {
	// 	role = 2
	// } else if req.Data["role"] == "proofreader" {
	// 	role = 3
	// } else if req.Data["role"] == "typesetter" {
	// 	role = 4
	// } else {
	// 	util.SendJSON(c, -1, "添加权限失败: 角色不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// // tags 是 JSON 里的数组时通常会被解码成 []interface{}。
	// // 这里正确判断：没传 / null / 空数组 时设置默认值。
	// if v, ok := req.Data["tags"]; !ok || v == nil {
	// 	req.Data["tags"] = []int{4}
	// } else if s, ok := v.([]interface{}); ok && len(s) == 0 {
	// 	req.Data["tags"] = []int{4}
	// }

	// row := model.Project{
	// 	CoverURL: util.GetPtrString(req.Data, "cover_url"),

	// 	Name:           util.GetString(req.Data, "name"),
	// 	TranslatedName: util.GetPtrString(req.Data, "translated_name"),
	// 	Author:         util.GetPtrString(req.Data, "author"),
	// 	SourceSite:     util.GetPtrString(req.Data, "source_site"),

	// 	Description:            util.GetPtrString(req.Data, "description"),
	// 	TranslationDescription: util.GetPtrString(req.Data, "translation_description"),
	// }

	// if config.DB.Where("id = ?", projID).First(&model.Project{}).Error != nil {
	// 	if config.DB.Where("id = ?", projID).First(&model.Project{}).Error == gorm.ErrRecordNotFound {
	// 		util.SendJSON(c, -1, "添加权限失败: 项目不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 		return
	// 	}
	// 	util.SendJSON(c, -1, "添加权限失败: 查询数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// var role_num int64
	// if config.DB.Model(&model.ProjectAssignment{}).Where("project_id = ? AND role = ?", projID, role).Count(&role_num).Error != nil {
	// 	util.SendJSON(c, -1, "添加权限失败: 查询数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// if config.DB.Transaction(func(tx *gorm.DB) error {
	// 	if req.Data["role"] == "admin" {
	// 		row.AdminMax = max(row.AdminMax, int(role_num)+len(req.Data["admin"].([]interface{})))
	// 	} else if req.Data["role"] == "source" {
	// 		row.SourceMax = max(row.AdminMax, int(role_num)+len(req.Data["admin"].([]interface{})))
	// 	} else if req.Data["role"] == "supervisor" {
	// 		row.SupervisorMax = max(row.AdminMax, int(role_num)+len(req.Data["admin"].([]interface{})))
	// 	} else if req.Data["role"] == "translator" {
	// 		row.TranslatorMax = max(row.AdminMax, int(role_num)+len(req.Data["admin"].([]interface{})))
	// 	} else if req.Data["role"] == "proofreader" {
	// 		row.ProofreaderMax = max(row.AdminMax, int(role_num)+len(req.Data["admin"].([]interface{})))
	// 	} else if req.Data["role"] == "typesetter" {
	// 		row.TypesetterMax = max(row.AdminMax, int(role_num)+len(req.Data["admin"].([]interface{})))
	// 	}

	// 	if tx.Where("id = ?", projID).Updates(&row).Error != nil {
	// 		return tx.Error
	// 	}

	// 	for _, item := range req.Data["admin"].([]interface{}) {
	// 		m, ok := item.(map[string]interface{})
	// 		if !ok {
	// 			return errors.New("添加权限失败: admin 数组元素格式不合法")
	// 		}
	// 		if m["user_id"] == nil || m["assigned_by"] == nil {
	// 			return errors.New("添加权限失败: 请求数据不合法")
	// 		}
	// 		assignedBy := util.AsInt64(m["assigned_by"], 0)
	// 		ass_row := model.ProjectAssignment{
	// 			ProjectID:  projID,
	// 			UserID:     m["user_id"].(int64),
	// 			Role:       role,
	// 			AssignedBy: &assignedBy,
	// 		}

	// 		if tx.Create(&ass_row).Error != nil {
	// 			return tx.Error
	// 		}
	// 	}
	// 	return nil
	// }) == nil {
	// 	util.SendJSON(c, 0, "添加权限成功", []gin.H{
	// 		{
	// 			"project_id": row.ID,
	// 			"name":       req.Data["name"],
	// 			"created_at": row.CreatedAt,
	// 		}}, 1, 1, c.FullPath(), c.Request.Method)
	// } else {
	// 	util.SendJSON(c, -1, "添加权限失败：添加数据失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }
}
