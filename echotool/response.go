// echotool 包,echo框架的一些便携操作
package echotool

import (
	"github.com/scofieldpeng/echo"
	"net/http"
)

// EchoJsonOk json格式返回200状态码下的数据
func EchoJsonOk(ctx *echo.Context, data interface{}) error {
	if assertData, ok := data.(string); ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"data": assertData})
	}
	return ctx.JSON(http.StatusOK, data)
}

// EchoJsonError json格式返回httpStatusCode的HTTP状态码下的错误信息,参数errCode为具体的错误码,msg为string类型的错误原因
func EchoJsonError(ctx *echo.Context, httpStatusCode, errCode int, msg string) error {
	errJson := map[string]interface{}{"errcode": errCode, "errmsg": msg}
	return ctx.JSON(httpStatusCode, errJson)
}