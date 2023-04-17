package util

import (
	"strings"

	"github.com/jwalton/gchalk"
)

func SelectLine(text string, line int) string {
	lines := strings.Split(text, "\n")
	if line < 0 || line > len(lines)-1 {
		return ""
	}
	return lines[line]
}

func HighlightLine(line string, column int, length int) string {
	return line[:column-1] + gchalk.Red(line[column-1:column+length]) + line[column+length:]
}

func FindLineNumber(text string, index int) int {
	if index < 0 || index >= len(text) {
		return 1
	}
	return strings.Count(text[:index], "\n") + 1
}

func FindColumnNumber(text string, index int) int {
	if index < 0 || index >= len(text) {
		return 0
	}
	lastNewlineIndex := strings.LastIndex(text[:index], "\n")
	return index - lastNewlineIndex
}
