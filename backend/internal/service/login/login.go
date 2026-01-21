package service

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		return
	}

	username, ok1 := req.Data["username"].(string)
	password, ok2 := req.Data["password"].(string)
	if !ok1 || !ok2 {
		util.SendJSON(c, -1, "用户名或密码缺失", []interface{}{}, 0, "/api/v1/auth/login", "POST")
		return
	}

	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		util.SendJSON(c, -1, "用户名不存在", []interface{}{}, 0, "/api/v1/auth/login", "POST")
		return
	}

	if user.PasswordHash == nil || !util.CheckPassword(*user.PasswordHash, password) {
		util.SendJSON(c, -1, "密码错误", []interface{}{}, 0, "/api/v1/auth/login", "POST")
		return
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLoginAt = &now
	config.DB.Save(&user)

	util.SendJSON(c, 0, "登录成功", []gin.H{{"id": user.ID}}, 1, "/api/v1/auth/login", "POST")
}
