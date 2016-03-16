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
//     filewatcher.FWatcher.AddFile("filePath",callback) /* 监听文件 */
//     filewatcher.FWatcher.AddFolder("folderPath",callback)/* 监听文件夹 */
//
// 3. 取消文件/文件夹监听
//     filewatcher.FWatcher.RmFile("filePath",callback) /* 移除对文件的监听 */
//     filewatcher.FWatcher.RmFolder("folderPath",callback) /* 移除对文件夹的监听 */
package filewatcher

import(
//	"github.com/fsnotify/fsnotify"
)

type (
    // fWatcher fWatcher结构体,用于实现文件监听的主要功能
    fWatcher struct {
        events map[string] []func(string)error
    }
)

var(
    FWatcher fWatcher
)

func init() {
    FWatcher = newFWatcher()
}

// newFWatcher 新建一个fWatcher对象结构体
func newFWatcher() fWatcher{
    return fWatcher{}
}

// Init 初始化文件监听
func (fw *fWatcher) Init() error {
    return nil
}

// AddFile 添加文件监听
func (fw *fWatcher) AddFile(filePath string,callback func(string)error) error {
    return nil
}

// AddFloder 添加文件夹变化
func (fw *fWatcher) AddFolder(folderPath string,callback func(string)error) error {
    return nil
}

// RmFile 移除对某文件的事件监听
func (fw *fWatcher) RmFile(filePath string,callback func(string)error) error {
    return nil
}

// RmFolder 移除对某文件夹的事件监听
func (fw *fWatcher) RmFolder(folderPath string,FCallback func(string)error) error {
    return nil
}