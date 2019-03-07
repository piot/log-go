package clog

import (
	"encoding/hex"
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
	Log(level clogint.LogLevel, timeString string, name string, output string)
}

type ConsoleLogger struct {
}

func (c *ConsoleLogger) Log(level clogint.LogLevel, timestring string, name string, output string) {
	var selectedColor *color.Color
	levelString := "unknown"
	switch level {
	case clogint.Info:
		selectedColor = color.New(color.FgCyan)
		levelString = "INFO "
	case clogint.Warning:
		selectedColor = color.New(color.FgYellow)
		levelString = "WARN "
	case clogint.Debug:
		selectedColor = color.New(color.FgGreen)
		levelString = "DEBUG"
	case clogint.Trace:
		selectedColor = color.New(color.FgMagenta)
		levelString = "TRACE"
	case clogint.Error:
		selectedColor = color.New(color.FgRed)
		levelString = "ERROR"
	case clogint.Panic:
		selectedColor = color.New(color.FgHiWhite).Add(color.BgRed)
		levelString = "PANIC"
	}
	color.New(color.FgWhite).Print(timestring + " ")
	selectedColor.Print(levelString + ": ")
	const maxCount = 40
	if len(name) > maxCount {
		name = name[:maxCount]
	}
	padding := strings.Repeat(" ", maxCount-len(name))
	color.New(color.FgHiBlue).Print(name + padding)
	fmt.Println("  " + output)
}

type Log struct {
	logger       Logger
	prefixFields []Field
	prefixName   string
	filterLevel  clogint.LogLevel
}

func DefaultLog() *Log {
	logger := &ConsoleLogger{}
	color.NoColor = false

	return &Log{logger: logger}
}

func convertToString(fields []Field) string {
	s := ""
	for _, f := range fields {
		s += color.CyanString(f.Key) + "="
		switch f.Type {
		case clogint.IntType:
			s += fmt.Sprintf("%v", f.Integer)
		case clogint.UintType:
			s += fmt.Sprintf("%v", f.Integer)
		case clogint.StringType:
			s += fmt.Sprintf("'%v'", f.String)
		case clogint.StringerType:
			s += fmt.Sprintf("'%v'", f.Other)
		case clogint.BoolType:
			s += fmt.Sprintf("%v", f.Integer)
		case clogint.BlobType:
			hexString := hex.Dump(f.Other.([]byte))
			s += hexString
		case clogint.ErrorType:
			s += fmt.Sprintf("%s", f.Other.(error).Error())
		}
		s += " "
	}

	return s
}

func (l *Log) log(level clogint.LogLevel, name string, fields []Field) {
	if level < l.filterLevel {
		return
	}
	t := time.Now()
	timeString := t.UTC().Format(time.RFC3339)
	allFields := append(l.prefixFields, fields...)
	output := convertToString(allFields)
	wholeName := l.prefixName + ":" + name
	l.logger.Log(level, timeString, wholeName, output)
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
	panicString := name + " " + convertToString(fields)
	panic(panicString)
}

func (l *Log) With(name string, fields ...Field) *Log {
	return &Log{logger: l.logger, filterLevel: l.filterLevel, prefixName: name, prefixFields: fields}
}
