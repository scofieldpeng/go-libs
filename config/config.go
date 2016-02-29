// Package config 用来读取ini程序配置项目,并且能够自动监听配置文件变化触发相应方法
//
// 用法:
// 1. 初始化
//     if err := config.Init(path,debug);err != nil{
//        log.Fatal(err)
//     }
//
// 2. 读取配置
//     val,ok := config.Config.Read("section:name"[,defaultVal])
//     intVal,ok := config.Config.ReadInt("section:name"[,defaultIntVal])
//     int64Val,ok := config.Config.ReadInt64("section:name"[,defaultInt64Val])
//     boolVal,ok := config.Config.ReadBool("section:name"[,defaultBoolVal])
//     sectionVal,ok := config.Config.ReadSection("section")
//
// 3. 如果需要监听某个配置文件的变化采取不同的动作,当某个配置文件变化时会通知且出发回调的执行
//    第一个参数为某个section中的某个name配置项的变化,如果要监听某个section内任意一个配置
//    项的变化,第一个参数设置为section:*,如果要监听多个数据,格式为section1:name1|section2:name2
//    第二个参数为配置变化后的回调函数,需要设置为func(string)格式
//    config.Config.AddEventListener("section:name",func(newValue))
// 4. 重新载入配置
//    err := config.Config.Reload()
package config

import "github.com/vaughan0/go-ini"

// config 配置文件结构体
type config struct {
    curCache ini.File  // 当前配置
    lastCache ini.File // 上次配置
    events map[string]func(string) // 事件监听配置
}

var(
    Config config
)

func init() {
    Config = config{}
}

// init TODO 读取配置文件,如果读取失败,返回错误
func (c *config) Init(path string) error {
    return nil
}

// Read TODO 读取配置文件
func (c *config) Read(node string,defaultVal...string)(string,bool){
    return "",true
}

// ReadInt TODO 读取int值的配置文件
func (c *config) ReadInt(node string,defaultVal ...int)(int,bool){
    return 0,true
}

// ReadInt64 TODO 读取int64类型值的配置文件
func (c *config) ReadInt64(node string,defaultVal ...int64)(int64,bool){
    return 0,true
}

// ReadBool TODO 读取bool类型值的配置文件
func (c *config) ReadBool(node string,defaultVal ...bool)(bool,bool){
    return false,true
}

// ReadSection TODO 读取section类型的配置文件
func (c *config) ReadSection(node string)(ini.Section,bool){
    return ini.Section{},true
}

// AddEventListener TODO 添加节点的时间监听
func (c *config) AddEventListener(node string,callback func(string)) error {
    return nil
}

// RmEventListener 移除节点的监听
func (c *config) RmEventListener(node string) error {
    return nil
}