package parser_run

import (
	"fmt"
	"os"
	"strings"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

func RunParser(cparser *parser.AwooParser) parser.AwooParserResult {
	result := parser.AwooParserResult{
		Context: cparser.Context,
	}
	logger.LogExtra(gchalk.Yellow("\n> Parser\n"))

	globalFunctionIdentifier := node.CreateNodeIdentifier(lexer_token.CreateTokenIdentifier(0, cc.AwooCompilerGlobalFunctionName))
	globalFunctionStatement := statement.CreateStatementFunc(globalFunctionIdentifier.Node)
	parser_context.PushParserScopeFunction(&cparser.Context, parser_context.AwooParserScopeFunction{
		Name: cc.AwooCompilerGlobalFunctionName,
	})

	var err *parser_error.AwooParserError
	for ; err == nil; err = parser.AdvanceParser(cparser) {
		logger.LogExtra("┏━ %s\n", lexer.PrintToken(&cparser.Settings.Lexer, &cparser.Current))
		st, err := statement_parse.ConstructStatement(cparser, cparser.Current, &parser_details.ConstructStatementDetails{
			EndToken: token.TokenTypeEndStatement,
		})
		if err != nil {
			PrintParserError(cparser, err)
		}
		parser.PrintNewStatement(&cparser.Settings, &cparser.Context, &st)
		if st.Type == statement.ParserStatementTypeFunc {
			result.Statements = append(result.Statements, st)
		} else {
			statement.AppendStatementFuncBody(&globalFunctionStatement, st)
		}
	}

	statement.AppendStatementFuncBody(&globalFunctionStatement, statement.CreateStatementReturn(nil))
	parser_context.PopParserScopeCurrentFunction(&cparser.Context)
	parser_context.PushParserFunction(&cparser.Context, parser_context.AwooParserFunction{
		Name:      cc.AwooCompilerGlobalFunctionName,
		Arguments: []statement.AwooParserStatementFuncArgument{},
	})
	result.Statements = append([]statement.AwooParserStatement{globalFunctionStatement}, result.Statements...)

	return result
}

func PrintParserError(cparser *parser.AwooParser, err *parser_error.AwooParserError) {
	fmt.Printf("%s: %s\n", gchalk.Red(fmt.Sprintf("error[E%#3x]", err.Type)), err.Message)
	for _, highlight := range err.Highlights {
		lineNumber := util.FindLineNumber(string(cparser.Contents.Text), int(highlight.Start))
		columnNumber := util.FindColumnNumber(string(cparser.Contents.Text), int(highlight.Start))
		line := util.SelectLine(string(cparser.Contents.Text), int(highlight.Start))
		fmt.Printf(" %s %s:%d:%d\n", gchalk.Blue("---->"), cparser.Settings.Lexer.Path, lineNumber, columnNumber)
		fmt.Printf("   %s \n", gchalk.Blue("|"))
		fmt.Printf(" %3s %s %s\n", gchalk.Blue(fmt.Sprint(lineNumber)), gchalk.Blue("|"), line)
		fmt.Printf("   %s %s%s %s\n",
			gchalk.Blue("|"),
			strings.Repeat(" ", columnNumber),
			gchalk.Red(strings.Repeat("^", int(highlight.Length))),
			gchalk.Red(highlight.Details),
		)
	}
	os.Exit(0)
}
