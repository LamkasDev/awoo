package parser_run

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement_parse"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunParser(cparser *parser.AwooParser) parser.AwooParserResult {
	result := parser.AwooParserResult{}
	logger.Log(gchalk.Yellow("\n> Parser\n"))
	logger.Log("Input: %s\n", gchalk.Magenta(fmt.Sprintf("%v", cparser.Contents.Tokens)))
	var err error = nil
	for ; err == nil; err = parser.AdvanceParser(cparser) {
		logger.Log("┏━ %s\n", lexer_token.PrintToken(&cparser.Contents.Context, &cparser.Current))
		st, err := statement_parse.ConstructStatement(cparser, cparser.Current, &parser_details.ConstructStatementDetails{})
		if err != nil {
			result.Error = err
			break
		}
		parser.PrintNewStatement(cparser, &st)
		result.Statements = append(result.Statements, st)
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
