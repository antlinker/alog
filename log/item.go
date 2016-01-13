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

// ToMap 将LogItem转换为字典值
func (lt *LogItem) ToMap() map[string]interface{} {
	data := map[string]interface{}{
		"ID":      lt.ID,
		"Time":    lt.Time,
		"Level":   lt.Level.ToString(),
		"Tag":     lt.Tag,
		"Message": lt.Message,
	}
	if v := lt.File.Name; v != "" {
		data["FileName"] = v
	}
	if v := lt.File.ShortName; v != "" {
		data["FileName"] = v
	}
	if v := lt.File.FuncName; v != "" {
		data["FileFuncName"] = v
	}
	if v := lt.File.Line; v != 0 {
		data["FileLine"] = v
	}
	return data
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
