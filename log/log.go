// package log 专门的log包管理工具
package log

// logger结构体信息
type logger struct {

}

// Info 输出info信息
func (l logger) Info() {}

// Warn 输出warn信息
func (l logger) Warn() {}

// Fetal 输出fetal信息
func (l logger) Fetal() {}

// Info 输出log信息,该函数一般用于调试信息等的输出
func Info(...interface{}) {

}

// Warn 输出warn信息,该函数一般用于警告信息的输出
func Warn(...interface{}) {

}

// Fetal 输出fetal错误(致命错误),该函数一般用户致命错误的输出
func Fetal(...interface{}) {

}


