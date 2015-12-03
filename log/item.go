package log

import (
	"time"
)

// LogItem 日志项
type LogItem struct {
	// ID 唯一标识
	ID uint64
	// Time 日志发生时间
	Time time.Time
	// Level 级别
	Level LogLevel
	// Tag 标签
	Tag LogTag
	// Message 日志明细
	Message string
	// File 发生日志的文件
	File LogFile
}

// LogTag 日志标签
type LogTag string

// LogFile 发生日志的文件
type LogFile struct {
	// Name 文件名
	Name string
	// FuncName 函数名
	FuncName string
	// Line 文件行
	Line int
}
