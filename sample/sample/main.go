package main

import (
	"time"

	"gopkg.in/alog.v1"
)

func main() {
	alog.RegisterAlog()
	alog.SetLogTag("Sample")
	alog.Debug("Debug info...")
	alog.DebugC("Debug console info...")
	alog.Info("Info info...")
	alog.InfoC("Info console info...")
	alog.Warn("Warn info...")
	alog.WarnC("Warn console info...")
	alog.Error("Error info...")
	alog.ErrorC("Error console info...")
	alog.Fatal("Fatal info...")
	alog.FatalC("Fatal console info...")
	time.Sleep(2 * time.Second)
}
