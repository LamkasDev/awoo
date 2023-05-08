package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructStatementGroup(cparser *parser.AwooParser, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	body := []statement.AwooParserStatement{}
	for t, err := parser.AdvanceParser(cparser); err == nil && t.Type != token.TokenTypeBracketCurlyRight; t, err = parser.AdvanceParser(cparser) {
		bodyStatement, err := ConstructStatement(cparser, *t, &parser_details.ConstructStatementDetails{
			EndToken:  token.TokenTypeEndStatement,
			CanReturn: details.CanReturn,
		})
		if err != nil {
			return nil, err
		}
		body = append(body, *bodyStatement)
	}
	groupStatement := statement.CreateStatementGroup(body)

	return &groupStatement, nil
}
