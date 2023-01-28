package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

func ExpectToken(fetchToken lexer_token.FetchToken, tokenTypes []uint16, tokenName string) (lexer_token.AwooLexerToken, error) {
	t, err := fetchToken()
	if err != nil {
		return t, err
	}
	if !util.Contains(tokenTypes, t.Type) {
		return t, fmt.Errorf("expected a %s", gchalk.Red(tokenName))
	}

	return t, nil
}

/*
Flow of functions
1) ConstructExpressionNegative (handles negatives)
2) ConstructExpressionPriority (handles brackets)
3) ConstructExpression (handles left side of expressions) & ConstructExpressionContinue (joins left and right side of expression)
-> ConstructExpressionNegativeFast & ConstructExpressionContinue
*/
func ConstructStatement(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserStatement, error) {
	statement, err := AwooParserStatement{}, fmt.Errorf("expected a %s", gchalk.Red("statement"))
	switch t.Type {
	case token.TokenTypeType:
		statement, err = ConstructStatementDefinition(context, t, fetchToken)
	case token.TokenTypeIdentifier:
		statement, err = ConstructStatementAssignment(context, t, fetchToken)
	}

	return statement, err
}
