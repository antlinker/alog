package main

import (
	"github.com/antlinker/alog"
)

func main() {
	alog.RegisterAlog()
	alog.SetLogTag("CONSOLE")
	for i := 0; i < 10; i++ {
		alog.InfoC("The console:", i)
	}
}
