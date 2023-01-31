package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionPriority(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if t.Type == token.TokenTypeBracketLeft {
		details.PendingBrackets++
		return ConstructExpressionBracketFast(context, fetchToken, details)
	}
	return ConstructNodeValue(context, t, fetchToken, details)
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
