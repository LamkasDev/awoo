package main

import (
	"os"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
)

func main() {
	file, err := os.ReadFile("C:\\Users\\PC\\Documents\\code\\go\\awoo-emu\\data\\input.txt")
	if err != nil {
		panic(err)
	}

	lexSettings := lexer.AwooLexerSettings{}
	lex := lexer.SetupLexer(lexSettings)
	lexer.LoadLexer(&lex, []rune(string(file)))
	lexRes := lexer.RunLexer(&lex)

	parSettings := parser.AwooParserSettings{}
	par := parser.SetupParser(parSettings, lex.Context)
	parser.LoadParser(&par, lexRes)
	parRes := parser.RunParser(&par)

	compSettings := compiler.AwooCompilerSettings{}
	comp := compiler.SetupCompiler(compSettings, par.Context)
	compiler.LoadCompiler(&comp, parRes)
	compiler.RunCompiler(&comp)
}
