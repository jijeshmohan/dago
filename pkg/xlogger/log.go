package xlogger

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger interface {
	Debug(str string, args ...interface{})
	Info(str string, args ...interface{})
	Warn(str string, args ...interface{})
	Error(str string, args ...interface{})
}

var levelPrefixMap = map[LogLevel]string{
	DEBUG: "DEBUG : ",
	INFO:  "INFO  : ",
	WARN:  "WARN  : ",
	ERROR: "ERROR : ",
}

var emptyPrefixMap = map[LogLevel]string{
	DEBUG: "",
	INFO:  "",
	WARN:  "",
	ERROR: "",
}
