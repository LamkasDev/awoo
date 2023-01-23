package lexer

import "testing"

func TestLexer(t *testing.T) {
	file := "int 123 int int"

	lexer := SetupLexer()
	LoadLexer(&lexer, []rune(file))
	RunLexer(&lexer)
}
