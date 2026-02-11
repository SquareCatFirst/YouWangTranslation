package user

import (
	"strconv"

	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/config"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/model"
	"github.com/SquareCatFirst/YouWangTranslation/backend/internal/util"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	// 1) page / pageSize
	// page, err1 := strconv.Atoi(c.DefaultQuery("page", "1"))
	// pageSize, err2 := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	//用DefultQueryRequest实现
	pageStr, ok1 := util.DefultQueryRequest(c, "page", "1")
	pageSizeStr, ok2 := util.DefultQueryRequest(c, "pageSize", "20")
	if !ok1 || !ok2 {
		util.SendJSON(c, -1, "获取用户列表失败：解析请求失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}
	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err1 != nil || err2 != nil || page <= 0 || pageSize <= 0 {
		util.SendJSON(c, -1, "分页参数错误", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}
	// fmt.Printf("DefultQueryRequest: page=%d, pageSize=%d\n", page1, pageSize1) // 打印分页参数，方便调试

	//用GetParseRRequest实现
	// req, ok := util.GetParseRequest(c)
	// if !ok {
	// 	util.SendJSON(c, -1, "获取用户列表失败：解析请求失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }
	// page := req.Page
	// pageSize := req.PageSize
	// // fmt.Printf("UserList: page=%d, pageSize=%d\n", page, pageSize) // 打印分页参数，方便调试
	// if page <= 0 || pageSize <= 0 {
	// 	util.SendJSON(c, -1, "分页参数错误", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
	// 	return
	// }

	// 2) total
	var total int64
	if err := config.DB.Model(&model.User{}).Count(&total).Error; err != nil {
		util.SendJSON(c, -1, "获取用户总数失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}
	if total == 0 {
		util.SendJSON(c, 0, "获取成功", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	// 3) list
	offset := (page - 1) * pageSize
	var users []model.User
	if err := config.DB.
		Model(&model.User{}).
		Order("id DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&users).Error; err != nil {
		util.SendJSON(c, -1, "获取用户列表失败", []interface{}{}, 0, 0, c.FullPath(), c.Request.Method)
		return
	}

	// 4) response data (只返回你结构体里存在的字段；不返回 password_hash)
	list := make([]gin.H, 0, len(users))
	for _, u := range users {
		list = append(list, gin.H{
			"id":           u.ID,
			"avatarUrl":    u.AvatarURL,
			"nickname":     u.Nickname,
			"role":         u.Role,
			"qqOpenId":     u.QQOpenID,
			"status":       u.Status,
			"username":     u.Username, // 不想暴露就删掉这一行
			"registeredAt": u.RegisteredAt,
			"lastLoginAt":  u.LastLoginAt,
		})
	}

	// count 我这里返回 total（更符合分页）；如果你 util.SendJSON 的 count 表示“本页数量”，改成 len(list)
	util.SendJSON(c, 0, "获取成功", list, len(list), int(total), c.FullPath(), c.Request.Method)
}
