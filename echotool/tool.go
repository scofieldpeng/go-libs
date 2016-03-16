package echotool

import (
	"github.com/scofieldpeng/echo"
	"strings"
)

// GetRealIp 获取用户真实的IP
func GetRealIp(ctx *echo.Context) string {
	host := ctx.Request().RemoteAddr
	if host == "" {
		host = ctx.Request().Header.Get("X-Forward-For")
	}
	if host == "" {
		host = ctx.Request().Header.Get("Proxy-Client-IP")
	}
	if host == "" {
		host = ctx.Request().Header.Get("WL-Proxy-Client-IP")
	}

	if host != "" {
		hostSplit := strings.Split(host, ":")
		if len(hostSplit) > 0 {
			host = hostSplit[0]
		}
	}

	return host
}
