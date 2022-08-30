package log

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// DefaultLog default log
var DefaultLog = New()

// New new logrus.Logger
func New() *logrus.Logger {
	return logrus.New()
}

// InitLogger init logger
func InitLogger(w io.Writer) {
	//DefaultLog.Formatter = &logrus.JSONFormatter{}
	DefaultLog.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	DefaultLog.Out = w
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = DefaultLog.Out
	DefaultLog.Level = logrus.DebugLevel // InfoLevel
}

// WithFields ...
func WithFields(fields logrus.Fields) *logrus.Entry {
	return DefaultLog.WithFields(fields)
}

// Debugf debug format
func Debugf(format string, args ...interface{}) {
	DefaultLog.Debugf(format, args...)
}

// Infof info format
func Infof(format string, args ...interface{}) {
	DefaultLog.Infof(format, args...)
}

// Errorf error format
func Errorf(format string, args ...interface{}) {
	DefaultLog.Errorf(format, args...)
}

// Fatalf fatal format
func Fatalf(format string, args ...interface{}) {
	DefaultLog.Fatalf(format, args...)
}
