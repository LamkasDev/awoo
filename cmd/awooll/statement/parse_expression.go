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

func ConstructExpression(context *parser_context.AwooParserContext, leftNode node.AwooParserNode, fetchToken lexer_token.FetchToken, requiredType types.AwooType) (node.AwooParserNode, error) {
	for true {
		op, err := fetchToken()
		if err != nil {
			// Return a singular value node
			return leftNode, nil
		}
		if op.Type == token.TokenOperatorEndStatement {
			break
		}
		switch op.Type {
		case token.TokenOperatorAddition,
			token.TokenOperatorSubstraction:
			rightNode, err := ConstructExpressionFast(context, fetchToken, requiredType)
			if err != nil {
				return leftNode, err
			}
			return node.CreateNodeExpression(op, leftNode, rightNode), nil
		case token.TokenOperatorMultiplication,
			token.TokenOperatorDivision:
			// get right value
			rightNode, err := node.CreateNodeValueFast(context, requiredType, fetchToken)
			if err != nil {
				return leftNode, err
			}
			// join the two so they cannot be separated
			leftNode, err := node.CreateNodeExpression(op, leftNode, rightNode), nil
			if err != nil {
				return leftNode, err
			}
			// continue as normal
			return ConstructExpression(context, leftNode, fetchToken, requiredType)
		default:
			return leftNode, fmt.Errorf("expected an %s", gchalk.Red("operator or ;"))
		}
	}

	return leftNode, nil
}

func ConstructExpressionFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, requiredType types.AwooType) (node.AwooParserNode, error) {
	leftNode, err := node.CreateNodeValueFast(context, requiredType, fetchToken)
	if err != nil {
		return leftNode, err
	}
	return ConstructExpression(context, leftNode, fetchToken, requiredType)
}
