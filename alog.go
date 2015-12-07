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
	// 提供全局的ALog
	GALog *ALog
)

// ALog 提供ALog日志模块的输出管理
type ALog struct {
	tag    log.LogTag
	config *log.LogConfig
	manage log.LogManage
}

// RegisterAlog 注册并初始化ALog
// configs 配置信息:
// 配置文件方式，包含yaml,json两种方式
// 动态配置LogConfig
func RegisterAlog(configs ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("===> [ALog]Initialization error:", err)
			os.Exit(-1)
		}
	}()
	cfg := new(log.LogConfig)
	if len(configs) > 0 {
		config := configs[0]
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
	GALog = &ALog{
		manage: manage.NewLogManage(cfg),
		config: cfg,
		tag:    log.DefaultTag,
	}
}

// SetLogTag 设置LogTag
func (a *ALog) SetLogTag(tag string) {
	a.tag = log.LogTag(tag)
}

// GetConfig 获取配置文件信息
func (a *ALog) GetConfig() *log.LogConfig {
	return a.config
}

// GetWriteNum 获取写入日志条数
func (a *ALog) GetWriteNum() int64 {
	return a.manage.TotalNum()
}

// Write 输出消息
func (a *ALog) Write(onlyConsole bool, level log.LogLevel, tag string, v ...interface{}) {
	t := log.LogTag(tag)
	if t == "" {
		t = a.tag
	}
	if onlyConsole {
		a.manage.Console(level, t, v...)
		return
	}
	a.manage.Write(level, t, v...)
}

// Writef 输出格式化消息
func (a *ALog) Writef(onlyConsole bool, level log.LogLevel, tag string, format string, v ...interface{}) {
	t := log.LogTag(tag)
	if t == "" {
		t = a.tag
	}
	if onlyConsole {
		a.manage.Consolef(level, t, format, v...)
		return
	}
	a.manage.Writef(level, t, format, v...)
}

// Debug Debug消息
func Debug(tag string, v ...interface{}) {
	GALog.Write(false, log.DEBUG, tag, v...)
}

// Debugf 格式化Debug消息
func Debugf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.DEBUG, tag, format, v...)
}

// Debug Debug控制台消息(只输出到控制台，不写入存储)
func DebugC(tag string, v ...interface{}) {
	GALog.Write(true, log.DEBUG, tag, v...)
}

// Debugf 格式化Debug控制台消息(只输出到控制台，不写入存储)
func DebugCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.DEBUG, tag, format, v...)
}

// Info Info消息
func Info(tag string, v ...interface{}) {
	GALog.Write(false, log.INFO, tag, v...)
}

// Infof 格式化Info消息
func Infof(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.INFO, tag, format, v...)
}

// Info Info控制台消息(只输出到控制台，不写入存储)
func InfoC(tag string, v ...interface{}) {
	GALog.Write(true, log.INFO, tag, v...)
}

// Infof 格式化Info控制台消息(只输出到控制台，不写入存储)
func InfoCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.INFO, tag, format, v...)
}

// Warn Warn消息
func Warn(tag string, v ...interface{}) {
	GALog.Write(false, log.WARN, tag, v...)
}

// Warnf 格式化Warn消息
func Warnf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.WARN, tag, format, v...)
}

// Warn Warn控制台消息(只输出到控制台，不写入存储)
func WarnC(tag string, v ...interface{}) {
	GALog.Write(true, log.WARN, tag, v...)
}

// Warnf 格式化Warn控制台消息(只输出到控制台，不写入存储)
func WarnCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.WARN, tag, format, v...)
}

// Error Error消息
func Error(tag string, v ...interface{}) {
	GALog.Write(false, log.ERROR, tag, v...)
}

// Errorf 格式化Error消息
func Errorf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.ERROR, tag, format, v...)
}

// Error Error控制台消息(只输出到控制台，不写入存储)
func ErrorC(tag string, v ...interface{}) {
	GALog.Write(true, log.ERROR, tag, v...)
}

// Errorf 格式化Error控制台消息(只输出到控制台，不写入存储)
func ErrorCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.ERROR, tag, format, v...)
}

// Fatal Fatal消息
func Fatal(tag string, v ...interface{}) {
	GALog.Write(false, log.FATAL, tag, v...)
}

// Fatalf 格式化Fatal消息
func Fatalf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.FATAL, tag, format, v...)
}

// Fatal Fatal控制台消息(只输出到控制台，不写入存储)
func FatalC(tag string, v ...interface{}) {
	GALog.Write(true, log.FATAL, tag, v...)
}

// Fatalf 格式化Fatal控制台消息(只输出到控制台，不写入存储)
func FatalCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.FATAL, tag, format, v...)
}
