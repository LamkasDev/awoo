package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructStatementFor(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	t, err := parser.AdvanceParser(cparser)
	if err != nil {
		return nil, err
	}
	forInitializationStatement, err := ConstructStatement(cparser, *t, &parser_details.ConstructStatementDetails{
		EndToken: token.TokenTypeEndStatement,
	})
	if err != nil {
		return nil, err
	}
	forStatement := statement.CreateStatementFor(*forInitializationStatement)
	forConditionExpression, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		EndTokens: []uint16{token.TokenTypeEndStatement},
	})
	if err != nil {
		return &forStatement, err
	}
	statement.SetStatementForCondition(&forStatement, forConditionExpression.Node)
	if t, err = parser.AdvanceParser(cparser); err != nil {
		return nil, err
	}
	forAdvancementStatement, err := ConstructStatement(cparser, *t, &parser_details.ConstructStatementDetails{
		EndToken: token.TokenTypeBracketCurlyLeft,
	})
	if err != nil {
		return &forStatement, err
	}
	statement.SetStatementForAdvancement(&forStatement, *forAdvancementStatement)

	forGroup, err := ConstructStatementGroup(cparser, details)
	if err != nil {
		return &forStatement, err
	}
	statement.SetStatementForBody(&forStatement, *forGroup)

	return &forStatement, nil
}
