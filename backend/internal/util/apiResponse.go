package util

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Status     int         `json:"status"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	RowCount   int         `json:"rowCount"`
	TotalCount int         `json:"totalCount"`
	API        string      `json:"api"`
	Method     string      `json:"method"`
	SN         string      `json:"SN"`
}

// SendJSON 统一返回函数
func SendJSON(c *gin.Context, status int, msg string, data interface{}, rowCount int, totalCount int, api, method string) {
	resp := APIResponse{
		Status:     status,
		Msg:        msg,
		Data:       data,
		RowCount:   rowCount,
		TotalCount: totalCount,
		API:        api,
		Method:     method,
		SN:         GenSN(),
	}

	c.JSON(200, resp)
}
