package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementFor(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	if err := parser.AdvanceParser(cparser); err != nil {
		return statement.AwooParserStatement{}, err
	}
	forInitializationStatement, err := ConstructStatement(cparser, cparser.Current, &parser_details.ConstructStatementDetails{
		EndToken: token.TokenTypeEndStatement,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	forStatement := statement.CreateStatementFor(forInitializationStatement)
	forConditionExpression, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		EndToken: token.TokenTypeEndStatement,
	})
	if err != nil {
		return forStatement, err
	}
	statement.SetStatementForCondition(&forStatement, forConditionExpression.Node)
	if err := parser.AdvanceParser(cparser); err != nil {
		return statement.AwooParserStatement{}, err
	}
	forAdvancementStatement, err := ConstructStatement(cparser, cparser.Current, &parser_details.ConstructStatementDetails{
		EndToken: token.TokenTypeBracketCurlyLeft,
	})
	if err != nil {
		return forStatement, err
	}
	statement.SetStatementForAdvancement(&forStatement, forAdvancementStatement)

	forGroup, err := ConstructStatementGroup(cparser, details)
	if err != nil {
		return forStatement, err
	}
	statement.SetStatementForBody(&forStatement, forGroup)

	return forStatement, nil
}
