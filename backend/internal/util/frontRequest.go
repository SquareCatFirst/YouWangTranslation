package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderByItem struct {
	Field string `json:"field"` // "asc" 或 "desc"
}

type BaseRequest struct {
	Action     string                 `json:"action"`     // 操作动作
	Sets       []string               `json:"sets"`       // 可选字段集合
	OrderBy    []OrderByItem          `json:"orderBy"`    // 排序
	Page       int                    `json:"page"`       // 页码
	PageSize   int                    `json:"pageSize"`   // 页大小
	Data       map[string]interface{} `json:"data"`       // 业务数据
	Filter     map[string]interface{} `json:"filter"`     // 过滤条件
	AuthFilter map[string]interface{} `json:"authFilter"` // 权限过滤
}

func ParseRequest(c *gin.Context) (*BaseRequest, bool) {
	var req BaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     -1,
			"msg":        "请求参数错误",
			"data":       []interface{}{},
			"rowCount":   0,
			"totalCount": 0,
			"api":        c.FullPath(),
			"method":     c.Request.Method,
			"SN":         GenSN(),
		})
		return nil, false
	}
	return &req, true
}
