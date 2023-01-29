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
		details.QueuedBrackets++
		details.PendingBrackets++
		if details.Negative > 0 {
			details.NegativeBracket = details.Negative
			details.Negative = 0
		}
		return ConstructExpressionNegativeFast(context, fetchToken, details)
	}
	if details.QueuedBrackets > 0 {
		details.QueuedBrackets--
	}
	return ConstructExpressionContinue(context, t, fetchToken, details)
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
