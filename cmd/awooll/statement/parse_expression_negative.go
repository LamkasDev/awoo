package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionNegative(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if t.Type == token.TokenOperatorSubstraction {
		n := ConstructExpressionNegativeFast(context, fetchToken, details)
		if n.Error != nil {
			return node.AwooParserNodeResult{
				Error: n.Error,
			}
		}
		return node.CreateNodeNegative(t, n.Node)
	}
	return ConstructExpressionPriority(context, t, fetchToken, details)
}

func ConstructExpressionNegativeFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketLeft, token.TokenOperatorSubstraction}, "primitive, identifier, ( or -")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructExpressionNegative(context, t, fetchToken, details)
}
