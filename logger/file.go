package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	// Size 使用大小来分割
	Size = iota
	// Time 使用时间来分割
	Time
)

type writer struct {
	filename  string // 日志文件名
	file      *os.File
	size      int64 // 日志大小，超过设定的大小会进行分割
	splitType int   // 日志分割方式：1.大小 2.时间
}

// 创建文件
func (w *writer) createFile(filename string) {
	if filename == "" {
		panic("filename con not be empty")
	}
	if w.file != nil {
		w.file.Close()
	}
	var err error
	w.file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
}

// 写入文件
func (w *writer) write(message string) {
	fmt.Println(w.filename)
	if w.file == nil {
		w.createFile(w.filename)
	} else {
		if w.isExceedLogSize() {
			// 超过设置大小，切割
			w.splitLog()
			w.createFile(w.filename)
		}
	}
	fmt.Fprint(w.file, message)
}

// 检查日志大小
func (w *writer) isExceedLogSize() bool {
	stat, error := w.file.Stat()
	if error != nil {
		panic(error)
	}
	size := stat.Size()
	if size >= w.size { // 如果超过设置的大小
		return true
	}
	return false
}

// 日志切割
func (w *writer) splitLog() {
	currentTime := time.Now().Format("20060102150405")
	newFilenameSplice := strings.Split(w.filename, ".log")
	newFilenameSplice = append(newFilenameSplice, currentTime, ".log")
	var newFilename string
	for _, v := range newFilenameSplice {
		newFilename += v
	}
	error := os.Rename(w.filename, newFilename)
	if error != nil {
		fmt.Println(error)
	}
	w.file.Close()
}

func newWriter(filename string, size int64, spliceType int) *writer {
	w := &writer{}
	w.filename = filename
	w.size = size
	w.splitType = spliceType
	return w
}
