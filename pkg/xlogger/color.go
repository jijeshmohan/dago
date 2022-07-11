package xlogger

import "github.com/gookit/color"

type colorLogger struct {
	level     LogLevel
	prefixMap map[LogLevel]string
}

func NewColorLogger(level LogLevel, prefix bool) Logger {
	prefixMap := emptyPrefixMap
	if prefix {
		prefixMap = levelPrefixMap
	}

	return &colorLogger{
		level:     level,
		prefixMap: prefixMap,
	}
}

func (c *colorLogger) Debug(str string, args ...interface{}) {
	if c.level <= DEBUG {
		color.Gray.Printf(c.prefixMap[DEBUG]+str+"\n", args...)
	}
}

func (c *colorLogger) Info(str string, args ...interface{}) {
	if c.level <= INFO {
		color.Info.Printf(c.prefixMap[INFO]+str+"\n", args...)
	}
}

func (c *colorLogger) Warn(str string, args ...interface{}) {
	if c.level <= WARN {
		color.Warn.Printf(c.prefixMap[WARN]+str+"\n", args...)
	}

}

func (c *colorLogger) Error(str string, args ...interface{}) {
	if c.level <= ERROR {
		color.Red.Printf(c.prefixMap[ERROR]+str+"\n", args...)
	}
}
