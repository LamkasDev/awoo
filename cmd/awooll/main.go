package main

import (
	"flag"
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_run"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
)

func main() {
	u, _ := user.Current()
	defaultInput := path.Join(u.HomeDir, "Documents", "awoo", "data", "input.awoo")
	defaultOutput := path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj")

	var input string
	var output string
	var quiet bool
	flag.StringVar(&input, "i", defaultInput, "path to input .awoo file")
	flag.StringVar(&output, "o", defaultOutput, "path to output .awooobj file")
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AwooLoggerEnabled = !quiet
	input, output = paths.ResolvePaths(input, ".awoo", output, ".awoobj")
	flags.ResolveColor()

	file, err := os.ReadFile(input)
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
	parRes := parser_run.RunParser(&par)

	compSettings := compiler.AwooCompilerSettings{
		Path: output,
	}
	comp := compiler.SetupCompiler(compSettings, par.Context)
	compiler.LoadCompiler(&comp, parRes)
	compiler.RunCompiler(&comp)
}
