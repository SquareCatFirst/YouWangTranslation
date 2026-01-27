package cookie

import (
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterSession(r *gin.Engine) {

	// 用一个强随机 secret（建议放到 config 里）
	store := cookie.NewStore([]byte(config.Cfg.Session.Secret)) // 自己放配置里
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   36000, // 1小时
		HttpOnly: true,  // 防 XSS 读 cookie
		Secure:   false, // 线上 HTTPS 要 true；本地可 false
		SameSite: 2,     // 2=SameSiteLaxMode（go1.20+ 的 net/http 是枚举；这里 gin-contrib 用 int）
	})

	r.Use(sessions.Sessions("sid", store)) // "sid" 是 cookie 名

}
