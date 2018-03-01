package glog

import (
	"fmt"
	"runtime"
	"strings"
)

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/* ================================================================================
 * 日志 for uber zap
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

type (
	Logger struct {
		log  *zap.Logger
		atom zap.AtomicLevel
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化Log
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewLogger(name string, level Level, args ...bool) *Logger {
	isProduct := true
	if len(args) > 0 {
		isProduct = args[0]
	}

	zapLogger, _ := newZapLogger(level, isProduct)
	atom := zap.NewAtomicLevel()

	logger := &Logger{
		log:  zapLogger.Named(name).WithOptions(zap.AddCallerSkip(1)),
		atom: atom,
	}

	return logger
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化ZapLog
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func newZapLogger(level Level, isProduct bool) (*zap.Logger, error) {
	var loggerConfig zap.Config
	if isProduct {
		loggerConfig = zap.NewProductionConfig()
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(
			strings.Join([]string{
				caller.TrimmedPath(),
				runtime.FuncForPC(caller.PC).Name(),
			}, ":"),
		)
	}

	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(level))

	return loggerConfig.Build()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * SetLevel
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) SetLevel(level Level) {
	s.atom.SetLevel(zapcore.Level(level))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Print
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Print(args ...interface{}) {
	s.Info(args...)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Print of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Printf(format string, args ...interface{}) {
	s.Infof(format, args...)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * PrintField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) PrintField(msg string, key string, value interface{}) {
	s.InfoField(msg, key, value)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Debug
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Debug(args ...interface{}) {
	s.log.Debug(fmt.Sprint(args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Debug of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Debugf(format string, args ...interface{}) {
	s.log.Debug(fmt.Sprintf(format, args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * DebugField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) DebugField(msg string, key string, value interface{}) {
	s.log.Debug(msg, zap.Any(key, value))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Info
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Info(args ...interface{}) {
	s.log.Info(fmt.Sprint(args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Info of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Infof(format string, args ...interface{}) {
	s.log.Info(fmt.Sprintf(format, args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * InfoField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) InfoField(msg string, key string, value interface{}) {
	s.log.Info(msg, zap.Any(key, value))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Warn
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Warn(args ...interface{}) {
	s.log.Warn(fmt.Sprint(args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Warn of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Warnf(format string, args ...interface{}) {
	s.log.Warn(fmt.Sprintf(format, args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * WarnField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) WarnField(msg string, key string, value interface{}) {
	s.log.Warn(msg, zap.Any(key, value))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Error
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Error(args ...interface{}) {
	s.log.Error(fmt.Sprint(args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Error of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Errorf(format string, args ...interface{}) {
	s.log.Error(fmt.Sprintf(format, args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * ErrorField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) ErrorField(msg string, key string, value interface{}) {
	s.log.Error(msg, zap.Any(key, value))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Fatal
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Fatal(args ...interface{}) {
	s.log.Fatal(fmt.Sprint(args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Fatal of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Fatalf(format string, args ...interface{}) {
	s.log.Fatal(fmt.Sprintf(format, args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * FatalField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) FatalField(msg string, key string, value interface{}) {
	s.log.Fatal(msg, zap.Any(key, value))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Panic
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Panic(args ...interface{}) {
	s.log.Panic(fmt.Sprint(args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Panic of format
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) Panicf(format string, args ...interface{}) {
	s.log.Panic(fmt.Sprintf(format, args...))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * PanicField
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Logger) PanicField(msg string, key string, value interface{}) {
	s.log.Panic(msg, zap.Any(key, value))
}
