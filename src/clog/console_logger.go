package clog

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/piot/log-go/src/clogint"
)

type ConsoleLogger struct {
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func convertToColorString(fields []Field) string {
	s := ""
	for _, f := range fields {
		s += color.CyanString(f.Key) + "="
		s += color.GreenString(fieldValueToString(f))
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
