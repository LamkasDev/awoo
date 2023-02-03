package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementIfBody(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
	body := []statement.AwooParserStatement{}
	for t, err := parser.FetchTokenParser(cparser); err == nil && t.Type != token.TokenTypeBracketCurlyRight; t, err = parser.FetchTokenParser(cparser) {
		bodyStatement, err := ConstructStatement(cparser, t)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		body = append(body, bodyStatement)
	}

	return statement.CreateStatementGroup(body), nil
}

func ConstructStatementIf(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (statement.AwooParserStatement, error) {
	n := ConstructExpressionStart(cparser, &ConstructExpressionDetails{EndWithCurly: true})
	if n.Error != nil {
		return statement.AwooParserStatement{}, n.Error
	}
	ifStatement := statement.CreateStatementIf(n.Node)
	ifGroup, err := ConstructStatementIfBody(cparser)
	if err != nil {
		return ifStatement, err
	}
	statement.SetStatementIfBody(&ifStatement, ifGroup)
	for t, ok := parser.PeekParser(cparser); ok && t.Type == token.TokenTypeElse; t, ok = parser.PeekParser(cparser) {
		t, _ = parser.FetchTokenParser(cparser)
		t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeBracketCurlyLeft}, "{")
		if err != nil {
			return ifStatement, err
		}
		elseGroup, err := ConstructStatementIfBody(cparser)
		if err != nil {
			return ifStatement, err
		}
		statement.SetStatementIfNext(&ifStatement, append(statement.GetStatementIfNext(&ifStatement), elseGroup))
	}

	return ifStatement, nil
}
