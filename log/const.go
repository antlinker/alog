package log

const (
	// DefaultEnabled 默认启用日志
	DefaultEnabled = 1

	// DefaultTag 默认日志标签
	DefaultTag LogTag = "ALOG"

	// DefaultSystemTag 默认系统日志消息标签
	DefaultSystemTag LogTag = "SYSTEM"

	// DefaultGlobalKey 默认Global存储键
	// DefaultGlobalKey = "global"

	// DefaultFilePath 默认日志文件存储路径
	DefaultFilePath = "logs"

	// DefaultFileSize 默认单个日志文件大小为2048KB
	DefaultFileSize = 2048

	// DefaultPrint 默认输出日志到控制台
	DefaultPrint = 1

	// DefaultLogLevel 默认日志级别
	DefaultLogLevel = DEBUG

	// DefaultInterval 默认写入时间间隔
	DefaultInterval = 1

	// DefaultShowFile 默认输出文件信息
	DefaultShowFile = 1

	// DefaultFileCaller 默认输出文件信息调用层级
	DefaultFileCaller = 6

	// DefaultElasticURL ElasticSearch默认节点
	DefaultElasticURL = "http://127.0.0.1:9200"

	// DefaultMongoURL 指定MongoDB的默认链接地址
	DefaultMongoURL = "mongodb://127.0.0.1:27017"
)

// LogRule 日志规则
type LogRule byte

const (
	// AlwaysRule 检查全部输出规则
	AlwaysRule LogRule = iota
	// GlobalRule 按照Global输出
	GlobalRule
	// TagRule 按照TagRule输出
	TagRule
	// LevelRule 按照LevelRule输出
	LevelRule
)

// LogBufferEngine 日志缓冲区引擎
type LogBufferEngine byte

const (
	// MEMORY_BUFFER 内存缓冲区
	MEMORY_BUFFER LogBufferEngine = iota + 1
	// REDIS_BUFFER redis缓冲区
	REDIS_BUFFER
)

// LogStoreEngine 日志存储引擎
type LogStoreEngine byte

const (
	// FILE_STORE 文件存储
	FILE_STORE LogStoreEngine = iota + 1
)
