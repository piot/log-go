package clog

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/piot/log-go/src/clogint"
)

type ConsoleLogger struct {
	output io.Writer
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{output: os.Stderr}
}

func convertToColorString(fields []Field) string {
	s := ""
	for _, f := range fields {
		value := fieldValueToString(f)
		lines := strings.Split(value, "\n")
		isMultiline := len(lines) > 1
		if isMultiline {
			s += "\n  "
		}
		s += color.CyanString(f.Key) + "="
		if isMultiline {
			for index, line := range lines {
				if index > 0 {
					s += "\n    "
				}
				s += color.GreenString(line)
			}
		} else {
			s += color.GreenString(fieldValueToString(f))
		}
		s += " "
	}

	return s
}

func (c *ConsoleLogger) Log(level clogint.LogLevel, timestring string, name string, fields []Field) {
	output := convertToColorString(fields)
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
	color.New(color.FgWhite).Fprint(c.output, timestring+" ")
	selectedColor.Fprint(c.output, levelString+": ")
	const maxCount = 40
	if len(name) > maxCount {
		name = name[:maxCount]
	}
	padding := strings.Repeat(" ", maxCount-len(name))
	color.New(color.FgHiBlue).Fprint(c.output, name+padding)
	fmt.Fprintln(c.output, "  "+output)
}
