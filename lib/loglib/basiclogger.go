package loglib

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type logLevelInt int

const (
	levelAllInt     = 0
	levelDebugInt   = 1
	levelInfoInt    = 2
	levelWarningInt = 3
	levelErrorInt   = 4
	levelNoneInt    = 5
)

func parseLogLevelToInt(level LogLevel) logLevelInt {
	switch level {
	case LevelDebug:
		return levelDebugInt
	case LevelInfo:
		return levelInfoInt
	case LevelWarning:
		return levelWarningInt
	case LevelError:
		return levelErrorInt
	case LevelNone:
		return levelNoneInt
	default:
		return levelAllInt
	}
}

func convLevelItos(level logLevelInt) string {
	switch level {
	case levelDebugInt:
		return string(LevelDebug)
	case levelInfoInt:
		return string(LevelInfo)
	case levelWarningInt:
		return string(LevelWarning)
	case levelErrorInt:
		return string(LevelError)
	default:
		return ""
	}
}

type BasicLogger struct {
	mu sync.Mutex
	out io.Writer
	level logLevelInt
}

func NewBasicLogger(level LogLevel) *BasicLogger {
	return &BasicLogger{
		out:   os.Stdout,
		level: parseLogLevelToInt(level),
	}
}

func (l *BasicLogger) Debug(v ...interface{}) {
	l.log(levelDebugInt, v...)
}

func (l *BasicLogger) Debugf(format string, v ...interface{}) {
	l.logf(levelDebugInt, format, v...)
}

func (l *BasicLogger) Info(v ...interface{}) {
	l.log(levelInfoInt, v...)
}

func (l *BasicLogger) Infof(format string, v ...interface{}) {
	l.logf(levelInfoInt, format, v...)
}

func (l *BasicLogger) Warning(v ...interface{}) {
	l.log(levelWarningInt, v...)
}

func (l *BasicLogger) Warningf(format string, v ...interface{}) {
	l.logf(levelWarningInt, format, v...)
}

func (l *BasicLogger) Error(v ...interface{}) {
	l.log(levelErrorInt, v...)
}

func (l *BasicLogger) Errorf(format string, v ...interface{}) {
	l.logf(levelErrorInt, format, v...)
}

func (l *BasicLogger) log(level logLevelInt, v ...interface{}) {
	// prepare message
	if level < l.level {
		return
	}
	s := format(level, fmt.Sprint(v...))
	l.out.Write([]byte(s))
}

func (l *BasicLogger) logf(level logLevelInt, f string, v ...interface{}) {
	// prepare message
	if level < l.level {
		return
	}
	s := format(level, fmt.Sprintf(f, v...))
	l.out.Write([]byte(s))
}

func format(level logLevelInt, s string) string {
	return fmt.Sprintf("%-7s %s: %s\n", convLevelItos(level), time.Now().Format(time.ANSIC), s)
}
