package alog

import "github.com/antlinker/alog/log"

var (
	// 提供全局的ALog
	GALog *ALog
)

// RegisterAlog 注册并初始化ALog
// configs 配置信息:
// 配置文件方式，包含yaml,json两种方式
func RegisterAlog(configs ...string) {
	GALog = NewALog(configs...)
	GALog.SetFileCaller(GALog.GetConfig().Global.FileCaller)
}

// SetLogTag 设置日志标签
func SetLogTag(tag string) {
	GALog.SetLogTag(tag)
}

// SetEnabled 设置是否启用日志
func SetEnabled(enabled bool) {
	GALog.SetEnabled(enabled)
}

// Debug Debug 消息
func Debug(v ...interface{}) {
	GALog.Write(false, log.DEBUG, "", v...)
}

// Debug Debug 格式化消息
func Debugf(format string, v ...interface{}) {
	GALog.Writef(false, log.DEBUG, "", format, v...)
}

// Debug Debug 标签消息
func DebugT(tag string, v ...interface{}) {
	GALog.Write(false, log.DEBUG, tag, v...)
}

// Debug Debug 标签格式化消息
func DebugTf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.DEBUG, tag, format, v...)
}

// Debug Debug 控制台消息
func DebugC(v ...interface{}) {
	GALog.Write(true, log.DEBUG, "", v...)
}

// Debug Debug 控制台格式化消息
func DebugCf(format string, v ...interface{}) {
	GALog.Writef(true, log.DEBUG, "", format, v...)
}

// Debug Debug 控制台标签消息
func DebugTC(tag string, v ...interface{}) {
	GALog.Write(true, log.DEBUG, tag, v...)
}

// Debug Debug 控制台标签格式化消息
func DebugTCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.DEBUG, tag, format, v...)
}

// Info Info 消息
func Info(v ...interface{}) {
	GALog.Write(false, log.INFO, "", v...)
}

// Info Info 格式化消息
func Infof(format string, v ...interface{}) {
	GALog.Writef(false, log.INFO, "", format, v...)
}

// Info Info 标签消息
func InfoT(tag string, v ...interface{}) {
	GALog.Write(false, log.INFO, tag, v...)
}

// Info Info 标签格式化消息
func InfoTf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.INFO, tag, format, v...)
}

// Info Info 控制台消息
func InfoC(v ...interface{}) {
	GALog.Write(true, log.INFO, "", v...)
}

// Info Info 控制台格式化消息
func InfoCf(format string, v ...interface{}) {
	GALog.Writef(true, log.INFO, "", format, v...)
}

// Info Info 控制台标签消息
func InfoTC(tag string, v ...interface{}) {
	GALog.Write(true, log.INFO, tag, v...)
}

// Info Info 控制台标签格式化消息
func InfoTCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.INFO, tag, format, v...)
}

// Warn Warn 消息
func Warn(v ...interface{}) {
	GALog.Write(false, log.WARN, "", v...)
}

// Warn Warn 格式化消息
func Warnf(format string, v ...interface{}) {
	GALog.Writef(false, log.WARN, "", format, v...)
}

// Warn Warn 标签消息
func WarnT(tag string, v ...interface{}) {
	GALog.Write(false, log.WARN, tag, v...)
}

// Warn Warn 标签格式化消息
func WarnTf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.WARN, tag, format, v...)
}

// Warn Warn 控制台消息
func WarnC(v ...interface{}) {
	GALog.Write(true, log.WARN, "", v...)
}

// Warn Warn 控制台格式化消息
func WarnCf(format string, v ...interface{}) {
	GALog.Writef(true, log.WARN, "", format, v...)
}

// Warn Warn 控制台标签消息
func WarnTC(tag string, v ...interface{}) {
	GALog.Write(true, log.WARN, tag, v...)
}

// Warn Warn 控制台标签格式化消息
func WarnTCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.WARN, tag, format, v...)
}

// Error Error 消息
func Error(v ...interface{}) {
	GALog.Write(false, log.ERROR, "", v...)
}

// Error Error 格式化消息
func Errorf(format string, v ...interface{}) {
	GALog.Writef(false, log.ERROR, "", format, v...)
}

// Error Error 标签消息
func ErrorT(tag string, v ...interface{}) {
	GALog.Write(false, log.ERROR, tag, v...)
}

// Error Error 标签格式化消息
func ErrorTf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.ERROR, tag, format, v...)
}

// Error Error 控制台消息
func ErrorC(v ...interface{}) {
	GALog.Write(true, log.ERROR, "", v...)
}

// Error Error 控制台格式化消息
func ErrorCf(format string, v ...interface{}) {
	GALog.Writef(true, log.ERROR, "", format, v...)
}

// Error Error 控制台标签消息
func ErrorTC(tag string, v ...interface{}) {
	GALog.Write(true, log.ERROR, tag, v...)
}

// Error Error 控制台标签格式化消息
func ErrorTCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.ERROR, tag, format, v...)
}

// Fatal Fatal 消息
func Fatal(v ...interface{}) {
	GALog.Write(false, log.FATAL, "", v...)
}

// Fatal Fatal 格式化消息
func Fatalf(format string, v ...interface{}) {
	GALog.Writef(false, log.FATAL, "", format, v...)
}

// Fatal Fatal 标签消息
func FatalT(tag string, v ...interface{}) {
	GALog.Write(false, log.FATAL, tag, v...)
}

// Fatal Fatal 标签格式化消息
func FatalTf(tag string, format string, v ...interface{}) {
	GALog.Writef(false, log.FATAL, tag, format, v...)
}

// Fatal Fatal 控制台消息
func FatalC(v ...interface{}) {
	GALog.Write(true, log.FATAL, "", v...)
}

// Fatal Fatal 控制台格式化消息
func FatalCf(format string, v ...interface{}) {
	GALog.Writef(true, log.FATAL, "", format, v...)
}

// Fatal Fatal 控制台标签消息
func FatalTC(tag string, v ...interface{}) {
	GALog.Write(true, log.FATAL, tag, v...)
}

// Fatal Fatal 控制台标签格式化消息
func FatalTCf(tag string, format string, v ...interface{}) {
	GALog.Writef(true, log.FATAL, tag, format, v...)
}
