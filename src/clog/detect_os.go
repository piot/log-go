package clog

import (
	"runtime"
)

type OperatingSystem int

const (
	Linux OperatingSystem = iota
	Windows
	MacOS
)

func DetectOperatingSystem() OperatingSystem {
	s := runtime.GOOS
	if s == "windows" {
		return Windows
	}
	if s == "darwin" {
		return MacOS
	}

	if s == "linux" {
		return Linux
	}

	panic("unknown operating system")
}
