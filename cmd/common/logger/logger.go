package logger

import "fmt"

var AwooLoggerEnabled = true
var AwooLoggerExtraEnabled = true

func Log(format string, v ...any) {
	if !AwooLoggerEnabled {
		return
	}
	fmt.Printf(format, v...)
}

func LogExtra(format string, v ...any) {
	if !AwooLoggerExtraEnabled {
		return
	}
	fmt.Printf(format, v...)
}

func LogError(format string, v ...any) {
	fmt.Printf(format, v...)
}
