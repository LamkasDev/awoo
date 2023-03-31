package parser_run

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunParser(cparser *parser.AwooParser) parser.AwooParserResult {
	result := parser.AwooParserResult{}
	logger.Log(gchalk.Yellow("\n> Parser\n"))

	globalFunctionIdentifier := node.CreateNodeIdentifier(lexer_token.CreateTokenIdentifier(0, cc.AwooCompilerGlobalFunctionName))
	globalFunctionStatement := statement.CreateStatementFunc(globalFunctionIdentifier.Node)
	parser_context.PushParserScopeFunction(&cparser.Context, parser_context.AwooParserScopeFunction{
		Name: cc.AwooCompilerGlobalFunctionName,
	})

	var err error
	for ; err == nil; err = parser.AdvanceParser(cparser) {
		logger.Log("┏━ %s\n", lexer.PrintToken(&cparser.Settings.Lexer, &cparser.Current))
		st, err := statement_parse.ConstructStatement(cparser, cparser.Current, &parser_details.ConstructStatementDetails{
			EndToken: token.TokenTypeEndStatement,
		})
		if err != nil {
			result.Error = err
			break
		}
		parser.PrintNewStatement(&cparser.Settings, &cparser.Context, &st)
		if st.Type == statement.ParserStatementTypeFunc {
			result.Statements = append(result.Statements, st)
		} else {
			statement.AppendStatementFuncBody(&globalFunctionStatement, st)
		}
	}
	if result.Error != nil {
		panic(result.Error)
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
