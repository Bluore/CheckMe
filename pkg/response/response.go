package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	StatusOK     = responseCode{Code: 200, Message: "成功"}
	BadRequest   = responseCode{Code: 400, Message: "请求错误"}
	ParamMissing = responseCode{Code: 400, Message: "参数错误"}
	Unauthorized = responseCode{Code: 401, Message: "身份未验证"}
	Forbidden    = responseCode{Code: 403, Message: "禁止访问"}
	NotFound     = responseCode{Code: 404, Message: "资源不存在"}
	ServerError  = responseCode{Code: 500, Message: "服务器内部错误"}
)

type responseCode struct {
	Code    int
	Message string
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (rc responseCode) ToResponse(data interface{}) Response {
	return Response{
		Code:    rc.Code,
		Message: rc.Message,
		Data:    data,
	}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, StatusOK.ToResponse(data))
}

func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: msg,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func Fail(c *gin.Context, code responseCode) {
	c.JSON(http.StatusOK, code.ToResponse(nil))
}
