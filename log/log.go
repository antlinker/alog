package log

// LogBuffer 提供日志缓冲区操作接口
type LogBuffer interface {
	// Push 将日志项追加写入到缓冲区
	Push(LogItem) error
	// Pop 弹出缓冲区的第一个元素
	Pop() (*LogItem, error)
}

// LogStore 提供日志持久化存储接口
type LogStore interface {
	// Store 将日志项写入到存储区
	Store(*LogItem) error
}

// LogManage 提供日志的写入、存储及控制台输出接口
type LogManage interface {
	// Write 写入日志信息
	Write(level LogLevel, tag LogTag, v ...interface{})
	// Writef 写入格式化日志信息
	Writef(level LogLevel, tag LogTag, format string, v ...interface{})
	// Console 将日志输出到控制台（不写入文件）
	Console(level LogLevel, tag LogTag, v ...interface{})
	// Consolef 将格式化日志输出到控制台（不写入文件）
	Consolef(level LogLevel, tag LogTag, format string, v ...interface{})
	// TotalNum 写入日志总条数
	TotalNum() int64
}
