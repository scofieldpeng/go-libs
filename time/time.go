//time包内主要为一些常用的时间类库包，通过对官方包中time包的一些封装提升时间操作的便利性
package time

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Sec   int
}

//Timestamp将time.Time类型的值转化为时间戳
//t传值为字符串，格式为"1970-12-12 12:12:12"的当地时间，如果空字符串，默认为当前时间
func Timestamp(s string) (timestamp int64, err error) {
	err = nil
	timestamp = 0

	var t time.Time

	if len(s) == 0 {
		t = time.Now()
	} else {
		t, err = parseDate(s)
	}
	if err != nil {
		return
	}

	timestamp = t.Unix()
	return
}

//Mktime返回相应时间的时间戳
func Mktime(hour, minute, second, month, day, year int) int64 {
	var monthArr = [12]time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	localLocation, _ := time.LoadLocation("UTC")
	t := time.Date(year, monthArr[month-1], day, hour, minute, second, 0, localLocation)

	return t.Unix()
}

//Getdate与PHP的getdate函数相同,不同的是返回的map中不含有键值0的数据
//seconds，minutes，hours，mday，wday，
func Getdate(timestamp int64) (res map[string]interface{}) {
	if timestamp == -1 {
		timestamp = time.Now().Unix()
	}

	timeIns := time.Unix(timestamp, 0)

	monthToNum := map[time.Month]int{
		time.January:   1,
		time.February:  2,
		time.March:     3,
		time.April:     4,
		time.May:       5,
		time.June:      6,
		time.July:      7,
		time.August:    8,
		time.September: 9,
		time.October:   10,
		time.November:  11,
		time.December:  12,
	}

	weekdayToNum := map[time.Weekday]int{
		time.Sunday:    0,
		time.Monday:    1,
		time.Tuesday:   2,
		time.Wednesday: 3,
		time.Thursday:  4,
		time.Friday:    5,
		time.Saturday:  6,
	}

	res = make(map[string]interface{})

	res["year"], res["month"], res["mday"] = timeIns.Date()
	res["hours"], res["minutes"], res["seconds"] = timeIns.Clock()
	res["weekday"] = timeIns.Weekday()
	res["yday"] = timeIns.YearDay()
	res["mon"] = monthToNum[res["month"].(time.Month)]
	res["wday"] = weekdayToNum[res["weekday"].(time.Weekday)]
	return
}

//Format用来将时间戳按照一定的格式进行输出，与PHP的date()所做的事情一样
//目前所支持的格式:
//年：
//Y 2015
//y 15
//月：
//F January
//m 01-12
//M Jan
//n 1-12
//日期：
//d 01-31
//j 1-31
//D Mon-Sun
//l Monday-Sunday
//时
//g 1-12
//G 0-23
//h 01-12
//H 00-23
//分钟
//i 00-59
//秒
//s 00-59
//上/下午
//a am/pm
//A AM/PM
func Format(format string, timestamp int64) string {
	dateFormat := map[string]string{
		"Y": "2006",
		"y": "06",
		"F": "January",
		"m": "01",
		"n": "1",
		"d": "02",
		"j": "2",
		"D": "Mon",
		"l": "Monday",
		"g": "3",
		"H": "15",
		"h": "03",
		"i": "04",
		"s": "05",
		"a": "pm",
		"A": "PM",
	}
	//转换格式
	for k, v := range dateFormat {
		format = strings.Replace(format, k, v, -1)
	}

	//将时间戳转化为string
	return time.Unix(timestamp, 0).Format(format)
}

//parseDate 将字符串转换为time.Time格式，字符串的格式必须为"2006-01-02 15:04:05"格式
func parseDate(ori string) (t time.Time, err error) {
	err = nil

	var date Date
	month := [12]time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}

	//分割日期和时间
	dateArr := strings.Split(ori, " ")
	if len(dateArr) != 2 {
		err = errors.New("format is unknown,please use format:1970-12-12 12:12:12!")
		return
	}

	//解析具体日期
	dateSlice := strings.Split(dateArr[0], "-")
	if len(dateSlice) != 3 {
		err = errors.New("date format is unknown,please use format:1970-12-12 12:12:12!")
		return
	}
	date.Year, _ = strconv.Atoi(dateSlice[0])
	date.Month, _ = strconv.Atoi(dateSlice[1])
	date.Day, _ = strconv.Atoi(dateSlice[2])

	//解析具体时间
	timeSlice := strings.Split(dateArr[1], ":")
	if len(timeSlice) != 3 {
		err = errors.New("time format is unknown,please use format:1970-12-12 12:12:12!")
		return
	}
	date.Hour, err = strconv.Atoi(timeSlice[0])
	if err != nil {
		return
	}
	date.Min, err = strconv.Atoi(timeSlice[1])
	if err != nil {
		return
	}
	if month[date.Month-1] == 0 {
		err = errors.New("month is exceed Maxium! 1 to 12 is correct!")
		return
	}
	date.Sec, _ = strconv.Atoi(timeSlice[2])
	nowLocation, _ := time.LoadLocation("UTC")

	t = time.Date(date.Year, month[date.Month-1], date.Day, date.Hour, date.Min, date.Sec, 0, nowLocation)

	return
}
