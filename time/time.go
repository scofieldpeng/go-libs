//time包内主要为一些常用的时间类库包，通过对官方包中time包的一些封装提升时间操作的便利性
package time

import (
	"time"
)

// GetMonth 将数字形式的月份转化为time.Month结构,入股输入的月份不正确，则默认返回一月份
func GetMonth(monthNum int) time.Month {
	switch monthNum {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		return time.January
	}
}

// GetMonthInt 将time.Month格式的月份转化为int型
func GetMonthInt(month time.Month) int {
	switch month {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	default:
		return 1
	}
}
