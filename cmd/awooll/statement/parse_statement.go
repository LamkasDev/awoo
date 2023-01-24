package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

type ConstructStatementFetchToken func() (lexer_token.AwooLexerToken, error)

func ExpectToken(fetchToken ConstructStatementFetchToken, tokenTypes []uint16, tokenName string) (lexer_token.AwooLexerToken, error) {
	t, err := fetchToken()
	if err != nil {
		return t, err
	}
	if util.Contains(tokenTypes, t.Type) {
		return t, fmt.Errorf("expected a %s", gchalk.Red(tokenName))
	}

	return t, nil
}

func ConstructStatement(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken ConstructStatementFetchToken) (AwooParserStatement, error) {
	statement, err := AwooParserStatement{}, fmt.Errorf("expected a %s", gchalk.Red("statement"))
	switch t.Type {
	case token.TokenTypeType:
		statement, err = ConstructStatementDefinition(context, t, fetchToken)
	case token.TokenTypeIdentifier:
		identifier := lexer_token.GetTokenIdentifierValue(&t)
		if _, ok := parser_context.GetContextVariable(context, identifier); !ok {
			return statement, fmt.Errorf("unknown identifier %s", gchalk.Red(identifier))
		}
		statement, err = ConstructStatementAssignment(context, t, fetchToken)
	}

	return statement, err
}
