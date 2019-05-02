package clog

import (
	"github.com/piot/log-go/src/clogint"
)

type MultiLogger struct {
	loggers []Logger
}

func NewMultiLogger(loggers ...Logger) (*MultiLogger, error) {
	return &MultiLogger{loggers: loggers}, nil
}

func (c *MultiLogger) Log(level clogint.LogLevel, timestring string, name string, fields []Field) {
	for _, logger := range c.loggers {
		logger.Log(level, timestring, name, fields)
	}
}
