package logger

import (
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

type Logger struct {
	logger   *zap.SugaredLogger
	operator string
}

func (l *Logger) New(o string) *Logger {
	// Zap Logger Config
	logConf := zap.NewDevelopmentConfig()
	logConf.OutputPaths = append(logConf.OutputPaths, "logfile.log")
	logConf.DisableCaller = true
	logConf.DisableStacktrace = true

	log, _ := logConf.Build(zap.AddCallerSkip(1))
	l.logger = log.Sugar()

	l.operator = color.GreenString("[" + o + "]")

	return l
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Info(append([]interface{}{l.operator, fmt.Sprintf("\t"), msg, ": "}, args...)...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(append([]interface{}{l.operator, fmt.Sprintf("\t"), msg, ": "}, args...)...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Error(append([]interface{}{l.operator, fmt.Sprintf("\t"), msg, ": "}, args...)...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(append([]interface{}{l.operator, fmt.Sprintf("\t"), msg, ": "}, args...)...)
}
