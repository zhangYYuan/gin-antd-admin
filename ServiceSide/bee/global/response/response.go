package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int 		 	`json:"code"`
	Data interface{} 	`json:"data"`
	Msg  string		 	`json:"msg"`
}

const (
	ERROR = -1
	SUCEESS = 0
)

func Result(code int, data interface{}, msg string, c * gin.Context)  {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Success(c *gin.Context) {
	Result(SUCEESS, map[string]interface{}{}, "操作成功", c)
}
func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCEESS, data, "操作成功", c)
}
func SuccessWithDescription(data interface{}, msg string, c *gin.Context)  {
	Result(SUCEESS, data, msg, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
func FailWithDescription(code int, data interface{}, msg string, c *gin.Context)  {
	Result(code, data, msg, c)
}

