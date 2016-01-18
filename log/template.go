package log

// TmplKey 提供模板键的标识
type TmplKey uint

const (
	// TmplConsoleTime 控制台时间模板
	TmplConsoleTime TmplKey = iota + 1
	// TmplConsole 控制台模板
	TmplConsole
)

const (
	// DefaultTimeTmpl 默认时间输出模板
	DefaultTimeTmpl = `{{.Year}}-{{.Month}}-{{.Day}} {{.Hour}}:{{.Minute}}:{{.Second}}.{{.MilliSecond}}`
	// DefaultConsoleTimeTmpl 默认控制台时间输出模板
	DefaultConsoleTimeTmpl = `{{.Hour}}:{{.Minute}}:{{.Second}}`

	// DefaultConsoleTmpl 默认控制台输出模板
	DefaultConsoleTmpl = `[{{.Time}}｜{{.Level}}｜{{.Tag}}]{{.ShortName}}:{{.FileLine}}:{{.Message}}`
	// DefaultSystemTmpl 默认系统控制台输出模板
	DefaultSystemTmpl = `[{{.Time}} {{.Level}} {{.Tag}}] {{.Message}}`
	// DefaultMsgTmpl 默认文件存储日志模板
	DefaultMsgTmpl = `{{.Time}} {{.Level}} {{.Tag}} "{{.ShortName}} {{.FileFuncName}} {{.FileLine}}" {{.Message}}`

	// DefaultFileNameTmpl 默认存储日志文件名模板
	DefaultFileNameTmpl = `{{.Year}}{{.Month}}{{.Day}}.log`

	// DefaultElasticIndexTmpl ElasticSearch文档索引名称模板
	DefaultElasticIndexTmpl = `{{.Year}}.{{.Month}}.{{.Day}}`
	// DefaultElasticTypeTmpl ElasticSearch文档类型名称模板
	DefaultElasticTypeTmpl = `ALogs`

	// DefaultMongoDBTmpl MongoDB数据库名称模板
	DefaultMongoDBTmpl = `alog`
	// DefaultMongoCollectionTmpl MongoDB集合名称模板
	DefaultMongoCollectionTmpl = `{{.Year}}{{.Month}}{{.Day}}`
)
