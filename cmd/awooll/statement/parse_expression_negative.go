package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionNegative(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	// TODO: negating brackets negates the whole subsequent expression
	if t.Type == token.TokenOperatorSubstraction {
		details.Negative++
		return ConstructExpressionNegativeFast(context, fetchToken, details)
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
