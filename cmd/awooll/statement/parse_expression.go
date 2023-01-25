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

func ConstructExpression(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, requiredType types.AwooType) (node.AwooParserNode, error) {
	leftNode, err := node.CreateNodeValueFast(context, requiredType, fetchToken)
	if err != nil {
		return leftNode, err
	}
	for true {
		op, err := fetchToken()
		if err != nil {
			return leftNode, err
		}
		if op.Type == token.TokenOperatorEndStatement {
			break
		}
		switch op.Type {
		case token.TokenOperatorAddition,
			token.TokenOperatorSubstraction,
			token.TokenOperatorMultiplication,
			token.TokenOperatorDivision:
			rightNode, err := node.CreateNodeValueFast(context, requiredType, fetchToken)
			if err != nil {
				return rightNode, err
			}
			leftNode = node.CreateNodeExpression(op, leftNode, rightNode)
		default:
			return leftNode, fmt.Errorf("expected an %s", gchalk.Red("operator or ;"))
		}
	}

	return leftNode, nil
}
