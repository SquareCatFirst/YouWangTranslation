package project

import (
	"time"

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

	if req.Data["tags"] == nil {
		util.SendJSON(c, -1, "新建项目失败，标签不能为空", []interface{}{}, 0, c.FullPath(), c.Request.Method)
		return
	}

	if req.Data["project_chapters"] == nil {
		util.SendJSON(c, -1, "项目章节不能为空", []interface{}{}, 0, c.FullPath(), c.Request.Method)
		return
	}

	if req.Data["translated_name"] == nil {
		req.Data["translated_name"] = "未命名作品_" + time.Now().Format("20060102150405")
	}

	if req.Data["name"] == nil {
		req.Data["name"] = "未命名项目_" + time.Now().Format("20060102150405")
	}

	if req.Data["cover_url"] == nil {
		req.Data["cover_url"] = "https://photogzmaz.photo.store.qq.com/psc?/V51HTR1T39Mw3r0KnhDf0s3GGb00Md4m/LiySpxowE0yeWXwBdXN*SZtR4lEJbZQy.tXoFnEAwuCc.zRtxEvY7TUHn1XeFAvH2Ap9TR6dZLWify3q1DlUXO1lm5iuv4eNwoFneoXOodU!/b&bo=qgHwAKoB8AACByM!&rf=viewer_4"
	}

	row := model.Project{
		CoverURL: util.GetPtrString(req.Data, "cover_url"),

		Name:           util.GetString(req.Data, "name"),
		TranslatedName: util.GetPtrString(req.Data, "translated_name"),
		Author:         util.GetPtrString(req.Data, "author"),
		SourceSite:     util.GetPtrString(req.Data, "source_site"),

		Description:            util.GetPtrString(req.Data, "description"),
		TranslationDescription: util.GetPtrString(req.Data, "translation_description"),

		SourceMax:      util.GetInt(req.Data, "source_max"),
		TranslatorMax:  util.GetInt(req.Data, "translator_max"),
		ProofreaderMax: util.GetInt(req.Data, "proofreader_max"),
		TypesetterMax:  util.GetInt(req.Data, "typesetter_max"),
		SupervisorMax:  util.GetInt(req.Data, "supervisor_max"),
		AdminMax:       util.GetInt(req.Data, "admin_max"),
	}

	type mem struct {
		user_id int64
		role    int32
	}

	// var mpusr map[int64]map[]

	if config.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Create(&row).Error != nil {
			return tx.Error
		}

		for _, item := range req.Data["project_chapters"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_chapter_row := model.ProjectChapter{
				ProjectID:   row.ID,
				ChapterName: util.GetString(mp, "chapter_name"),
				IsVisible:   util.GetBool(mp, "is_visible"),
				OrderIndex:  util.GetInt(mp, "order_index"),
			}

			if tx.Create(&proj_chapter_row).Error != nil {
				return tx.Error
			}

			for _, chapter_item := range mp["images"].([]interface{}) {
				chapter_mp := chapter_item.(map[string]interface{})
				chapter_work_img_row := model.ChapterWorkImage{
					ChapterID: proj_chapter_row.ID,
					Role:      0,
					ImageURL:  chapter_mp["image_url"].(string),

					OrderIndex: util.GetInt(chapter_mp, "order_index"),
				}

				if tx.Create(&chapter_work_img_row).Error != nil {
					return tx.Error
				}
			}
		}

		for _, item := range req.Data["admin"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_ass_row := model.ProjectAssignment{
				ProjectID:  row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       0,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&proj_ass_row).Error != nil {
				return tx.Error
			}

			chapter_ass_row := model.ChapterAssignment{
				ChapterID:  proj_ass_row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       0,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&chapter_ass_row).Error != nil {
				return tx.Error
			}
		}

		for _, item := range req.Data["source"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_ass_row := model.ProjectAssignment{
				ProjectID:  row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       1,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&proj_ass_row).Error != nil {
				return tx.Error
			}

			chapter_ass_row := model.ChapterAssignment{
				ChapterID:  proj_ass_row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       1,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&chapter_ass_row).Error != nil {
				return tx.Error
			}
		}

		for _, item := range req.Data["translator"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_ass_row := model.ProjectAssignment{
				ProjectID:  row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       2,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&proj_ass_row).Error != nil {
				return tx.Error
			}

			chapter_ass_row := model.ChapterAssignment{
				ChapterID:  proj_ass_row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       2,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&chapter_ass_row).Error != nil {
				return tx.Error
			}
		}

		for _, item := range req.Data["proofreader"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_ass_row := model.ProjectAssignment{
				ProjectID:  row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       3,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&proj_ass_row).Error != nil {
				return tx.Error
			}

			chapter_ass_row := model.ChapterAssignment{
				ChapterID:  proj_ass_row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       3,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&chapter_ass_row).Error != nil {
				return tx.Error
			}
		}

		for _, item := range req.Data["typesetter"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_ass_row := model.ProjectAssignment{
				ProjectID:  row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       4,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&proj_ass_row).Error != nil {
				return tx.Error
			}

			chapter_ass_row := model.ChapterAssignment{
				ChapterID:  proj_ass_row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       4,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&chapter_ass_row).Error != nil {
				return tx.Error
			}
		}

		for _, item := range req.Data["supervisor"].([]interface{}) {
			mp := item.(map[string]interface{})
			proj_ass_row := model.ProjectAssignment{
				ProjectID:  row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       5,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&proj_ass_row).Error != nil {
				return tx.Error
			}

			chapter_ass_row := model.ChapterAssignment{
				ChapterID:  proj_ass_row.ID,
				UserID:     util.GetInt64(mp, "user_id"),
				Role:       5,
				AssignedBy: util.GetPtrInt64(mp, "assigned_by"),
			}

			if tx.Create(&chapter_ass_row).Error != nil {
				return tx.Error
			}
		}

		for _, item := range req.Data["tags"].([]interface{}) {
			proj_tag_row := model.ProjectTag{
				ProjectID: row.ID,
				TagID:     util.GetInt64(item.(map[string]interface{}), "tag_id"),
			}

			if tx.Create(&proj_tag_row).Error != nil {
				return tx.Error
			}
		}

		return nil
	}) == nil {
		chapterCount := 0
		if v := req.Data["project_chapters"]; v != nil {
			if chapters, ok := v.([]interface{}); ok {
				chapterCount = len(chapters)
			}
		}

		util.SendJSON(c, 0, "新建项目成功", []gin.H{
			{
				"project_id":    row.ID,
				"name":          req.Data["name"],
				"cover_url":     req.Data["cover_url"],
				"created_at":    row.CreatedAt,
				"chapter_count": chapterCount,
			}}, 1, c.FullPath(), c.Request.Method)
	} else {
		util.SendJSON(c, -1, "新建项目失败", []interface{}{}, 0, c.FullPath(), c.Request.Method)
		return
	}
}
