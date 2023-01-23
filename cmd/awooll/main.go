package main

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
)

func main() {
	file := "int 123 int int"
	lex := lexer.SetupLexer()
	lexer.LoadLexer(&lex, []rune(file))
	lexer.RunLexer(&lex)
}
