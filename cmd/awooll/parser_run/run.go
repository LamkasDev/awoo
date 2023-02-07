package parser_run

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunParser(cparser *parser.AwooParser) parser.AwooParserResult {
	result := parser.AwooParserResult{}
	logger.Log(gchalk.Yellow("\n> Parser\n"))
	logger.Log("Input: %s\n", gchalk.Magenta(fmt.Sprintf("%v", cparser.Contents.Tokens)))
	for ok := true; ok; ok = parser.AdvanceParser(cparser) {
		logger.Log("┏━ %s\n", lexer_token.PrintToken(&cparser.Contents.Context, &cparser.Current))
		st, err := statement_parse.ConstructStatement(cparser, cparser.Current, &statement_parse.ConstructStatementDetails{})
		if err != nil {
			result.Error = err
			break
		}
		statement.PrintNewStatement(&cparser.Contents.Context, &st)
		result.Statements = append(result.Statements, st)
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
