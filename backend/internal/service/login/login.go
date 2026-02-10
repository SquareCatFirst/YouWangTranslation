package login

import (
	"time"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req, ok := util.ParseRequest(c)
	if !ok {
		return
	}

	username, ok1 := req.Data["username"].(string)
	password, ok2 := req.Data["password"].(string)
	if !ok1 || !ok2 {
		util.SendJSON(c, -1, "用户名或密码缺失", []interface{}{}, 0, 1, "/api/v1/auth/login", "POST")
		return
	}

	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		util.SendJSON(c, -1, "用户名不存在", []interface{}{}, 0, 1, "/api/v1/auth/login", "POST")
		return
	}

	if user.PasswordHash == nil || !util.CheckPassword(*user.PasswordHash, password) {
		util.SendJSON(c, -1, "密码错误", []interface{}{}, 0, 1, "/api/v1/auth/login", "POST")
		return
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLoginAt = &now
	config.DB.Save(&user)

	// ====== 关键：写 session（cookie 会通过 Set-Cookie 返回给浏览器）======
	sess := sessions.Default(c)

	// 防 session fixation：登录前清一下旧数据
	sess.Clear()

	// 存登录态（建议只存最小信息：user_id）
	sess.Set("user_id", user.ID)

	// 持久化（不 Save 就不会写回 cookie）
	if err := sess.Save(); err != nil {
		util.SendJSON(c, -1, "session保存失败", []interface{}{}, 0, 1, "/api/v1/auth/login", "POST")
		return
	}
	// ============================================================

	util.SendJSON(c, 0, "登录成功", []gin.H{{"id": user.ID}}, 1, 1, "/api/v1/auth/login", "POST")
}

func Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	// MaxAge=-1 会让 cookie 过期
	sess.Options(sessions.Options{MaxAge: -1, Path: "/"})
	_ = sess.Save()

	util.SendJSON(c, 0, "已退出", []interface{}{}, 0, 1, "/api/v1/auth/logout", "DELETE")
}
