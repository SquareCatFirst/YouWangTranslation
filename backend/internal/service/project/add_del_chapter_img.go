package project

import (
	"errors"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDelChapterImg(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		util.SendJSON(c, -1, "添加权限失败：解析请求失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	if req.Action == "add" {
		imagesSlice, ok := util.AsSlice(req.Data["images"])
		if !ok || len(imagesSlice) == 0 {
			util.SendJSON(c, -1, "添加图片失败：缺少图片数据", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		var resultImages []interface{}

		err := config.DB.Transaction(func(tx *gorm.DB) error {
			for _, item := range imagesSlice {
				imgMap, ok := util.AsMap(item)
				if !ok {
					return errors.New("图片数据格式错误")
				}

				chapterID := util.AsInt64(imgMap["chapter_id"], 0)
				if chapterID == 0 {
					return errors.New("缺少章节ID")
				}

				// 检查章节是否存在
				var count int64
				if tx.Model(&model.ProjectChapter{}).Where("id = ?", chapterID).Count(&count).Error != nil {
					return errors.New("查询章节失败")
				}
				if count == 0 {
					return errors.New("章节ID不存在")
				}

				newImg := model.ChapterWorkImage{
					ChapterID: chapterID,
					Role:      int16(util.AsInt(imgMap["role"], 0)),
					ImageURL:  util.GetString(imgMap, "image_url"),
				}

				if newImg.ImageURL == "" {
					return errors.New("图片URL不能为空")
				}

				if !util.ChapterWorkImageInsert(chapterID, newImg.Role, newImg.ImageURL) {
					return errors.New("插入图片失败")
				}

				if err := tx.Where("chapter_id = ? AND role = ? AND image_url = ?", chapterID, newImg.Role, newImg.ImageURL).
					Order("order_index desc, id desc").
					First(&newImg).Error; err != nil {
					return errors.New("读取新增图片失败")
				}

				resultImages = append(resultImages, gin.H{
					"image_id":    newImg.ID,
					"image_url":   newImg.ImageURL,
					"order_index": newImg.OrderIndex,
				})
			}
			return nil
		})

		if err != nil {
			util.SendJSON(c, -1, "添加图片失败: "+err.Error(), []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		util.SendJSON(c, 0, "添加图片成功", []interface{}{
			gin.H{
				"images": resultImages,
			},
		}, len(imagesSlice), len(imagesSlice), c.FullPath(), c.Request.Method)
		return

	} else if req.Action == "delete" {
		idsSlice, ok := util.AsSlice(req.Data["image_ids"])
		if !ok || len(idsSlice) == 0 {
			util.SendJSON(c, -1, "删除图片失败：缺少图片ID数据", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		err := config.DB.Transaction(func(tx *gorm.DB) error {
			for _, item := range idsSlice {
				id := util.AsInt64(item, 0)
				if id == 0 {
					return errors.New("包含无效的图片ID")
				}

				if err := tx.Delete(&model.ChapterWorkImage{}, id).Error; err != nil {
					return err
				}
			}
			return nil
		})

		if err != nil {
			util.SendJSON(c, -1, "删除图片失败: "+err.Error(), []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}

		util.SendJSON(c, 0, "删除图片成功", []interface{}{}, len(idsSlice), len(idsSlice), c.FullPath(), c.Request.Method)
		return

	} else {
		util.SendJSON(c, -1, "修改图片失败，请求数据不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

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
