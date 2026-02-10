package util

import (
	"encoding/json"
	"net/http"
	"strings"

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

func GetParseRequest(c *gin.Context) (*BaseRequest, bool) {
	var req BaseRequest
	raw := c.Query("request")
	//当row为空字符串时，json.Unmarshal会报错。

	// //没传 / null 时设置默认值。
	// if raw == "" || raw == "null" {
	// 	raw = "{}"
	// }
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
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
	// log.Printf("GetParseRequest%+v\n", req) // 打印解析后的请求对象，方便调试
	return &req, true
}

func DefultQueryRequest(c *gin.Context, key string, defultValue string) (string, bool) {
	raw := c.Query("request")
	dec := json.NewDecoder(strings.NewReader(raw))
	dec.UseNumber()
	var mp map[string]any
	if err := dec.Decode(&mp); err != nil {
		// 解析失败
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
		return "", false
	}
	// log.Printf("DefultQueryRequest：%+v\n", mp) // 打印解析后的请求对象，方便调试
	if v, ok := mp[key]; !ok || v == nil {
		return defultValue, true
	} else {
		switch v := v.(type) {
		case string:
			if v == "null" {
				return defultValue, true
			} else {
				return v, true
			}
		case json.Number:
			return v.String(), true
		default:
			return defultValue, true
		}
	}

	// //parse_json
	// m := map[string]interface{}{}
	// //没传 / null 时设置默认值。
	// if raw == "" || raw == "null" {
	// 	raw = "{}"
	// }
	// if err := json.Unmarshal([]byte(raw), &m); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":     -1,
	// 		"msg":        "请求参数错误111",
	// 		"data":       []interface{}{},
	// 		"rowCount":   0,
	// 		"totalCount": 0,
	// 		"api":        c.FullPath(),
	// 		"method":     c.Request.Method,
	// 		"SN":         GenSN(),
	// 	})
	// 	return 0, false
	// }
	// log.Printf("DefultQueryRequest：%+v\n", m) // 打印解析后的请求对象，方便调试
	// // set default value
	// if v, ok := m[key]; !ok || v == nil {
	// 	return defultValue, true
	// } else {
	// 	return int(v.(float64)), true
	// }
}
