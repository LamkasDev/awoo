package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructNodeValue(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	var n node.AwooParserNodeResult
	switch t.Type {
	case token.TokenTypePrimitive:
		n = node.CreateNodePrimitiveSafe(context, t)
	case token.TokenTypeIdentifier:
		n = node.CreateNodeIdentifierSafe(context, t)
	default:
		return node.AwooParserNodeResult{
			Error: fmt.Errorf("expected a %s", gchalk.Red("primitive or identifier")),
		}
	}
	if n.Error != nil {
		return n
	}
	if details.Negative > 0 {
		details.Negative--
		// TODO: this is missing token
		return node.CreateNodeNegative(lexer_token.AwooLexerToken{}, n.Node)
	}

	return n
}

func ConstructNodeValueFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier}, "primitive or identifier")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructNodeValue(context, t, fetchToken, details)
}
