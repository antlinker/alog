package main

import (
	"time"

	"gopkg.in/alog.v1"
)

func main() {
	alog.RegisterAlog("config.yaml")
	alog.Debug("Debug", "Debug info...")
	alog.DebugC("Debug", "Debug console info...")
	alog.Info("Info", "Info info...")
	alog.InfoC("Info", "Info console info...")
	alog.Warn("Warn", "Warn info...")
	alog.WarnC("Warn", "Warn console info...")
	alog.Error("Error", "Error info...")
	alog.ErrorC("Error", "Error console info...")
	alog.Fatal("Fatal", "Fatal info...")
	alog.FatalC("Fatal", "Fatal console info...")
	time.Sleep(2 * time.Second)
}
