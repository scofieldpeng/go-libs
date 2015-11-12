//filelog 文件日志操作包 Copyright 2015 Author Scofield

/**
 * 使用方法：
 *  //设置log
 *	if err := filelog.FileLog.Load(); err != nil {
 * 	    fmt.Println(err)
 *      os.Exit(1)
 *  }
 *	//后台协程自动更新log的文件名称
 *	go filelog.FileLog.AutoUpdate()
 *
 *  其他地方使用log包内的函数即可（如log.Println(),log.Printf()...)
 */
package filelog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

//FileLogger 结构体
type FileLogger struct {
	logPath     string   //log存放的绝对路径
	logFile     string   //log的文件名称
	fileHandler *os.File //log文件句柄
}

// 新建一个NewFileLogger
func NewFileLogger(confpath ...string) *FileLogger {
	if len(confpath) == 0 {
		tmp, _ := filepath.Abs(os.Args[0])
		baseDir := filepath.Dir(tmp)
		confpath = []string{baseDir + string(os.PathSeparator) + "log" + string(os.PathSeparator)}
	}
	return &FileLogger{
		logPath: confpath[0],
	}
}

//实现Write方法
func (this *FileLogger) Write(str []byte) (int, error) {
	if num, err := os.Stdout.Write(str); err != nil {
		return num, err
	}

	if num, err := this.fileHandler.Write(str); err != nil {
		return num, err
	} else {
		return num, nil
	}

}

//设置log文件句柄
func (this *FileLogger) SetFileHandler() error {
	confFile, err := os.OpenFile(this.LogFileName(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0655))
	if err != nil {
		return err
	}

	this.fileHandler = confFile
	return nil
}

//Load 载入日志
func (this *FileLogger) Load() error {
	this.SetLogFile()

	if err := this.SetFileHandler(); err != nil {
		return err
	}

	log.SetOutput(this)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetPrefix("[front]")

	return nil
}

//SetLogFile 设置日志文件名称
func (this *FileLogger) SetLogFile() {
	this.logFile = time.Now().Format("2006-01-02") + ".log"
}

//LogFileName 获取日志文件名称
func (this *FileLogger) LogFileName() string {
	return this.logPath + this.logFile
}

//AutoRun 设置自动更新log文件名称
func (this *FileLogger) AutoUpdate() {
	//设置log文件夹
	if err := os.MkdirAll(this.logPath, os.FileMode(0755)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	this.SetLogFile()
	if err := this.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	now := time.Now().Unix()
	tommorrow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local).Unix()
	time.Sleep(time.Duration(tommorrow-now) * time.Second)
	for {
		this.SetLogFile()
		if err := this.Load(); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 86400)
	}
}

var FileLog = NewFileLogger()
