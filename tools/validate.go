package tools

import (
	"regexp"
)

// IsMobile 判断是否是手机号,返回bool值
func IsMobile(mobile string) bool {
	ok, _ := regexp.MatchString(`^1[3|4|5|7|8]\d{9}$`, mobile)
	return ok
}

// IsEmail 判断是否是合法邮箱,返回bool值
func IsEmail(email string) bool {
	ok, _ := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, email)
	return ok
}
