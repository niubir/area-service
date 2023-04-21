package logs

import (
	"time"

	"github.com/niubir/area-service/config"
	"github.com/niubir/logger"
)

var (
	l *logger.Logger
)

func Init() error {
	var err error
	l, err = NewLogger("area_service")
	return err
}

func Debug(s string, i ...interface{}) {
	l.Debug(s, i...)
}

func Info(s string, i ...interface{}) {
	l.Info(s, i...)
}

func Warn(s string, i ...interface{}) {
	l.Warn(s, i...)
}

func Error(s string, i ...interface{}) {
	l.Error(s, i...)
}

func NewLogger(prefix string) (*logger.Logger, error) {
	loggerLevel := logger.ErrorLevel
	loggerWithStdout := false
	loggerWithStack := false
	if config.Config.Debug {
		loggerLevel = logger.DebugLevel
		loggerWithStdout = true
		loggerWithStack = true
	}

	return logger.NewLogger(
		logger.WithLevel(loggerLevel),
		logger.WithTimeFormat(time.RFC3339Nano),
		logger.WithStdout(loggerWithStdout),
		logger.WithPath("./logs"),
		logger.WithPrefix(prefix),
		logger.WithDuration(24*time.Hour),
		logger.WithMaxByte(10*1024*1024),
		logger.WithStack(loggerWithStack),
	)
}

func NewFileLogger(prefix string) (*logger.FileLogger, error) {
	loggerLevel := logger.ErrorLevel
	loggerWithStdout := false
	loggerWithStack := false
	if config.Config.Debug {
		loggerLevel = logger.DebugLevel
		loggerWithStdout = true
		loggerWithStack = true
	}

	return logger.NewFileLogger(
		logger.WithLevel(loggerLevel),
		logger.WithTimeFormat(time.RFC3339Nano),
		logger.WithStdout(loggerWithStdout),
		logger.WithPath("./logs"),
		logger.WithPrefix(prefix),
		logger.WithDuration(24*time.Hour),
		logger.WithMaxByte(10*1024*1024),
		logger.WithStack(loggerWithStack),
	)
}
