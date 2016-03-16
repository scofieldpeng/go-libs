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

import (
	"errors"
	"github.com/scofieldpeng/go-libs/filewatcher"
	"github.com/vaughan0/go-ini"
	"strconv"
	"strings"
)

// config 配置文件结构体
type config struct {
	path          string                                     // 配置文件路径
	curCache      ini.File                                   // 当前配置
	lastCache     ini.File                                   // 上次配置
	nodeEvents    map[string]map[string]func(string, string) // 事件监听配置,外面的map的key代表配置文件节点,value为该节点的事件回调,value为一个map,其key为回调函数
	sectionEvents map[string]map[string]func(string, string) // section变化监听配置,外面的map的key代表section节点,value为该节点的事件回调,value为一个map,其key为回调函数
}

var (
	Config config
)

func init() {
	Config = config{}
}

// init 读取配置文件,如果读取失败,返回错误
func (c *config) Init(path string) error {
	c.path = path

	// 载入配置
	if err := c.load(); err != nil {
		return err
	}
	// 载入文件监听
	if err := c.configChangeListen(); err != nil {
		return err
	}

	return nil
}

// Reload 重新载入配置
func (c *config) Reload() error {
	return c.load()
}

// Rollback 回退配置
func (c *config) Rollback() {
	c.curCache = c.lastCache
}

// load 载入配置
func (c *config) load() error {
	file, err := ini.LoadFile(c.path)
	if err != nil {
		return err
	}

	c.lastCache = c.curCache
	c.curCache = file
	return nil
}

// configChangeListen 配置文件变化监听
func (c *config) configChangeListen() error {
	return filewatcher.FWatcher.AddFile(c.path, c.configChangeTrigger)
}

// configChangeTrigger 配置文件变化时的触发器
func (c *config) configChangeTrigger(path string) error {
	// 重新载入文件
	if err := c.load(); err != nil {
		return errors.New("自动刷新配置文件(" + c.path + ")失败,错误原因:" + err.Error())
	}

	changedSections := make([]string, 0)
	changedNodes := make([]string, 0)
	changedNodesValue := make(map[string][]string, 0) // 变化节点的新旧配置,offset为0的值为旧值,1为新值

	// 遍历文件所有配置变化,然后进行比对,当发现有变化时如果有监听,依次进行回调
	for nodeKey, section := range c.lastCache {
		for subNodeKey, preValue := range section {
			sectionRecord := false
			// 发现变化
			if curValue, ok := c.Read(nodeKey + ":" + subNodeKey); ok && curValue != preValue {
				if !sectionRecord {
					changedSections = append(changedSections, nodeKey)
				}
				changedNodes = append(changedNodes, nodeKey+":"+subNodeKey)
				changedNodesValue[nodeKey+":"+subNodeKey] = make([]string, 2)
				changedNodesValue[nodeKey+":"+subNodeKey][0] = preValue
				changedNodesValue[nodeKey+":"+subNodeKey][1] = curValue
			}
		}
	}

	// 触发section监听
	for _, sectionName := range changedSections {
		if callbacks, ok := c.sectionEvents[sectionName]; ok {
			for _, callback := range callbacks {
				go callback("", "")
			}
		}
	}
	// 触发node监听
	for _, nodeName := range changedNodes {
		if callbacks, ok := c.sectionEvents[nodeName]; ok {
			for _, callback := range callbacks {
				go callback(changedNodesValue[nodeName][0], changedNodesValue[nodeName][1])
			}
		}
	}

	return nil
}

// Read 读取配置文件,返回string格式的配置,如果出错,第二个参数为false
func (c *config) Read(node string, defaultVal ...string) (string, bool) {
	nodeArr := make([]string, 2)
	tmp := strings.Split(node, ":")
	if len(tmp) == 2 {
		nodeArr = tmp
	} else {
		nodeArr[0] = ""
		nodeArr[1] = tmp[0]
	}
	return c.curCache.Get(nodeArr[0], nodeArr[1])
}

// ReadInt 读取int值的配置文件,如果出错,第二个参数为false
func (c *config) ReadInt(node string, defaultVal ...int) (int, bool) {
	tmp, ok := c.Read(node)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], true
		}
		return 0, false
	}
	res, err := strconv.Atoi(tmp)
	if err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0], true
		}
		return 0, false
	}
	return res, true
}

// ReadInt64 读取int64类型值的配置文件
func (c *config) ReadInt64(node string, defaultVal ...int64) (int64, bool) {
	tmp, ok := c.Read(node)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], true
		}
		return 0, false
	}
	res, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0], true
		}
		return 0, false
	}
	return res, true
}

// ReadBool 读取bool类型值的配置文件
func (c *config) ReadBool(node string, defaultVal ...bool) (bool, bool) {
	tmp, ok := c.Read(node)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], true
		}
		return false, false
	}
	res, err := strconv.ParseBool(tmp)
	if err != nil {
		if len(defaultVal) > 0 {
			return defaultVal[0], true
		}
		return false, false
	}
	return res, true
}

// ReadSection 读取section类型的配置文件
func (c *config) ReadSection(node string) (ini.Section, bool) {
	hasFind := false
	section := c.curCache.Section(node)
	if len(section) > 0 {
		hasFind = true
	}
	return section, hasFind
}

// AddSectionListener 添加section的事件监听
func (c *config) AddSectionListener(section string, callbackName string, callback func(string, string)) {
	if _, ok := c.sectionEvents[section]; !ok {
		c.sectionEvents[section] = make(map[string]func(string, string))
	}
	c.sectionEvents[section][callbackName] = callback
}

// AddNodeListener 添加节点的事件监听
func (c *config) AddNodeListener(node string, callbackName string, callback func(string, string)) {
	if _, ok := c.sectionEvents[node]; !ok {
		c.nodeEvents[node] = make(map[string]func(string, string))
	}
	c.nodeEvents[node][callbackName] = callback
}

// RmNodeLisnter 移除某节点的具体某个监听事件
func (c *config) RmNodeListener(node, callbackName string) {
	if _, ok := c.sectionEvents[node]; ok {
		delete(c.sectionEvents[node], callbackName)
	}
}

// ClearNodeListener 清除某节点所有的监听事件
func (c *config) ClearNodeListener(node string) {
	delete(c.sectionEvents, node)
}

// RmEventListener 移除某section的具体某个监听
func (c *config) RmSectionListener(section, callbackName string) {
	if _, ok := c.sectionEvents[section]; ok {
		delete(c.sectionEvents[section], callbackName)
	}
}

// ClearSectionListener 清除某section的所有监听事件
func (c *config) ClearSectionListener(section string) {
	delete(c.sectionEvents, section)
}
