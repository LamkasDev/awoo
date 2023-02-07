package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionComparison(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchTokenParser(cparser)
	if err != nil {
		return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	if t.Type == token.TokenOperatorEq {
		if op.Type == token.TokenOperatorLT {
			op.Type = token.TokenOperatorLTEQ
		} else {
			op.Type = token.TokenOperatorGTEQ
		}
		t, err = parser.FetchTokenParser(cparser)
		if err != nil {
			return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
	}
	rightNode, err := ConstructExpressionReference(cparser, t, details)
	if err != nil {
		return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
	}, nil
}
