package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructStatementGroup(cparser *parser.AwooParser, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, *parser_error.AwooParserError) {
	body := []statement.AwooParserStatement{}
	for t, err := parser.FetchToken(cparser); err == nil && t.Type != token.TokenTypeBracketCurlyRight; t, err = parser.FetchToken(cparser) {
		bodyStatement, err := ConstructStatement(cparser, t, &parser_details.ConstructStatementDetails{
			EndToken:  token.TokenTypeEndStatement,
			CanReturn: details.CanReturn,
		})
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		body = append(body, bodyStatement)
	}

	return statement.CreateStatementGroup(body), nil
}
