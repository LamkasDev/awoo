package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementIf(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserStatement, error) {
	n := ConstructExpressionStart(context, fetchToken, &ConstructExpressionDetails{EndWithCurly: true})
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	statement := CreateStatementIf(n.Node)
	body := []AwooParserStatement{}
	for t, err := fetchToken(); err == nil && t.Type != token.TokenTypeBracketCurlyRight; t, err = fetchToken() {
		bodyStatement, err := ConstructStatement(context, t, fetchToken)
		if err != nil {
			return statement, err
		}
		body = append(body, bodyStatement)
	}
	SetStatementIfBody(&statement, body)

	return statement, nil
}
