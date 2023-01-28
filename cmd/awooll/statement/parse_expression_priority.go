package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

// TODO: handle double brackets -> ((1 + 2))
func ConstructExpressionPriority(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if t.Type == token.TokenTypeBracketLeft {
		details.Bracket++
		if details.Value > 0 {
			details.Value--
		}
		if details.Negative > 0 {
			details.Negative--
			n := ConstructExpressionNegativeFast(context, fetchToken, details)
			if n.Error != nil {
				return n
			}
			// TODO: this is missing token
			return node.CreateNodeNegative(lexer_token.AwooLexerToken{}, n.Node)
		}
		return ConstructExpressionNegativeFast(context, fetchToken, details)
	}
	return ConstructExpression(context, t, fetchToken, details)
}

func ConstructExpressionPriorityFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketLeft}, "primitive, identifier or (")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructExpressionPriority(context, t, fetchToken, details)
}
