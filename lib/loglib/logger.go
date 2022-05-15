package loglib

var (
	logger Logger = NewBasicLogger(LevelInfo)
)

func SetLogger(l Logger) {
	logger = l
}

func Info(v ...interface{}) {
	logger.Info(v)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v)
}

func Debug(v ...interface{}) {
	logger.Debug(v)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v)
}

func Warning(v ...interface{}) {
	logger.Warning(v)
}

func Warningf(format string, v ...interface{}) {
	logger.Warningf(format, v)
}

func Error(v ...interface{}) {
	logger.Error(v)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v)
}

