// Package filewatcher 包用于监听监听文件变化
//
// 安装:
//    go get github.com/scofieldpeng/go-libs/filewatcher
//
// 使用方法:
//     import github.com/scofieldpeng/go-libs/filewatcher
//
// 1. 初始化
//     filewatcher.FWatcher.Init()
//
// 2. 添加某文件监听事件,第一个参数为要监听的文件/文件夹的绝对路径,第二个参数为要回调的函数名称,
//    注意,如果是文件夹,那么只能监听该文件夹下的第一级文件的变化
//     filewatcher.FWatcher.AddFile("filePath",callbackName,callback)     /* 监听文件 */
//     filewatcher.FWatcher.AddFolder("folderPath",callbackName,callback) /* 监听文件夹 */
//
// 3. 取消文件/文件夹监听
//     filewatcher.FWatcher.RmFile("filePath",callbackName) /* 移除对文件的监听 */
//     filewatcher.FWatcher.RmFolder("folderPath",callbackName) /* 移除对文件夹的监听 */
package filewatcher

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"log"
)

type (
	// fWatcher fWatcher结构体,用于实现文件监听的主要功能
	fWatcher struct {
		fileEvents   map[string]map[string]func(string) error // 文件监听事件回调
		folderEvents map[string]map[string]func(string) error // 文件夹监听事件回调
		status       int                                      // 当前监听程序的状态,init(初始化完毕),run(运行)
		watcher      *fsnotify.Watcher                        // fsnotify监听对象
		debug        bool                                     // 是否是调试模式
	}
)

var (
	STATUS_STOP int = 0 // 停止状态
	STATUS_INIT int = 1 // 初始化的状态
	STATUS_RUN  int = 2 // 运行的状态
)

var (
	FWatcher fWatcher
)

func init() {
	FWatcher = newFWatcher()
}

// newFWatcher 新建一个fWatcher对象结构体
func newFWatcher() fWatcher {
	return fWatcher{}
}

// Init 初始化文件监听
func (fw *fWatcher) Init(debug ...bool) error {
	var err error = nil

	if len(debug) > 0 {
		debug = make([]bool,1)
		debug[0] = false
	}

	fw.status = STATUS_INIT
	fw.watcher,err = fsnotify.NewWatcher()
	fw.debug = debug[0]
	if err != nil {
		return errors.New("初始化filewatcher失败,失败原因:" + err.Error())
	}
	return nil
}

// Run 开始文件监听
func (fw *fWatcher) Run() {
	defer fw.watcher.Close()
	fw.status = STATUS_RUN
	for{
		select {
		case event := <-fw.watcher.Events:
			if fw.debug {
				log.Println("本次发现变化文件:",event.String())
			}

			go fw.watcherTrigger(event.Name)
		}
	}
}

// watcherTrigger TODO 事件监听触发函数,该函数将便利用户设置的监听文件/文件夹,如果有对变化文件/文件夹的监听,那么一一进行用户回调函数的触发
func (fw *fWatcher) watcherTrigger(changedPath string) {

}

// Status 获取当前程序运行状态
func (fw *fWatcher) Status() int {
	return fw.status
}

// AddFile 添加文件监听
func (fw *fWatcher) AddFile(filePath, callbackName string, callback func(string) error) {
	if _, ok := fw.fileEvents[filePath]; !ok {
		fw.fileEvents[filePath] = make(map[string]func(string) error)
	}
	fw.fileEvents[filePath][callbackName] = callback
}

// AddFloder 添加文件夹变化
func (fw *fWatcher) AddFolder(folderPath, callbackName string, callback func(string) error) {
	if _, ok := fw.fileEvents[folderPath]; !ok {
		fw.fileEvents[folderPath] = make(map[string]func(string) error)
	}
	fw.fileEvents[folderPath][callbackName] = callback
}

// RmFile 移除对某文件的事件监听
func (fw *fWatcher) RmFile(filePath, callbackName string) error {
	var errNotFound error = errors.New("没有设置对文件:" + filePath + "为" + callbackName + "的监听")
	filePathCallbacks, ok := fw.fileEvents[filePath]
	if !ok {
		return errNotFound
	}
	if _, ok := filePathCallbacks[callbackName]; !ok {
		return errNotFound
	}

	delete(fw.fileEvents[filePath], callbackName)
	return nil
}

// RmFolder 移除对某文件夹的事件监听
func (fw *fWatcher) RmFolder(folderPath, callbackName string) error {
	var errNotFound error = errors.New("没有设置对文件:" + folderPath + "为" + callbackName + "的监听")
	folderPathCallbacks, ok := fw.folderEvents[folderPath]
	if !ok {
		return errNotFound
	}
	if _, ok := folderPathCallbacks[callbackName]; !ok {
		return errNotFound
	}

	delete(fw.folderEvents[folderPath], callbackName)
	return nil
}
