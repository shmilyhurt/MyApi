package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const (
	SuccessCode = 0
	ErrorCode   = -1
)

// 自定义成功
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: data,
	})
}

// 自定义错误码
func ErrorWithCode(c *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		data = gin.H{} // 空对象，而不是 null
	}
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
