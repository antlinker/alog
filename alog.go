package alog

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/alog.v1/log"
	"gopkg.in/alog.v1/manage"
	"gopkg.in/alog.v1/utils"
)

var (
	// 提供全局的LogManage
	GLogManage log.LogManage
)

// RegisterAlog 注册并初始化ALog
// config 配置信息:
// 配置文件方式，包含yaml,json两种方式
// 动态配置LogConfig
func RegisterAlog(config interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("===> [ALog]Initialization error:", err)
			os.Exit(-1)
		}
	}()
	cfg := new(log.LogConfig)
	if config != nil {
		if v, ok := config.(string); ok {
			err := utils.NewConfig(v).Read(cfg)
			if err != nil {
				panic(err)
			}
		} else if v, ok := config.(log.LogConfig); ok {
			cfg = &v
		} else if v, ok := config.(*log.LogConfig); ok {
			cfg = v
		} else {
			panic(errors.New("Wrong configuration."))
		}
	}
	if cfg.Console.Item.Tmpl == "" {
		cfg.Console.Item.Tmpl = log.DefaultConsoleTmpl
	}
	if cfg.Console.Item.TimeTmpl == "" {
		cfg.Console.Item.TimeTmpl = log.DefaultConsoleTimeTmpl
	}
	if cfg.Global.Interval == 0 {
		cfg.Global.Interval = log.DefaultInterval
	}
	if cfg.Global.Buffer.Engine == 0 {
		cfg.Global.Buffer.Engine = log.MEMORY_BUFFER
	}
	if cfg.Global.TargetStore == "" {
		cfg.Global.TargetStore = log.DefaultGlobalKey
	}
	if cfg.Store.File == nil {
		cfg.Store.File = map[string]log.FileConfig{
			log.DefaultGlobalKey: log.FileConfig{},
		}
	}
	if cfg.Global.FileCaller == 0 {
		cfg.Global.FileCaller = log.DefaultFileCaller
	}

	GLogManage = manage.NewLogManage(cfg)
}

// Debug Debug消息
func Debug(tag log.LogTag, v ...interface{}) {
	GLogManage.Write(log.DEBUG, tag, v...)
}

// Debugf 格式化Debug消息
func Debugf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Writef(log.DEBUG, tag, format, v...)
}

// Debug Debug控制台消息(只输出到控制台，不写入存储)
func DebugC(tag log.LogTag, v ...interface{}) {
	GLogManage.Console(log.DEBUG, tag, v...)
}

// Debugf 格式化Debug控制台消息(只输出到控制台，不写入存储)
func DebugCf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Consolef(log.DEBUG, tag, format, v...)
}

// Info Info消息
func Info(tag log.LogTag, v ...interface{}) {
	GLogManage.Write(log.INFO, tag, v...)
}

// Infof 格式化Info消息
func Infof(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Writef(log.INFO, tag, format, v...)
}

// Info Info控制台消息(只输出到控制台，不写入存储)
func InfoC(tag log.LogTag, v ...interface{}) {
	GLogManage.Console(log.INFO, tag, v...)
}

// Infof 格式化Info控制台消息(只输出到控制台，不写入存储)
func InfoCf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Consolef(log.INFO, tag, format, v...)
}

// Warn Warn消息
func Warn(tag log.LogTag, v ...interface{}) {
	GLogManage.Write(log.WARN, tag, v...)
}

// Warnf 格式化Warn消息
func Warnf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Writef(log.WARN, tag, format, v...)
}

// Warn Warn控制台消息(只输出到控制台，不写入存储)
func WarnC(tag log.LogTag, v ...interface{}) {
	GLogManage.Console(log.WARN, tag, v...)
}

// Warnf 格式化Warn控制台消息(只输出到控制台，不写入存储)
func WarnCf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Consolef(log.WARN, tag, format, v...)
}

// Error Error消息
func Error(tag log.LogTag, v ...interface{}) {
	GLogManage.Write(log.ERROR, tag, v...)
}

// Errorf 格式化Error消息
func Errorf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Writef(log.ERROR, tag, format, v...)
}

// Error Error控制台消息(只输出到控制台，不写入存储)
func ErrorC(tag log.LogTag, v ...interface{}) {
	GLogManage.Console(log.ERROR, tag, v...)
}

// Errorf 格式化Error控制台消息(只输出到控制台，不写入存储)
func ErrorCf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Consolef(log.ERROR, tag, format, v...)
}

// Fatal Fatal消息
func Fatal(tag log.LogTag, v ...interface{}) {
	GLogManage.Write(log.FATAL, tag, v...)
}

// Fatalf 格式化Fatal消息
func Fatalf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Writef(log.FATAL, tag, format, v...)
}

// Fatal Fatal控制台消息(只输出到控制台，不写入存储)
func FatalC(tag log.LogTag, v ...interface{}) {
	GLogManage.Console(log.FATAL, tag, v...)
}

// Fatalf 格式化Fatal控制台消息(只输出到控制台，不写入存储)
func FatalCf(tag log.LogTag, format string, v ...interface{}) {
	GLogManage.Consolef(log.FATAL, tag, format, v...)
}
