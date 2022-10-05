package helper

import "github.com/gin-gonic/gin"

type Map map[string]interface{}

type Format struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func JSON(c *gin.Context, statusCode int, msg string, data interface{}) {
	var value Format
	value.Code = statusCode
	value.Msg = msg
	value.Data = data
	c.JSON(statusCode, value)
}
