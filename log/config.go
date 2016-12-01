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
	// ShortName 短文件名
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
	// IsEnabled 是否启用日志
	// 参数说明：
	// 1表示启用
	// 2表示不启用
	// 默认值为1
	IsEnabled int `yaml:"enabled" json:"enabled"`
	// IsPrint 是否打印到控制台
	// 1表示打印
	// 2表示不打印
	IsPrint int `yaml:"print" json:"print"`
	// Rule 日志输出规则
	// 参数说明：
	// 0表示所有配置输出
	// 1表示指定Global配置输出
	// 2表示指定Tag配置输出
	// 3表示指定Level配置输出
	// 默认值为0
	Rule LogRule `json:"rule" yaml:"rule"`
	// Level 日志级别
	Level LogLevel `json:"level" yaml:"level"`
	// ShowFile 是否输出日志文件信息，包括：文件名、行数、函数名
	// 1表示输出
	// 2表示不输出
	ShowFile int `json:"showfile" yaml:"showfile"`
	// FileCaller 文件信息调用层级
	// 默认为5(当前调用)
	FileCaller int `json:"caller" yaml:"caller"`
	// Interval 读取缓冲区时间间隔（以秒为单位）
	// 默认为1秒
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
	// Elastic ElasticSearch存储配置
	Elastic map[string]ElasticConfig `json:"elastic" yaml:"elastic"`
	// Mongo MongoDB存储配置
	Mongo map[string]MongoConfig `json:"mongo" yaml:"mongo"`
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

	// 文件保留天数,
	// 默认值为0(保留全部文件)
	RetainDay int `json:"retian" yaml:"retain"`

	// 清理文件周期(单位分钟)
	// 默认为720分钟
	GCInterval int `json:"interval" yaml:"interval"`
}

// ElasticConfig ElasticSearch持久化存储
type ElasticConfig struct {
	// URL 指定ElasticSearch的请求节点
	// 默认值为http://127.0.0.1:9200
	URL string `json:"url" yaml:"url"`
	// IndexTmpl 索引模板
	// 模板字段说明：
	// Year 年份
	// Month 月份
	// Day 天数
	// Level 日志级别
	// Tag 标签
	// 默认值为{{.Year}}.{{.Month}}.{{.Day}}
	IndexTmpl string `json:"index" yaml:"index"`
	// TypeTmpl 文档类型模板
	// 模板字段说明：
	// Year 年份
	// Month 月份
	// Day 天数
	// Level 日志级别
	// Tag 标签
	// 默认值为ALogs
	TypeTmpl string `json:"type" yaml:"type"`
}

// MongoConfig 提供MongoDB持久化存储
type MongoConfig struct {
	// URL 指定MongoDB的链接地址
	// 默认值为mongodb://127.0.0.1:27017
	URL string `json:"url" yaml:"url"`
	// DBTmpl 存储数据库名称模板
	// 模板字段说明：
	// Year 年份
	// Month 月份
	// Day 天数
	// Level 日志级别
	// Tag 标签
	// 默认值为alog
	DBTmpl string `json:"db" yaml:"db"`
	// CollectionTmpl 存储集合名称模板
	// 模板字段说明：
	// Year 年份
	// Month 月份
	// Day 天数
	// Level 日志级别
	// Tag 标签
	// 默认值为{{.Year}}{{.Month}}{{.Day}}
	CollectionTmpl string `json:"collection" yaml:"collection"`
}
