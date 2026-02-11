package project

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProj(c *gin.Context) {
	req, ok := util.GetParseRequest(c)
	if !ok {
		util.SendJSON(c, -1, "获取项目失败：解析请求失败", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}

	if v, ok := req.Data["project_id"]; !ok || v == nil {
		util.SendJSON(c, -1, "获取项目失败：缺少项目ID", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}

	// 类似 dict 的结构：外层 key -> 内层 dict（string -> any）
	// 例如：mpusr["123"] = map[string]any{"role": 1, "name": "alice"}

	projID := util.GetInt64(req.Data, "project_id")

	var proj model.Project
	err := config.DB.Where("id = ?", projID).First(&proj).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.SendJSON(c, -1, "获取项目失败：项目不存在", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
			return
		}
		util.SendJSON(c, -1, "获取项目失败：查询数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	row := proj

	type MemberAssign struct {
		UserID     int64 `json:"user_id"`
		AssignedBy int64 `json:"assigned_by"`
	}

	type ChapterAssignmentDTO struct {
		Role       int     `json:"role"`
		ChapterIDs []int64 `json:"chapter_ids"`
		UserIDs    []int64 `json:"user_ids"`
		AssignedBy int64   `json:"assigned_by"` // 示例里有这个字段
	}

	type ChapterImageDTO struct {
		ImageID    int64  `json:"image_id,omitempty"` // 如果协定不需要可去掉
		ImageURL   string `json:"image_url"`
		OrderIndex int    `json:"order_index"`
	}

	type ProjectChapterDTO struct {
		ChapterID   int64             `json:"chapter_id"`
		ChapterName string            `json:"chapter_name,omitempty"`
		IsVisible   bool              `json:"is_visible,omitempty"`
		OrderIndex  int               `json:"order_index,omitempty"`
		Images      []ChapterImageDTO `json:"images"`
	}

	mp := make(map[string]any)
	mp["cover_url"] = row.CoverURL
	mp["name"] = row.Name
	mp["translated_name"] = row.TranslatedName
	mp["author"] = row.Author
	mp["source_site"] = row.SourceSite
	mp["description"] = row.Description
	mp["translation_description"] = row.TranslationDescription
	mp["admin_max"] = row.AdminMax
	mp["supervisor_max"] = row.SupervisorMax
	mp["typesetter_max"] = row.TypesetterMax
	mp["proofreader_max"] = row.ProofreaderMax
	mp["translator_max"] = row.TranslatorMax
	mp["source_max"] = row.SourceMax
	mp["tags"] = []int64{}
	mp["admin"] = []MemberAssign{}
	mp["supervisor"] = []MemberAssign{}
	mp["typesetter"] = []MemberAssign{}
	mp["proofreader"] = []MemberAssign{}
	mp["translator"] = []MemberAssign{}
	mp["source"] = []MemberAssign{}
	mp["chapter_assignments"] = []ChapterAssignmentDTO{}
	mp["project_chapters"] = []ProjectChapterDTO{}

	var tags []model.ProjectTag
	err = config.DB.Where("id = ?", projID).Find(&tags).Error
	if err == nil {
		for _, item := range tags {
			mp["tags"] = append(mp["tags"].([]int64), item.TagID)
		}
	}

	var ass []model.ProjectAssignment
	err = config.DB.Where("id = ?", projID).Find(&ass).Error
	if err == nil {
		for _, item := range ass {

			var p int64
			if item.AssignedBy == nil {
				p = 0
			} else {
				p = *item.AssignedBy
			}

			if item.Role == 0 {
				mp["admin"] = append(mp["admin"].([]MemberAssign), MemberAssign{UserID: item.UserID, AssignedBy: p})
			} else if item.Role == 1 {
				mp["source"] = append(mp["source"].([]MemberAssign), MemberAssign{UserID: item.UserID, AssignedBy: p})
			} else if item.Role == 2 {
				mp["translator"] = append(mp["translator"].([]MemberAssign), MemberAssign{UserID: item.UserID, AssignedBy: p})
			} else if item.Role == 3 {
				mp["proofreader"] = append(mp["proofreader"].([]MemberAssign), MemberAssign{UserID: item.UserID, AssignedBy: p})
			} else if item.Role == 4 {
				mp["typesetter"] = append(mp["typesetter"].([]MemberAssign), MemberAssign{UserID: item.UserID, AssignedBy: p})
			} else if item.Role == 5 {
				mp["supervisor"] = append(mp["supervisor"].([]MemberAssign), MemberAssign{UserID: item.UserID, AssignedBy: p})
			}
		}
	}

	chapter_ids := []int64{}
	var chapters []model.ProjectChapter

	var chapter_mp = map[int64]ProjectChapterDTO{}
	err = config.DB.Where("id = ?", projID).Find(&chapters).Error
	if err == nil {
		for _, item := range chapters {
			chapter_ids = append(chapter_ids, item.ID)
			chapter_mp[item.ID] = ProjectChapterDTO{
				ChapterID:   item.ID,
				ChapterName: item.ChapterName,
				IsVisible:   item.IsVisible,
				OrderIndex:  item.OrderIndex,
				Images:      []ChapterImageDTO{},
			}

			var img []model.ChapterWorkImage
			err = config.DB.Where("id = ?", item.ID).Find(&img).Error
			if err == nil {
				for _, itm := range img {
					dto := chapter_mp[item.ID] // 取出副本
					dto.Images = append(dto.Images, ChapterImageDTO{
						ImageID:    itm.ID,
						ImageURL:   itm.ImageURL,
						OrderIndex: itm.OrderIndex,
					})
					chapter_mp[item.ID] = dto // 写回 map
				}
			}
		}
	}

	type mem struct {
		Role       int
		AssignedBy int64
	}
	var chapter_id_map = map[mem][]int64{}
	var usr_id_map = map[mem][]int64{}

	for _, cid := range chapter_ids {
		var row []model.ChapterAssignment
		err := config.DB.Where("id = ?", cid).Find(&row).Error
		if err == nil {
			for _, item := range row {

				var p int64
				if item.AssignedBy == nil {
					p = 0
				} else {
					p = *item.AssignedBy
				}
				chapter_id_map[mem{Role: item.Role, AssignedBy: p}] = append(chapter_id_map[mem{Role: item.Role, AssignedBy: p}], cid)
				usr_id_map[mem{Role: item.Role, AssignedBy: p}] = append(usr_id_map[mem{Role: item.Role, AssignedBy: p}], item.UserID)
			}
		}
	}

	for key, item := range chapter_id_map {
		mp["chapter_assignments"] = append(mp["chapter_assignments"].([]ChapterAssignmentDTO), ChapterAssignmentDTO{
			Role:       key.Role,
			ChapterIDs: item,
			UserIDs:    usr_id_map[key],
			AssignedBy: key.AssignedBy,
		})
	}
	util.SendJSON(c, 0, "获取项目成功", []interface{}{mp}, 1, 1, c.FullPath(), c.Request.Method)
}
