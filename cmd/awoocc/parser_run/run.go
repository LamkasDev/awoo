package parser_run

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunParser(cparser *parser.AwooParser) parser.AwooParserResult {
	result := parser.AwooParserResult{}
	logger.Log(gchalk.Yellow("\n> Parser\n"))
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
		result.Statements = append(result.Statements, st)
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
