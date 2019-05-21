package clog

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/piot/log-go/src/clogint"
)

type Field = clogint.Field

func Bool(key string, val bool) Field {
	var ival int64
	if val {
		ival = 1
	}
	return Field{Key: key, Type: clogint.BoolType, Integer: ival}
}

func Blob(key string, payload []uint8) Field {
	return Field{Key: key, Type: clogint.BlobType, Other: payload}
}

func Error(key string, err error) Field {
	return Field{Key: key, Type: clogint.ErrorType, Other: err}
}

func Int(key string, val int) Field {
	return Field{Key: key, Type: clogint.IntType, Integer: int64(val)}
}

func Int8(key string, val int8) Field {
	return Field{Key: key, Type: clogint.IntType, Integer: int64(val)}
}

func Int16(key string, val int16) Field {
	return Field{Key: key, Type: clogint.IntType, Integer: int64(val)}
}

func Int32(key string, val int32) Field {
	return Field{Key: key, Type: clogint.IntType, Integer: int64(val)}
}

func Int64(key string, val int64) Field {
	return Field{Key: key, Type: clogint.IntType, Integer: int64(val)}
}

func String(key string, val string) Field {
	return Field{Key: key, Type: clogint.StringType, String: val}
}

func Stringer(key string, val fmt.Stringer) Field {
	return Field{Key: key, Type: clogint.StringerType, Other: val}
}

func Interface(key string, val interface{}) Field {
	return Field{Key: key, Type: clogint.InterfaceType, Other: val}
}

func Uint(key string, val uint) Field {
	return Field{Key: key, Type: clogint.UintType, Integer: int64(val)}
}

func Uint8(key string, val uint8) Field {
	return Field{Key: key, Type: clogint.UintType, Integer: int64(val)}
}

func Uint16(key string, val uint16) Field {
	return Field{Key: key, Type: clogint.UintType, Integer: int64(val)}
}

func Uint32(key string, val uint32) Field {
	return Field{Key: key, Type: clogint.UintType, Integer: int64(val)}
}

func Uint64(key string, val uint64) Field {
	return Field{Key: key, Type: clogint.UintType, Integer: int64(val)}
}

type Logger interface {
	Log(level clogint.LogLevel, timeString string, name string, fields []Field)
}

type Log struct {
	logger       Logger
	prefixFields []Field
	prefixName   string
	filterLevel  clogint.LogLevel
}

func DefaultLog() *Log {
	logger := NewConsoleLogger()
	color.NoColor = false

	return &Log{logger: logger}
}

func DefaultFileLog(organization string, applicationName string) (*Log, error) {
	directory, locationErr := logLocation(organization, applicationName)
	if locationErr != nil {
		return nil, locationErr
	}

	logger, err := NewFileLogger(directory, applicationName)
	if err != nil {
		return nil, err
	}
	color.NoColor = false

	return &Log{logger: logger}, nil
}

func NewFileLog(directory string, applicationName string) (*Log, error) {
	logger, err := NewFileLogger(directory, applicationName)
	if err != nil {
		return nil, err
	}
	color.NoColor = false

	return &Log{logger: logger}, nil
}

func NewMultiLog(loggers ...Logger) (*Log, error) {
	logger, loggerErr := NewMultiLogger(loggers...)
	if loggerErr != nil {
		return nil, loggerErr
	}
	return &Log{logger: logger}, nil
}

func (l *Log) log(level clogint.LogLevel, name string, fields []Field) {
	if level < l.filterLevel {
		return
	}
	t := time.Now()
	timeString := t.UTC().Format(time.RFC3339)
	allFields := append(l.prefixFields, fields...)
	wholeName := name
	if l.prefixName != "" {
		wholeName = l.prefixName + ":" + wholeName
	}
	l.logger.Log(level, timeString, wholeName, allFields)
}

func stringToLogLevel(logLevelString string, defaultLevel clogint.LogLevel) clogint.LogLevel {
	logLevelString = strings.ToLower(logLevelString)
	switch logLevelString {
	case "trace":
		return clogint.Trace
	case "debug":
		return clogint.Debug
	case "info":
		return clogint.Info
	case "warn":
		return clogint.Warning
	case "error":
		return clogint.Error
	case "panic":
		return clogint.Panic
	}

	return defaultLevel
}

func (l *Log) SetLogLevelUsingString(logLevelString string, defaultLevel clogint.LogLevel) {
	logLevel := stringToLogLevel(logLevelString, defaultLevel)
	l.SetLogLevel(logLevel)
}

func (l *Log) SetLogLevel(level clogint.LogLevel) {
	l.filterLevel = level
}

func (l *Log) Err(err error) {
	l.log(clogint.Error, "error", []Field{Error("err", err)})
}

func (l *Log) Error(name string, fields ...Field) {
	l.log(clogint.Error, name, fields)
}

func (l *Log) Warn(name string, fields ...Field) {
	l.log(clogint.Warning, name, fields)
}

func (l *Log) Info(name string, fields ...Field) {
	l.log(clogint.Info, name, fields)
}

func (l *Log) Debug(name string, fields ...Field) {
	l.log(clogint.Debug, name, fields)
}

func (l *Log) Trace(name string, fields ...Field) {
	l.log(clogint.Trace, name, fields)
}

func (l *Log) Panic(name string, fields ...Field) {
	l.log(clogint.Panic, name, fields)
	panicString := name + " " + ConvertToString(fields)
	panic(panicString)
}

func (l *Log) With(name string, fields ...Field) *Log {
	return &Log{logger: l.logger, filterLevel: l.filterLevel, prefixName: name, prefixFields: fields}
}

func (l *Log) CheckTrace() bool {
	return l.filterLevel <= clogint.Trace
}

func (l *Log) CheckDebug() bool {
	return l.filterLevel <= clogint.Debug
}

func (l *Log) CheckInfo() bool {
	return l.filterLevel <= clogint.Info
}
