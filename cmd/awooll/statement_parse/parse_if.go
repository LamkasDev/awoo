package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementIfOuter(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
	n, err := ConstructExpressionStart(cparser, &ConstructExpressionDetails{EndWithCurly: true})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	ifStatement := statement.CreateStatementIf(n.Node)
	ifGroup, err := ConstructStatementIfGroup(cparser)
	if err != nil {
		return ifStatement, err
	}
	statement.SetStatementIfBody(&ifStatement, ifGroup)

	return ifStatement, nil
}

func ConstructStatementIfGroup(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
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

func ConstructStatementIf(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
	ifStatement, err := ConstructStatementIfOuter(cparser)
	if err != nil {
		return ifStatement, err
	}
	for t, ok := parser.PeekParser(cparser); ok && t.Type == token.TokenTypeElse; t, ok = parser.PeekParser(cparser) {
		t, _ = parser.FetchTokenParser(cparser)
		t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeIf, token.TokenTypeBracketCurlyLeft}, "if or {")
		if err != nil {
			return ifStatement, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
		}
		switch t.Type {
		case token.TokenTypeIf:
			elifStatement, err := ConstructStatementIfOuter(cparser)
			if err != nil {
				return ifStatement, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
			}
			statement.SetStatementIfNext(&ifStatement, append(statement.GetStatementIfNext(&ifStatement), elifStatement))
		case token.TokenTypeBracketCurlyLeft:
			elseStatement, err := ConstructStatementIfGroup(cparser)
			if err != nil {
				return ifStatement, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
			}
			statement.SetStatementIfNext(&ifStatement, append(statement.GetStatementIfNext(&ifStatement), elseStatement))
		}
	}

	return ifStatement, nil
}
