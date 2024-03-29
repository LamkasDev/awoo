package parser_run

import (
	"fmt"
	"os"
	"strings"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

// TODO: avoid empty global function

func LoadParserSymbols(cparser *parser.AwooParser, celf *elf.AwooElf) error {
	scope.PushFunction(&cparser.Context.Scopes, scope.AwooScopeFunction{
		Name: cc.AwooCompilerGlobalFunctionName,
	})
	for _, symbol := range celf.SymbolTable.External {
		if _, err := scope.PushFunctionBlockSymbolExternal(&cparser.Context.Scopes, symbol); err != nil {
			return err
		}
	}

	return nil
}

func RunParser(cparser *parser.AwooParser) parser.AwooParserResult {
	result := parser.AwooParserResult{}
	logger.LogExtra(gchalk.Yellow("\n> Parser\n"))

	globalFunctionIdentifier := node.CreateNodeIdentifier(lexer_token.CreateTokenIdentifier(lexer_token.AwooLexerTokenPosition{}, cc.AwooCompilerGlobalFunctionName))
	globalFunctionStatement := statement.CreateStatementFunc(globalFunctionIdentifier.Node)

	for t, err := parser.AdvanceParserFor(cparser, 0); err == nil; t, err = parser.AdvanceParser(cparser) {
		logger.LogExtra("┏━ %s\n", lexer.PrintToken(&cparser.Settings.Lexer, t))
		st, err := statement_parse.ConstructStatement(cparser, *t, &parser_details.ConstructStatementDetails{
			EndToken: token.TokenTypeEndStatement,
		})
		if err != nil {
			PrintParserError(cparser, err)
		}
		if st == nil {
			continue
		}
		parser.PrintNewStatement(&cparser.Settings, &cparser.Context, st)
		if st.Type == statement.ParserStatementTypeFunc {
			result.Statements = append(result.Statements, *st)
		} else {
			statement.AppendStatementFuncBody(&globalFunctionStatement, *st)
		}
	}

	result.Context = cparser.Context
	result.Statements = append([]statement.AwooParserStatement{globalFunctionStatement}, result.Statements...)
	return result
}

func PrintParserError(cparser *parser.AwooParser, err *parser_error.AwooParserError) {
	fmt.Printf("%s: %s\n", gchalk.Red(fmt.Sprintf("error[E%#3x]", err.Type)), err.Message)
	text := strings.ReplaceAll(string(cparser.Contents.Data.Contents.Text), "\t", strings.Repeat(" ", cc.AwooTabIndent))
	for _, highlight := range err.Highlights {
		line := util.HighlightLine(util.SelectLine(text, int(highlight.Position.Line-1)), int(highlight.Position.Column), int(highlight.Position.Length))
		fmt.Printf(" %s %s:%d:%d\n", gchalk.Blue("---->"), cparser.Settings.Lexer.Path, highlight.Position.Line, highlight.Position.Column)
		fmt.Printf("   %s \n", gchalk.Blue("|"))
		fmt.Printf(" %3s %s %s\n", gchalk.Blue(fmt.Sprint(highlight.Position.Line)), gchalk.Blue("|"), line)
		fmt.Printf("   %s %s%s %s\n",
			gchalk.Blue("|"),
			strings.Repeat(" ", int(highlight.Position.Column-1)),
			gchalk.Red(strings.Repeat("^", int(highlight.Position.Length))),
			gchalk.Red(highlight.Details),
		)
	}
	os.Exit(0)
}
