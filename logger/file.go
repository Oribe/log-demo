package logger

import (
	"fmt"
	"os"
)

type writer struct {
	filename string
	file     *os.File
}

func (w *writer) createFile(filename string) {
	if filename == "" {
		panic("filename con not be empty")
	}
	if w.file != nil {
		w.file.Close()
	}
	var err error
	w.file, err = os.Create(filename)
	if err != nil {
		panic("create file error")
	}
}

func (w *writer) write(message string) {
	fmt.Println(w.filename)
	if w.file == nil {
		w.createFile(w.filename)
	}
	fmt.Fprint(w.file, message)
}

func newWriter(filename string) *writer {
	w := &writer{}
	w.filename = filename
	return w
}
