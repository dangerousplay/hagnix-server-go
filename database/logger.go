package database

import (
	"github.com/InVisionApp/go-logger"
	"github.com/go-xorm/core"
)

var bLogger = log.NewSimple()

type dBLogger struct{}

func (dBLogger) Debug(v ...interface{}) {
	bLogger.Debug(v)
}

// Debugf empty implementation
func (dBLogger) Debugf(format string, v ...interface{}) {
	bLogger.Debugf(format, v)
}

// Error empty implementation
func (dBLogger) Error(v ...interface{}) {
	bLogger.Error(v)
}

// Errorf empty implementation
func (dBLogger) Errorf(format string, v ...interface{}) {
	bLogger.Errorf(format, v)
}

// Info empty implementation
func (dBLogger) Info(v ...interface{}) {
	bLogger.Info(v)
}

// Infof empty implementation
func (dBLogger) Infof(format string, v ...interface{}) {
	bLogger.Infof(format, v)
}

// Warn empty implementation
func (dBLogger) Warn(v ...interface{}) {
	bLogger.Warn(v)
}

// Warnf empty implementation
func (dBLogger) Warnf(format string, v ...interface{}) {
	bLogger.Warnf(format, v)
}

// Level empty implementation
func (dBLogger) Level() core.LogLevel {
	return core.LOG_UNKNOWN
}

func (dBLogger) IsShowSQL() bool {
	return false
}

func (dBLogger) ShowSQL(show ...bool) {

}

func (dBLogger) SetLevel(l core.LogLevel) {

}
