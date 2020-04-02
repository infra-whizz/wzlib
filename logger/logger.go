package wzlib_logger

import (
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type WzLogger struct {
	log   *logrus.Logger
	level logrus.Level
}

// SetLogLevel sets a log level
func (wl *WzLogger) SetLogLevel(level logrus.Level) {
	wl.level = level
	if wl.log != nil {
		wl.log.SetLevel(wl.level)
	}
}

// GetLogger returns lazy-instantiated logger.
func (wl *WzLogger) GetLogger() *logrus.Logger {
	if wl.log == nil {
		wl.log = GetTextLogger(logrus.TraceLevel, nil)
	}
	return wl.log
}

// GetTextLogger create a logger instance
func GetTextLogger(level logrus.Level, out *os.File) *logrus.Logger {
	formatter := new(nested.Formatter)
	formatter.HideKeys = true
	formatter.FieldsOrder = []string{"component", "category"}
	formatter.ShowFullLevel = true

	logger := logrus.New()

	if out == nil {
		logger.Out = os.Stderr
		formatter.NoColors = false
		formatter.NoFieldsColors = false
	} else {
		logger.Out = out
		formatter.NoColors = true
		formatter.NoFieldsColors = true
	}

	logger.Level = level
	logger.SetFormatter(formatter)

	return logger
}

// GormLogger object
type GormLogger struct {
	lgr *logrus.Logger
}

// NewGormLogger creates an instance of GormLogger
func NewGormLogger(lg *logrus.Logger) *GormLogger {
	lgr := new(GormLogger)
	lgr.SetLogger(lg)
	return lgr
}

// SetLogger creates a new logrus instance of the logger or assigns an existing one
func (lgr *GormLogger) SetLogger(lg *logrus.Logger) *GormLogger {
	if lg == nil {
		lg = GetTextLogger(logrus.DebugLevel, nil)
	}
	lgr.lgr = lg
	return lgr
}

// Print prints logging data of GORM
func (lgr *GormLogger) Print(v ...interface{}) {
	var data interface{}
	if v[0] == "sql" {
		data = v[3]
	} else if v[0] == "log" {
		data = v[2]
	}
	lgr.lgr.Debugln(data)
}
