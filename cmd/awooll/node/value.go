package node

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

func CreateNodeValue(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, primitiveType types.AwooType) (AwooParserNode, error) {
	switch t.Type {
	case token.TokenOperatorSubstraction:
		// TODO: Move this logic so it can be applied to expressions
		n, err := CreateNodeValueFast(context, fetchToken, primitiveType)
		if err != nil {
			return n, err
		}
		return CreateNodeNegative(t, n), nil
	case token.TokenTypePrimitive:
		return CreateNodePrimitiveSafe(context, t)
	case token.TokenTypeIdentifier:
		return CreateNodeIdentifierSafe(context, t)
	}

	return AwooParserNode{}, fmt.Errorf("expected a %s", gchalk.Red("primitive or identifier"))
}

func CreateNodeValueFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, primitiveType types.AwooType) (AwooParserNode, error) {
	t, err := fetchToken()
	if err != nil {
		return AwooParserNode{}, err
	}
	return CreateNodeValue(context, t, fetchToken, primitiveType)
}
