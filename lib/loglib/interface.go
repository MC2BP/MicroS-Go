package loglib

type LogLevel string

const (
	LevelDebug   LogLevel = "DEBUG"
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
	LevelNone    LogLevel = "NONE"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warning(v ...interface{})
	Warningf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}
