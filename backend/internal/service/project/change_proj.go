package project

import (
	"time"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangeProj(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		util.SendJSON(c, -1, "修改项目失败：解析请求失败", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}

	if req.Data["source_site"] == "" ||
		req.Data["description"] == "" ||
		req.Data["translation_description"] == "" {
		util.SendJSON(c, -1, "修改项目失败: 缺少来源站点、项目简介或翻译描述", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}

	if req.Data["translated_name"] == "" {
		req.Data["translated_name"] = "未命名作品_" + time.Now().Format("20060102150405")
	}

	if req.Data["name"] == "" {
		req.Data["name"] = "未命名项目_" + time.Now().Format("20060102150405")
	}

	if req.Data["cover_url"] == "" {
		req.Data["cover_url"] = "https://photogzmaz.photo.store.qq.com/psc?/V51HTR1T39Mw3r0KnhDf0s3GGb00Md4m/LiySpxowE0yeWXwBdXN*SZtR4lEJbZQy.tXoFnEAwuCc.zRtxEvY7TUHn1XeFAvH2Ap9TR6dZLWify3q1DlUXO1lm5iuv4eNwoFneoXOodU!/b&bo=qgHwAKoB8AACByM!&rf=viewer_4"
	}

	if req.Data["author"] == "" {
		req.Data["author"] = "未知作者"
	}

	// tags 是 JSON 里的数组时通常会被解码成 []interface{}。
	// 这里正确判断：没传 / null / 空数组 时设置默认值。
	if v, ok := req.Data["tags"]; !ok || v == nil {
		req.Data["tags"] = []interface{}{int64(4)}
	} else if s, ok := v.([]interface{}); ok && len(s) == 0 {
		req.Data["tags"] = []interface{}{int64(4)}
	}

	row := model.Project{
		CoverURL: util.GetPtrString(req.Data, "cover_url"),

		Name:           util.GetString(req.Data, "name"),
		TranslatedName: util.GetPtrString(req.Data, "translated_name"),
		Author:         util.GetPtrString(req.Data, "author"),
		SourceSite:     util.GetPtrString(req.Data, "source_site"),

		Description:            util.GetPtrString(req.Data, "description"),
		TranslationDescription: util.GetPtrString(req.Data, "translation_description"),
	}

	if config.DB.Transaction(func(tx *gorm.DB) error {
		// JSON 反序列化到 map[string]interface{} 时，数字通常会变成 float64，
		// 直接做 (int32) 断言会 panic；同时 model.Project.AdminMax 是 int。
		if v, ok := req.Data["admin_max"]; ok && v != nil {
			row.AdminMax = util.AsInt(v, 1)
		}
		if v, ok := req.Data["source_max"]; ok && v != nil {
			row.SourceMax = util.AsInt(v, 1)
		}
		if v, ok := req.Data["translator_max"]; ok && v != nil {
			row.TranslatorMax = util.AsInt(v, 1)
		}
		if v, ok := req.Data["supervisor_max"]; ok && v != nil {
			row.SupervisorMax = util.AsInt(v, 1)
		}
		if v, ok := req.Data["proofreader_max"]; ok && v != nil {
			row.ProofreaderMax = util.AsInt(v, 1)
		}
		if v, ok := req.Data["typesetter_max"]; ok && v != nil {
			row.TypesetterMax = util.AsInt(v, 1)
		}

		projID := util.GetInt64(req.Data, "project_id")
		err := config.DB.Where("id = ?", projID).Updates(&row).Error
		if err != nil {
			return err
		}

		if err := tx.Delete(&model.ProjectTag{}, "project_id = ?", projID).Error; err != nil {
			return err
		}

		tags, ok := util.AsSlice(req.Data["tags"])
		if !ok {
			return gorm.ErrInvalidData
		}

		for _, tagID := range tags {
			proj_tags_row := model.ProjectTag{
				ProjectID: projID,
				TagID:     util.AsInt64(tagID, 0),
			}
			if proj_tags_row.TagID == 0 {
				return gorm.ErrInvalidData
			}
			if !util.ProjectTagInsert(proj_tags_row.ProjectID, proj_tags_row.TagID) {
				return tx.Error
			}
		}

		return nil
	}) == nil {
		util.SendJSON(c, 0, "修改项目成功", []gin.H{
			{
				"project_id": row.ID,
				"name":       req.Data["name"],
				"created_at": row.CreatedAt,
			}}, 1, 1, c.FullPath(), c.Request.Method)
	} else {
		util.SendJSON(c, -1, "修改项目失败：修改数据失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}
}
