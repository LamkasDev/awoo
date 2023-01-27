package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

type ConstructExpressionDetails struct {
	Type    types.AwooType
	Bracket uint8
}

type ConstructExpressionResult struct {
	Node  node.AwooParserNode
	Error error
	End   bool
}

func ConstructExpression(context *parser_context.AwooParserContext, leftNode node.AwooParserNode, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) ConstructExpressionResult {
	op, err := fetchToken()
	if err != nil {
		return ConstructExpressionResult{
			Node: leftNode,
			End:  true,
		}
	}
	switch op.Type {
	case token.TokenTypeEndStatement:
		if details.Bracket > 0 {
			return ConstructExpressionResult{
				Node:  leftNode,
				Error: fmt.Errorf("expected a %s", gchalk.Red(")")),
			}
		}
		return ConstructExpressionResult{
			Node: leftNode,
			End:  true,
		}
	case token.TokenTypeBracketRight:
		if details.Bracket > 0 {
			details.Bracket--
			return ConstructExpressionResult{
				Node: leftNode,
			}
		}
		return ConstructExpressionResult{
			Node:  leftNode,
			Error: fmt.Errorf("unexpected %s", gchalk.Red(")")),
		}
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction:
		rightNode := ConstructExpressionPriorityFast(context, fetchToken, details)
		if rightNode.Error != nil {
			return rightNode
		}
		leftNode = node.CreateNodeExpression(op, leftNode, rightNode.Node)
		if rightNode.End {
			return ConstructExpressionResult{
				Node: leftNode,
				End:  true,
			}
		}
		return ConstructExpression(context, leftNode, fetchToken, details)
	case token.TokenOperatorMultiplication,
		token.TokenOperatorDivision:
		// get a singular right value (or bracket expression)
		rightNode := CreateNodeValuePriorityFast(context, fetchToken, details)
		if rightNode.Error != nil {
			return rightNode
		}
		// join the two so they cannot be separated
		leftNode := node.CreateNodeExpression(op, leftNode, rightNode.Node)
		if rightNode.End {
			return ConstructExpressionResult{
				Node: leftNode,
				End:  true,
			}
		}
		return ConstructExpression(context, leftNode, fetchToken, details)
	}

	return ConstructExpressionResult{
		Node:  leftNode,
		Error: fmt.Errorf("expected an %s", gchalk.Red("operator or ;")),
	}
}

func ConstructExpressionFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) ConstructExpressionResult {
	leftNode := ConstructExpressionPriorityFast(context, fetchToken, details)
	if leftNode.Error != nil || leftNode.End {
		return leftNode
	}
	return ConstructExpression(context, leftNode.Node, fetchToken, details)
}
