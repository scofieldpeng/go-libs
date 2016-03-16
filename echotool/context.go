package echotool

import (
	"encoding/json"
	"github.com/scofieldpeng/echo"
	"io/ioutil"
)

// GetBody 返回request请求中body的内容,如果出错,第二个参数为error
func GetBody(ctx *echo.Context) ([]byte, error) {
	return ioutil.ReadAll(ctx.Request().Body)
}

// ParseBodyToStruct 将body的内容解析到ouput中,如果出错,返回error
func ParseBodyToStruct(ctx *echo.Context, output interface{}) error {
	body, err := GetBody(ctx)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}

	return json.Unmarshal(body, output)
}

// GetMultiForm 获取多媒体文件 TODO
func GetMultiForm(ctx *echo.Context) {
}
