package main

import (
	"gopkg.in/alog.v1"
)

func main() {
	alog.RegisterAlog()
	alog.GALog.SetLogTag("CONSOLE")
	for i := 0; i < 10; i++ {
		alog.InfoC("The console:", i)
	}
}
