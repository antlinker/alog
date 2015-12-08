package alog

import (
	"fmt"
	"os"

	"gopkg.in/alog.v1/log"
	"gopkg.in/alog.v1/manage"
)

// ALog 提供ALog日志模块的输出管理
type ALog struct {
	tag    log.LogTag
	config *log.LogConfig
	manage log.LogManage
}

// NewALog 获取ALog实例
func NewALog(configs ...interface{}) *ALog {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("===> [ALog]Initialization error:", err)
			os.Exit(-1)
		}
	}()
	var config *log.LogConfig
	if len(configs) > 0 {
		cfg, err := parseConfig(configs[0])
		if err != nil {
			panic(err)
		}
		config = cfg
	} else if _GConfig != nil {
		cfg := *_GConfig
		config = &cfg
	} else {
		config = defaultConfig()
	}
	alg := &ALog{
		config: config,
		manage: manage.NewLogManage(config),
		tag:    log.DefaultTag,
	}
	return alg
}

// SetLogTag 设置LogTag
func (a *ALog) SetLogTag(tag string) {
	a.tag = log.LogTag(tag)
}

// SetFileCaller 设置文件调用层次
func (a *ALog) SetFileCaller(caller int) {
	(*(a.config)).Global.FileCaller = caller
}

// SetRule 设置输出规则
func (a *ALog) SetRule(rule log.LogRule) {
	(*(a.config)).Global.Rule = rule
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

func (a *ALog) Debug(v ...interface{}) {
	a.Write(false, log.DEBUG, "", v...)
}

func (a *ALog) Debugf(format string, v ...interface{}) {
	a.Writef(false, log.DEBUG, "", format, v...)
}

func (a *ALog) DebugT(tag string, v ...interface{}) {
	a.Write(false, log.DEBUG, tag, v...)
}

func (a *ALog) DebugTf(tag string, format string, v ...interface{}) {
	a.Writef(false, log.DEBUG, tag, format, v...)
}

func (a *ALog) DebugC(v ...interface{}) {
	a.Write(true, log.DEBUG, "", v...)
}

func (a *ALog) DebugCf(format string, v ...interface{}) {
	a.Writef(true, log.DEBUG, "", format, v...)
}

func (a *ALog) DebugTC(tag string, v ...interface{}) {
	a.Write(true, log.DEBUG, tag, v...)
}

func (a *ALog) DebugTCf(tag string, format string, v ...interface{}) {
	a.Writef(true, log.DEBUG, tag, format, v...)
}

func (a *ALog) Info(v ...interface{}) {
	a.Write(false, log.INFO, "", v...)
}

func (a *ALog) Infof(format string, v ...interface{}) {
	a.Writef(false, log.INFO, "", format, v...)
}

func (a *ALog) InfoT(tag string, v ...interface{}) {
	a.Write(false, log.INFO, tag, v...)
}

func (a *ALog) InfoTf(tag string, format string, v ...interface{}) {
	a.Writef(false, log.INFO, tag, format, v...)
}

func (a *ALog) InfoC(v ...interface{}) {
	a.Write(true, log.INFO, "", v...)
}

func (a *ALog) InfoCf(format string, v ...interface{}) {
	a.Writef(true, log.INFO, "", format, v...)
}

func (a *ALog) InfoTC(tag string, v ...interface{}) {
	a.Write(true, log.INFO, tag, v...)
}

func (a *ALog) InfoTCf(tag string, format string, v ...interface{}) {
	a.Writef(true, log.INFO, tag, format, v...)
}

func (a *ALog) Warn(v ...interface{}) {
	a.Write(false, log.WARN, "", v...)
}

func (a *ALog) Warnf(format string, v ...interface{}) {
	a.Writef(false, log.WARN, "", format, v...)
}

func (a *ALog) WarnT(tag string, v ...interface{}) {
	a.Write(false, log.WARN, tag, v...)
}

func (a *ALog) WarnTf(tag string, format string, v ...interface{}) {
	a.Writef(false, log.WARN, tag, format, v...)
}

func (a *ALog) WarnC(v ...interface{}) {
	a.Write(true, log.WARN, "", v...)
}

func (a *ALog) WarnCf(format string, v ...interface{}) {
	a.Writef(true, log.WARN, "", format, v...)
}

func (a *ALog) WarnTC(tag string, v ...interface{}) {
	a.Write(true, log.WARN, tag, v...)
}

func (a *ALog) WarnTCf(tag string, format string, v ...interface{}) {
	a.Writef(true, log.WARN, tag, format, v...)
}

func (a *ALog) Error(v ...interface{}) {
	a.Write(false, log.ERROR, "", v...)
}

func (a *ALog) Errorf(format string, v ...interface{}) {
	a.Writef(false, log.ERROR, "", format, v...)
}

func (a *ALog) ErrorT(tag string, v ...interface{}) {
	a.Write(false, log.ERROR, tag, v...)
}

func (a *ALog) ErrorTf(tag string, format string, v ...interface{}) {
	a.Writef(false, log.ERROR, tag, format, v...)
}

func (a *ALog) ErrorC(v ...interface{}) {
	a.Write(true, log.ERROR, "", v...)
}

func (a *ALog) ErrorCf(format string, v ...interface{}) {
	a.Writef(true, log.ERROR, "", format, v...)
}

func (a *ALog) ErrorTC(tag string, v ...interface{}) {
	a.Write(true, log.ERROR, tag, v...)
}

func (a *ALog) ErrorTCf(tag string, format string, v ...interface{}) {
	a.Writef(true, log.ERROR, tag, format, v...)
}

func (a *ALog) Fatal(v ...interface{}) {
	a.Write(false, log.FATAL, "", v...)
}

func (a *ALog) Fatalf(format string, v ...interface{}) {
	a.Writef(false, log.FATAL, "", format, v...)
}

func (a *ALog) FatalT(tag string, v ...interface{}) {
	a.Write(false, log.FATAL, tag, v...)
}

func (a *ALog) FatalTf(tag string, format string, v ...interface{}) {
	a.Writef(false, log.FATAL, tag, format, v...)
}

func (a *ALog) FatalC(v ...interface{}) {
	a.Write(true, log.FATAL, "", v...)
}

func (a *ALog) FatalCf(format string, v ...interface{}) {
	a.Writef(true, log.FATAL, "", format, v...)
}

func (a *ALog) FatalTC(tag string, v ...interface{}) {
	a.Write(true, log.FATAL, tag, v...)
}

func (a *ALog) FatalTCf(tag string, format string, v ...interface{}) {
	a.Writef(true, log.FATAL, tag, format, v...)
}
