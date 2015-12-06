package log

// LogConfig 提供日志配置信息
type LogConfig struct {
	// Console 控制台输出
	Console ConsoleConfig `json:"console" yaml:"console"`
	// Global 全局配置
	Global GlobalConfig `json:"global" yaml:"global"`
	// Tags 标签配置
	Tags []TagConfig `json:"tags" yaml:"tags"`
	// Level 日志级别配置
	Levels []LevelConfig `json:"levels" yaml:"levels"`
	// Store 存储配置
	Store StoreConfig `json:"store" yaml:"store"`
}

// ConsoleConfig 控制台输出配置
type ConsoleConfig struct {
	// Level 输出级别
	Level LogLevel `json:"level" yaml:"level"`
	// Item 日志项配置
	Item LogItemConfig `json:"item" yaml:"item"`
}

// LogItemConfig 日志项配置
type LogItemConfig struct {
	// Tmpl 日志项模板
	// 模板字段说明：
	// ID 唯一标识
	// Time 日志发生时间
	// Level 级别
	// Tag 标签
	// Message 日志明细
	// FileName 文件名
	// FileFuncName 函数名
	// FileLine 文件行
	Tmpl string `json:"tmpl" yaml:"tmpl"`
	// TimeTmpl 时间模板
	// 模板字段说明：
	// Year 年份
	// Month 月份
	// Day 天数
	// Hour 小时
	// Minute 分钟
	// Second 秒
	// MilliSecond 毫秒
	TimeTmpl string `json:"time" yaml:"time"`
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	// IsPrint 是否控制台打印
	IsPrint int `yaml:"print" json:"print"`
	// Rule 日志输出规则
	// 参数说明：
	// 0表示所有配置输出
	// 1表示指定Global配置输出
	// 2表示指定Tag配置输出
	// 3表示指定Level配置输出
	// 默认值为0
	Rule LogRule `json:"rule" yaml:"rule" `
	// ShowFile 是否输出日志文件信息，包括：文件名、行数、函数名
	// 0表示不输出
	// 1表示输出
	ShowFile int `json:"showfile" yaml:"showfile"`
	// FileCaller 文件信息调用层级
	// 默认为4(当前调用)
	FileCaller int `json:"caller" yaml:"caller"`
	// Interval 读取缓冲区时间间隔（以秒为单位）
	// 默认为2秒
	Interval int `json:"interval" yaml:"interval"`
	// Buffer 缓冲区配置
	Buffer BufferConfig `json:"buffer" yaml:"buffer"`
	// TargetStore 目标存储
	TargetStore string `json:"target" yaml:"target" `
}

// BufferConfig 缓冲区配置
type BufferConfig struct {
	// Engine 存储引擎（{1:memory,2:redis}）
	// 默认为内存存储
	Engine LogBufferEngine `json:"engine" yaml:"engine"`
	// TargetStore 指向存储的名称
	TargetStore string `json:"target" yaml:"target"`
}

// CustomConfig 定制配置
type CustomConfig struct {
	// IsPrint 是否控制台打印
	IsPrint int `json:"print" yaml:"print"`
	// TargetStore 目标存储
	TargetStore string `json:"target" yaml:"target"`
}

// TagConfig 标签配置
type TagConfig struct {
	// Name 标签名
	Names []LogTag `json:"names" yaml:"names"`
	// Level 日志级别
	Level LogLevel `json:"level" yaml:"level"`
	// Config 配置
	Config CustomConfig `json:"config" yaml:"config"`
}

// LevelConfig 日志级别配置
type LevelConfig struct {
	// Values 日志级别
	Values []LogLevel `json:"values" yaml:"values"`
	// Config 配置
	Config CustomConfig `json:"config" yaml:"config"`
}

// StoreConfig 存储配置
type StoreConfig struct {
	// Redis redis存储配置
	Redis map[string]RedisConfig `json:"redis" yaml:"redis"`
	// File 文件存储配置
	File map[string]FileConfig `json:"file" yaml:"file"`
}

// RedisConfig redis配置
type RedisConfig struct {
	// Default is tcp.
	Network string `json:"network" yaml:"network"`
	// host:port address.
	Addr string `json:"addr" yaml:"addr"`
	// A database to be selected after connecting to server.
	DB int64 `json:"db" yaml:"db"`
	// The maximum number of socket connections.
	// Default is 10 connections.
	PoolSize int `json:"poolsize" yaml:"poolsize"`
}

// FileStoreConfig 文件存储
type FileConfig struct {
	// FilePath 文件存储路径,
	// 默认值为logs
	FilePath string `json:"filepath" yaml:"filepath"`
	// FileNameTmpl 文件名格式模板
	// 模板字段说明：
	// Year 年份
	// Month 月份
	// Day 天数
	// Level 日志级别
	// Tag 标签
	// 默认值为{{.Year}}{{.Month}}{{.Day}}.log
	FileNameTmpl string `json:"filename" yaml:"filename"`
	// FileSize 单个文件大小（单位KB）,
	// 默认值为512KB
	FileSize int64 `json:"filesize" yaml:"filesize"`
	// Item 日志项配置
	// 默认值：
	// time:{{.Year}}-{{.Month}}-{{.Day}} {{.Hour}}:{{.Minute}}:{{.Second}}.{{.MilliSecond}}
	// tmpl:{{.ID}} {{.Time}} {{.Level}} {{.Tag}} "{{.FileName}} {{.FileFuncName}} {{.FileLine}}" {{.Message}}
	Item LogItemConfig `json:"item" yaml:"item"`
}
