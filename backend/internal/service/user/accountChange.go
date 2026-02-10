package user

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
)

func AccountChange(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		util.SendJSON(c, -1, "请求参数错误", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}
	data := req.Data
	userID := util.GetInt64(data, "user_id")
	status := util.GetString(data, "account")
	if status != "enable" && status != "disable" {
		util.SendJSON(c, -1, "account参数错误，应为enable或disable", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}

	var user model.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		util.SendJSON(c, -1, "用户不存在", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}

	var newStatus int
	if status == "enable" {
		newStatus = 1
	} else {
		newStatus = 0
	}
	user.Status = newStatus
	if err := config.DB.Save(&user).Error; err != nil {
		util.SendJSON(c, -1, "更新用户状态失败", []interface{}{}, 0, 1, c.FullPath(), c.Request.Method)
		return
	}
	util.SendJSON(c, 0, "更新用户状态成功", []interface{}{}, 1, 1, c.FullPath(), c.Request.Method)
}
