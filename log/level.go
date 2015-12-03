package log

// LogLevel 日志级别
type LogLevel byte

const (
	DEBUG LogLevel = iota + 1
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) ToString() string {
	var v string
	switch l {
	case DEBUG:
		v = "Debug"
	case INFO:
		v = "Info"
	case WARN:
		v = "Warn"
	case ERROR:
		v = "Error"
	case FATAL:
		v = "Fatal"
	}
	return v
}
