package wzlib_logger

import (
	"io/ioutil"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var _Logger *logrus.Logger

type WzLogger struct{}

// SetLogLevel sets a log level
func (wl *WzLogger) SetLogLevel(level logrus.Level) {
	wl.GetLogger().SetLevel(level)
}

// MuteLoggerToPanic logs only panic messages, when program is crashing
func (wl *WzLogger) MuteLoggerToPanic() {
	wl.SetLogLevel(logrus.PanicLevel)
}

// MuteLogger completely mutes the logger, discarding everything
func (wl *WzLogger) MuteLogger() {
	wl.MuteLoggerToPanic()
	wl.GetLogger().SetOutput(ioutil.Discard)
}

// GetLogger returns lazy-instantiated logger.
func (wl *WzLogger) GetLogger() *logrus.Logger {
	if _Logger == nil {
		_Logger = GetTextLogger(logrus.TraceLevel, nil)
	}
	return _Logger
}

// GetCurrentLogger that was already initialised
func GetCurrentLogger() *logrus.Logger {
	return new(WzLogger).GetLogger()
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
