package logrus

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"strings"
)

type Logger struct {
	level      log.Level
	baseLogger *log.Logger
}

func convertLevel(strLevel string) log.Level {
	var level log.Level = log.DebugLevel
	switch strings.ToLower(strLevel) {
	case "panic":
		level = log.PanicLevel
	case "fatal":
		level = log.FatalLevel
	case "error":
		level = log.ErrorLevel
	case "warn":
		level = log.WarnLevel
	case "info":
		level = log.InfoLevel
	case "debug":
		level = log.DebugLevel
	case "trace":
		level = log.TraceLevel
	}
	return level
}

func newWithLogger(config *Config, logger *log.Logger) *Logger {
	return &Logger{
		level:      convertLevel(config.Level),
		baseLogger: logger,
	}
}

func newWithConfig(config *Config) (*Logger, error) {
	var (
		level log.Level = convertLevel(config.Level)
		pathname string = config.Path
		baseLogger *log.Logger
		errorWriter, infoWriter io.Writer
		err error
	)

	filedMap := log.FieldMap{
		log.FieldKeyTime:  "time",
		log.FieldKeyLevel: "level",
		log.FieldKeyMsg:   "msg",
		log.FieldKeyFunc: "caller",
	}
	if pathname != "" {
		errorFileName := "error.log"
		infoFileName := "access.log"
		errorPath := filepath.Join(pathname, errorFileName)
		infoPath := filepath.Join(pathname, infoFileName)
		if errorWriter, err = rotatelogs.New(
			errorPath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(errorPath),
			//rotatelogs.WithMaxAge(24*time.Hour),
			//rotatelogs.WithRotationTime(24*time.Hour),
			rotatelogs.WithClock(rotatelogs.Local),
		); err != nil {
			return nil, err
		}
		if infoWriter, err = rotatelogs.New(
			infoPath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(infoPath),
			//rotatelogs.WithMaxAge(24*time.Hour),
			//rotatelogs.WithRotationTime(24*time.Hour),
			rotatelogs.WithClock(rotatelogs.Local),
		); err != nil {
			return nil, err
		}

		baseLogger = log.New()
		baseLogger.Level = level
		baseLogger.Hooks.Add(lfshook.NewHook(
			lfshook.WriterMap{
				log.InfoLevel:  infoWriter,
				log.DebugLevel: infoWriter,
				log.WarnLevel:  infoWriter,
				log.TraceLevel: infoWriter,
				log.ErrorLevel: errorWriter,
				log.FatalLevel: errorWriter,
				log.PanicLevel: errorWriter,
			},
			&log.JSONFormatter{
				TimestampFormat: config.TimestampFormat,
				DisableTimestamp: config.DisableTimestamp,
				FieldMap: filedMap,
			},
		))
	} else {
		baseLogger = log.New()
		baseLogger.Formatter = &log.TextFormatter{
			TimestampFormat: config.TimestampFormat,
			DisableTimestamp: config.DisableTimestamp,
			FullTimestamp: config.FullTimestamp,
			DisableColors: config.DisableColors,
			FieldMap:      filedMap,
		}
		baseLogger.Level = level
	}
	baseLogger.SetReportCaller(config.ReportCaller)
	logger := new(Logger)
	logger.level = level
	logger.baseLogger = baseLogger
	return logger, nil
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.baseLogger.Debug(fmt.Sprintf(format, a...))
}

func (logger *Logger) Info(format string, a ...interface{}) {
	logger.baseLogger.Info(fmt.Sprintf(format, a...))
}

func (logger *Logger) Error(format string, a ...interface{}) {
	logger.baseLogger.Error(fmt.Sprintf(format, a...))
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	logger.baseLogger.Fatal(fmt.Sprintf(format, a...))
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	logger.baseLogger.Warn(fmt.Sprintf(format, a...))
}

func (logger *Logger) Trace(format string, a ...interface{}) {
	logger.baseLogger.Trace(fmt.Sprintf(format, a...))
}

func (logger *Logger) GetLogger() *log.Logger {
	return logger.baseLogger
}

var gLogger, _ = newWithConfig(DefaultConfig())

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	gLogger.Info(format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.Fatal(format, a...)
}

func Warn(format string, a ...interface{}) {
	gLogger.Warn(format, a...)
}

func Trace(format string, a ...interface{}) {
	gLogger.Trace(fmt.Sprintf(format, a...))
}

func GetLogger() *log.Logger {
	return gLogger.baseLogger
}