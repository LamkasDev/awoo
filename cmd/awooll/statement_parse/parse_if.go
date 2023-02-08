package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementIfOuter(cparser *parser.AwooParser, details *ConstructStatementDetails) (statement.AwooParserStatement, error) {
	n, err := ConstructExpressionStart(cparser, &ConstructExpressionDetails{EndWithCurly: true})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	ifStatement := statement.CreateStatementIf(n.Node)
	ifGroup, err := ConstructStatementGroup(cparser, details)
	if err != nil {
		return ifStatement, err
	}
	statement.SetStatementIfBody(&ifStatement, ifGroup)

	return ifStatement, nil
}

func ConstructStatementIf(cparser *parser.AwooParser, details *ConstructStatementDetails) (statement.AwooParserStatement, error) {
	ifStatement, err := ConstructStatementIfOuter(cparser, details)
	if err != nil {
		return ifStatement, err
	}
	for t, ok := parser.PeekParser(cparser); ok && t.Type == token.TokenTypeElse; t, ok = parser.PeekParser(cparser) {
		t, _ = parser.FetchTokenParser(cparser)
		t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeIf, token.TokenTypeBracketCurlyLeft}, "if or {")
		if err != nil {
			return ifStatement, err
		}
		switch t.Type {
		case token.TokenTypeIf:
			elifStatement, err := ConstructStatementIfOuter(cparser, details)
			if err != nil {
				return ifStatement, err
			}
			statement.SetStatementIfNext(&ifStatement, append(statement.GetStatementIfNext(&ifStatement), elifStatement))
		case token.TokenTypeBracketCurlyLeft:
			elseStatement, err := ConstructStatementGroup(cparser, details)
			if err != nil {
				return ifStatement, err
			}
			statement.SetStatementIfNext(&ifStatement, append(statement.GetStatementIfNext(&ifStatement), elseStatement))
		}
	}

	return ifStatement, nil
}
