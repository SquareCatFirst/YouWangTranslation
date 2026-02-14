package project

import (
	"errors"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func roleNameToCode(roleName string) (int, string, string, string, bool) {
	switch roleName {
	case "admin":
		return 0, "admin", "admin_max", "admin_max", true
	case "source":
		return 1, "source", "source_max", "source_max", true
	case "translator":
		return 2, "translator", "translator_max", "translator_max", true
	case "proofreader":
		return 3, "proofreader", "proofreader_max", "proofreader_max", true
	case "typesetter":
		return 4, "typesetter", "typesetter_max", "typesetter_max", true
	case "supervisor":
		return 5, "supervisor", "supervisor_max", "supervisor_max", true
	default:
		return 0, "", "", "", false
	}
}

func AddProjAss(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		util.SendJSON(c, -1, "添加权限失败：解析请求失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	projID := util.GetInt64(req.Data, "project_id")
	if projID == 0 {
		util.SendJSON(c, -1, "添加权限失败: 缺少项目ID", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	roleName := util.GetString(req.Data, "role")
	if roleName == "" {
		util.SendJSON(c, -1, "添加权限失败: 缺少添加角色", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	role, memberKey, maxColumn, reqMaxKey, ok := roleNameToCode(roleName)
	if !ok {
		util.SendJSON(c, -1, "添加权限失败: 角色不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	membersRaw, ok := util.AsSlice(req.Data[memberKey])
	if !ok || len(membersRaw) == 0 {
		util.SendJSON(c, -1, "添加权限失败: 缺少角色成员数据", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	memberRows := make([]model.ProjectAssignment, 0, len(membersRaw))
	for _, item := range membersRaw {
		m, ok := util.AsMap(item)
		if !ok {
			util.SendJSON(c, -1, "添加权限失败: 成员数据格式不合法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}
		userID := util.AsInt64(m["user_id"], 0)
		assignedBy := util.AsInt64(m["assigned_by"], 0)
		if userID == 0 || assignedBy == 0 {
			util.SendJSON(c, -1, "添加权限失败: user_id 或 assigned_by 非法", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}
		assignedByCopy := assignedBy
		memberRows = append(memberRows, model.ProjectAssignment{
			ProjectID:  projID,
			UserID:     userID,
			Role:       role,
			AssignedBy: &assignedByCopy,
		})
	}

	row := model.Project{}
	if err := config.DB.Where("id = ?", projID).First(&row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.SendJSON(c, -1, "添加权限失败: 项目不存在", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
			return
		}
		util.SendJSON(c, -1, "添加权限失败: 查询数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	var roleNum int64
	if config.DB.Model(&model.ProjectAssignment{}).Where("project_id = ? AND role = ?", projID, role).Count(&roleNum).Error != nil {
		util.SendJSON(c, -1, "添加权限失败: 查询数据库失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	if config.DB.Transaction(func(tx *gorm.DB) error {
		newMax := int(roleNum) + len(memberRows)
		maxValue := max(newMax, util.GetInt(req.Data, reqMaxKey))
		if err := tx.Model(&model.Project{}).
			Where("id = ?", projID).
			Updates(map[string]interface{}{maxColumn: maxValue}).Error; err != nil {
			return err
		}

		for _, assRow := range memberRows {
			if !util.ProjectAssignmentInsert(tx, assRow.ProjectID, assRow.UserID, assRow.Role, assRow.AssignedBy) {
				return errors.New("添加权限失败: 写入权限记录失败")
			}
		}
		return nil
	}) == nil {
		util.SendJSON(c, 0, "添加权限成功", []gin.H{{}}, len(memberRows), len(memberRows), c.FullPath(), c.Request.Method)
	} else {
		util.SendJSON(c, -1, "添加权限失败：添加数据失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}
}
