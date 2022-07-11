package xlogger

type emptyLogger struct {
}

func NewEmptyLogger() Logger {
	return emptyLogger{}
}

func (c emptyLogger) Debug(str string, args ...interface{}) {
}

func (c emptyLogger) Info(str string, args ...interface{}) {

}

func (c emptyLogger) Warn(str string, args ...interface{}) {

}

func (c emptyLogger) Error(str string, args ...interface{}) {
}
