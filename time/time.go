//time包内主要为一些常用的时间类库包，通过对官方包中time包的一些封装提升时间操作的便利性
package time

import (
	"time"
)

// GetMonth 将数字形式的月份转化为time.Month结构,入股输入的月份不正确，则默认返回一月份
func GetMonth(monthNum int) time.Month {
	if monthNum<1 || monthNUm>12{
		return time.January
	}
	return time.Month(monthNum)
}

// GetMonthInt 将time.Month格式的月份转化为int型
func GetMonthInt(month time.Month) int {
	return int(month)
}
