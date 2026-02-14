package project

import (
	"errors"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddChangeChapterAss(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		return
	}

	if req.Action == "add" {
		projID := util.GetInt64(req.Data, "project_id")
		if projID == 0 {
			util.SendJSON(c, -1, "添加权限失败：缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		if err := config.DB.Where("id = ?", projID).First(&model.Project{}).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				util.SendJSON(c, -1, "添加权限失败：项目不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
				return
			}
			util.SendJSON(c, -1, "添加权限失败：查询项目失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		chapterID := util.GetInt64(req.Data, "chapter_id")
		if chapterID == 0 {
			var firstChapter model.ProjectChapter
			if err := config.DB.Where("project_id = ?", projID).Order("order_index asc, id asc").First(&firstChapter).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					util.SendJSON(c, -1, "添加权限失败：项目下无章节可分配", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
					return
				}
				util.SendJSON(c, -1, "添加权限失败：查询章节失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
				return
			}
			chapterID = firstChapter.ID
		}

		if err := config.DB.Where("id = ?", chapterID).First(&model.ProjectChapter{}).Error; err != nil {
			util.SendJSON(c, -1, "添加权限失败：章节不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		ctxUID, hasUID := c.Get("user_id")
		loginUID := int64(0)
		if hasUID {
			loginUID = util.AsInt64(ctxUID, 0)
		}

		userID := util.GetInt64(req.Data, "user_id")
		if userID == 0 {
			userID = loginUID
		}
		if userID == 0 {
			util.SendJSON(c, -1, "添加权限失败：缺少用户ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		role := util.GetInt(req.Data, "role")
		if role < 0 || role > 5 {
			role = 0
		}

		assignedBy := util.GetInt64(req.Data, "assigned_by")
		if assignedBy == 0 {
			assignedBy = loginUID
		}
		if assignedBy == 0 {
			assignedBy = userID
		}

		if err := config.DB.Where("id = ?", userID).First(&model.User{}).Error; err != nil {
			util.SendJSON(c, -1, "添加权限失败：用户不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		if config.DB.Transaction(func(tx *gorm.DB) error {
			if !util.ChapterAssignmentInsert(tx, chapterID, userID, role, assignedBy) {
				return errors.New("写入数据库失败")
			}
			return nil
		}) != nil {
			util.SendJSON(c, -1, "添加权限失败：写入数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		util.SendJSON(c, 0, "添加权限成功", []interface{}{
			gin.H{
				"chapter_id":  chapterID,
				"user_id":     userID,
				"role":        role,
				"assigned_by": assignedBy,
			},
		}, 1, 1, c.FullPath(), c.Request.Method)
		return

	} else if req.Action == "change" {
		if req.Data["assignment_id"] == nil {
			util.SendJSON(c, -1, "修改权限失败：缺少权限记录ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		assignmentID := util.GetInt64(req.Data, "assignment_id")
		row := model.ChapterAssignment{}
		if config.DB.Where("id = ?", assignmentID).First(&row).Error != nil {
			util.SendJSON(c, -1, "修改权限失败：权限记录不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		updates := map[string]interface{}{}

		if req.Data["chapter_id"] != nil {
			chapterID := util.GetInt64(req.Data, "chapter_id")
			if config.DB.Where("id = ?", chapterID).First(&model.ProjectChapter{}).Error != nil {
				util.SendJSON(c, -1, "修改权限失败：章节不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
				return
			}
			updates["chapter_id"] = chapterID
		}

		if req.Data["user_id"] != nil {
			updates["user_id"] = util.GetInt64(req.Data, "user_id")
		}

		if req.Data["role"] != nil {
			updates["role"] = util.GetInt(req.Data, "role")
		}

		if req.Data["assigned_by"] != nil {
			assignedBy := util.GetInt64(req.Data, "assigned_by")
			updates["assigned_by"] = assignedBy
		}

		if len(updates) == 0 {
			util.SendJSON(c, -1, "修改权限失败：没有可更新字段", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		if config.DB.Model(&row).Updates(updates).Error != nil {
			util.SendJSON(c, -1, "修改权限失败：写入数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		if config.DB.Where("id = ?", assignmentID).First(&row).Error != nil {
			util.SendJSON(c, -1, "修改权限失败：读取更新结果失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		util.SendJSON(c, 0, "修改权限成功", []interface{}{
			gin.H{
				"assignment_id": row.ID,
				"chapter_id":    row.ChapterID,
				"user_id":       row.UserID,
				"role":          row.Role,
				"assigned_by":   row.AssignedBy,
			},
		}, 1, 1, c.FullPath(), c.Request.Method)
		return

	} else {
		util.SendJSON(c, -1, "操作失败: 操作类型不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	// if req.Data["action"] == nil {
	// 	util.SendJSON(c, -1, "操作失败: 缺少操作类型", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// } else if req.Data["action"] == "POST" {
	// 	if req.Data["project_id"] == nil {
	// 		util.SendJSON(c, -1, "添加权限失败: 缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 		return
	// 	}

	// 	projID := util.GetInt64(req.Data, "project_id")

	// } else if req.Data["action"] == "change" {

	// } else {
	// 	util.SendJSON(c, -1, "操作失败: 操作类型不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// var projID int64
	// if req.Data["project_id"] == "" {
	// 	util.SendJSON(c, -1, "添加权限失败: 缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
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
