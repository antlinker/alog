package main

import (
	"gopkg.in/alog.v1"
)

const (
	_LogNum = 10
	_LogTag = "CONSOLE"
)

func main() {
	alog.RegisterAlog(nil)
	for i := 0; i < _LogNum; i++ {
		alog.InfoC(_LogTag, "The console:", i)
	}
}
