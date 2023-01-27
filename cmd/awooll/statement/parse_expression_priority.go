package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionPriority(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) ConstructExpressionResult {
	priorityNode := ConstructExpressionPriorityFast(context, fetchToken, details)
	if priorityNode.Error != nil || priorityNode.End {
		return priorityNode
	}
	return ConstructExpression(context, priorityNode.Node, fetchToken, details)
}

func ConstructExpressionPriorityFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) ConstructExpressionResult {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketLeft}, "primitive, identifier or (")
	if err != nil {
		return ConstructExpressionResult{
			Error: err,
		}
	}
	if t.Type == token.TokenTypeBracketLeft {
		details.Bracket++
		return ConstructExpressionPriority(context, t, fetchToken, details)
	}
	valueNode, err := node.CreateNodeValue(context, t, fetchToken, details.Type)
	if err != nil {
		return ConstructExpressionResult{
			Error: err,
		}
	}
	return ConstructExpression(context, valueNode, fetchToken, details)
}

func CreateNodeValuePriorityFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) ConstructExpressionResult {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketLeft}, "primitive, identifier or (")
	if err != nil {
		return ConstructExpressionResult{
			Error: err,
		}
	}
	if t.Type == token.TokenTypeBracketLeft {
		details.Bracket++
		return ConstructExpressionPriority(context, t, fetchToken, details)
	}
	valueNode, err := node.CreateNodeValue(context, t, fetchToken, details.Type)
	return ConstructExpressionResult{
		Node:  valueNode,
		Error: err,
	}
}
