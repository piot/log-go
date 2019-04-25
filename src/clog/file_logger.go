package clog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/piot/log-go/src/clogint"
)

type FileLogger struct {
	write io.Writer
}

func NewFileLogger(company string, application string) (*FileLogger, error) {
	home, homeErr := homedir.Dir()
	if homeErr != nil {
		return nil, homeErr
	}
	directory := filepath.Join(home, "Library/Logs/", company)
	os.MkdirAll(directory, os.ModePerm)
	filename := filepath.Join(directory, application+".log")
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileLogger{write: f}, nil
}

func (c *FileLogger) Log(level clogint.LogLevel, timestring string, name string, fields []Field) {
	output := ConvertToString(fields)
	levelString := "unknown"
	switch level {
	case clogint.Info:
		levelString = "INFO "
	case clogint.Warning:
		levelString = "WARN "
	case clogint.Debug:
		levelString = "DEBUG"
	case clogint.Trace:
		levelString = "TRACE"
	case clogint.Error:
		levelString = "ERROR"
	case clogint.Panic:
		levelString = "PANIC"
	}
	fmt.Fprintf(c.write, "%s [%s] %s  %s\n", timestring, levelString, name, output)
}
