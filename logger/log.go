package logger

import (
	"Demo/exercise/logger/util"
	"fmt"
	"os"
	"path"
	"strings"
)

/*
	日志级别：
		1. debug		最低级别
		2. info			重要
		3. warning	可修复
		4. error		可修复
		5. Fatal		相当严重
*/

// 日志类型级别
const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger log对象
type Logger struct {
	level                     string // 日志级别
	basePath                  string // 日志保存路径
	info, warning, err, fatal *writer
}

// 获取当前时间
func getNowTime() string {
	return util.GetCurrent()
}

func (l *Logger) switchLevel(level string) int {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return 0
	}
}

// 日志可调用？
func (l *Logger) logEnable(level int) bool {
	return level >= l.switchLevel(l.level)
}

func (l *Logger) writeFile(level int, message string) {
	switch level {
	case INFO:
		l.info.write(message)
	case WARNING:
		l.warning.write(message)
		l.info.write(message)
	case ERROR:
		l.info.write(message)
		l.err.write(message)
	case FATAL:
		l.info.write(message)
		l.fatal.write(message)
	}
}

// 封裝统一打印方法
func (l *Logger) printInfo(level string, format string, a ...interface{}) {
	_, file, line := errorPosition()
	if !l.logEnable(l.switchLevel(level)) {
		return
	}
	message := fmt.Sprintf(format, a...)
	// fmt.Printf("[%s] [%s] [file:%s:%d] %v\n", getNowTime(), level, file, line, message)
	message = fmt.Sprintf("[%s] [%s] [file:%s:%d] %v\n", getNowTime(), level, file, line, message)
	fmt.Println(message)
	l.writeFile(l.switchLevel(level), message)
}

// Debug 最低级别调试用
func (l *Logger) Debug(format string, a ...interface{}) {
	// if !l.logEnable(DEBUG) {
	// 	return
	// }
	// message := fmt.Sprintf(format, a...)
	// fmt.Printf("[%s] [Debug] %v\n", getNowTime(), message)
	l.printInfo("Debug", format, a...)
}

// Info 打印所有信息
func (l *Logger) Info(format string, a ...interface{}) {
	// if !l.logEnable(INFO) {
	// 	return
	// }
	// message := fmt.Sprintf(format, a...)
	// fmt.Printf("[%s] [Info] %v\n", getNowTime(), message)
	l.printInfo("Info", format, a...)
}

// Warning 警告信息
func (l *Logger) Warning(format string, a ...interface{}) {
	// if !l.logEnable(WARNING) {
	// 	return
	// }
	// message := fmt.Sprintf(format, a...)
	// fmt.Printf("[%s] [Warning] %v\n", getNowTime(), message)
	l.printInfo("Warning", format, a...)
}

// Error 错误信息
func (l *Logger) Error(format string, a ...interface{}) {
	// if !l.logEnable(ERROR) {
	// 	return
	// }
	// message := fmt.Sprintf(format, a...)
	// fmt.Printf("[%s] [Error] %v\n", getNowTime(), message)
	l.printInfo("Error", format, a...)
}

// Fatal 严重错误信息
func (l *Logger) Fatal(format string, a ...interface{}) {
	// if !l.logEnable(FATAL) {
	// 	return
	// }
	// message := fmt.Sprintf(format, a...)
	// fmt.Printf("[%s] [Fatal] %v\n", getNowTime(), message)
	l.printInfo("Fatal", format, a...)
}

// Config 配置文件
type Config struct {
	level string
}

// NewLogger 构造函数
func NewLogger(level string, basePath string) *Logger {
	if level == "" {
		level = "Info"
	}
	if basePath == "" {
		basePath = path.Dir("./log/")
	}
	_, error := os.Stat(basePath)
	if error != nil {
		err := os.Mkdir(basePath, 0777)
		if err != nil {
			panic("dir make error")
		}
	}
	fmt.Println(basePath)
	info := newWriter(path.Join(basePath, "./info.log"))
	warning := newWriter(path.Join(basePath, "./warning.log"))
	err := newWriter(path.Join(basePath, "./error.log"))
	fatal := newWriter(path.Join(basePath, "./fatal.log"))

	return &Logger{
		level,
		basePath,
		info,
		warning,
		err,
		fatal,
	}
}
