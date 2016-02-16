package alog

import "gopkg.in/alog.v1/log"

// 加载默认配置
func loadDefaultConfig() *log.LogConfig {
	config := new(log.LogConfig)
	config.Console.Level = log.DefaultLogLevel
	config.Console.Item.Tmpl = log.DefaultConsoleTmpl
	config.Console.Item.TimeTmpl = log.DefaultConsoleTimeTmpl
	config.Global.IsEnabled = log.DefaultEnabled
	config.Global.IsPrint = log.DefaultPrint
	config.Global.Level = log.DefaultLogLevel
	config.Global.Buffer.Engine = log.MEMORY_BUFFER
	config.Global.Interval = log.DefaultInterval
	config.Global.ShowFile = log.DefaultShowFile
	config.Global.FileCaller = log.DefaultFileCaller
	// config.Global.TargetStore = log.DefaultGlobalKey
	// config.Store.File = map[string]log.FileConfig{
	// 	log.DefaultGlobalKey: log.FileConfig{},
	// }
	return config
}
