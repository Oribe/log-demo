package util

import (
	"time"
)

// GetCurrent 获取当前时间
func GetCurrent() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}
