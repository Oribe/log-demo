package logger

import (
	"fmt"
	"runtime"
)

func errorPosition() (pc uintptr, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		fmt.Println("can not location error")
	}
	return pc, file, line
}
