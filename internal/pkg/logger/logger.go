package logger

import "io"

var Log Logger

// logger interface based on logrus' exported methods
// since this is an interface, we can also swap to other log libraries if they also support log levels, and structured logging with fields
type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	WithFields(args map[string]interface{}) Logger
	GetWriter() io.Writer
	Printf(format string, args ...interface{})
}

// SetLogger set to exported to make it easier when we want to mock the logger during tests
func SetLogger(logger Logger) {
	Log = logger
}
