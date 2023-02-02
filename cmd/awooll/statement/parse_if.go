package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
)

func ConstructStatementIf(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserStatement, error) {
	n := ConstructExpressionStart(context, fetchToken, &ConstructExpressionDetails{})
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	statement := CreateStatementIf(n.Node)

	return statement, nil
}
