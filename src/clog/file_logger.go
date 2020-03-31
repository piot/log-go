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

func findXDGConfig(home string) string {
	xdkConfig := os.Getenv("XDG_CONFIG_HOME")
	if xdkConfig != "" {
		return xdkConfig
	}

	return filepath.Join(home, ".config")
}

func logLocation(organization string, application string) (string, error) {
	home, homeErr := homedir.Dir()
	if homeErr != nil {
		return "", homeErr
	}

	os := DetectOperatingSystem()
	switch os {
	case Windows:
		logDirectory := filepath.Join(home, organization, application, "logs")
		return logDirectory, nil
	case MacOS:
		return filepath.Join(home, "Library/Logs/", organization), nil
	case Linux:
		xdgConfig := findXDGConfig(home)
		logDirectory := filepath.Join(xdgConfig, organization, application, "logs")
		return logDirectory, nil

	}

	return "", fmt.Errorf("clog: unknown operating system")
}

func NewFileLogger(directory string, application string) (*FileLogger, error) {
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
