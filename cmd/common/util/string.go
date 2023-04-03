package util

import "strings"

func SelectLine(text string, index int) string {
	lineStart := strings.LastIndex(text[:index], "\n") + 1
	lineEnd := index + strings.Index(text[index:], "\n")
	if lineEnd < index {
		lineEnd = len(text)
	}
	return text[lineStart:lineEnd]
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
