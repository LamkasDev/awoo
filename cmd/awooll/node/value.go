package node

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

func CreateNodeValue(context *parser_context.AwooParserContext, primitiveType types.AwooType, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserNode, error) {
	switch t.Type {
	case token.TokenTypePrimitive:
		return CreateNodePrimitiveSafe(primitiveType, t)
	case token.TokenTypeIdentifier:
		return CreateNodeIdentifierSafe(context, t)
	}

	return AwooParserNode{}, fmt.Errorf("expected a %s", gchalk.Red("primitive or identifier"))
}

func CreateNodeValueFast(context *parser_context.AwooParserContext, primitiveType types.AwooType, fetchToken lexer_token.FetchToken) (AwooParserNode, error) {
	t, err := fetchToken()
	if err != nil {
		return AwooParserNode{}, err
	}
	return CreateNodeValue(context, primitiveType, t, fetchToken)
}
