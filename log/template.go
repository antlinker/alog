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
	DefaultConsoleTimeTmpl = `{{.Hour}}:{{.Minute}}`

	// DefaultConsoleTmpl 默认控制台输出模板
	DefaultConsoleTmpl = `[{{.ID}} {{.Time}} {{.Level}} {{.Tag}}] {{.Message}}`
	// DefaultSystemTmpl 默认系统控制台输出模板
	DefaultSystemTmpl = `[{{.Time}} {{.Level}} {{.Tag}}] {{.Message}}`
	// DefaultMsgTmpl 默认文件存储日志模板
	DefaultMsgTmpl = `{{.ID}} {{.Time}} {{.Level}} {{.Tag}} "{{.FileName}} {{.FileFuncName}} {{.FileLine}}" {{.Message}}`

	// DefaultFileNameTmpl 默认存储日志文件名模板
	DefaultFileNameTmpl = `{{.Year}}{{.Month}}{{.Day}}.log`
)
