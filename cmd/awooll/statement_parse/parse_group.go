package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementGroup(cparser *parser.AwooParser, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	body := []statement.AwooParserStatement{}
	for t, err := parser.FetchToken(cparser); err == nil && t.Type != token.TokenTypeBracketCurlyRight; t, err = parser.FetchToken(cparser) {
		bodyStatement, err := ConstructStatement(cparser, t, details)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		body = append(body, bodyStatement)
	}

	return statement.CreateStatementGroup(body), nil
}
