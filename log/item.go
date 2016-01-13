package log

import (
	"time"
)

// LogItem 日志项
type LogItem struct {
	ID      uint64    `json:",omitempty"` // ID 唯一标识
	Time    time.Time `json:",omitempty"` // Time 日志发生时间
	Level   LogLevel  `json:",omitempty"` // Level 级别
	Tag     LogTag    `json:",omitempty"` // Tag 标签
	Message string    `json:",omitempty"` // Message 日志明细
	File    LogFile   `json:",omitempty"` // File 发生日志的文件
}

// LogTag 日志标签
type LogTag string

// LogFile 发生日志的文件
type LogFile struct {
	Name      string `json:",omitempty"` // Name 文件名
	ShortName string `json:",omitempty"` // ShortName 短文件名
	FuncName  string `json:",omitempty"` // FuncName 函数名
	Line      int    `json:",omitempty"` // Line 文件行
}
