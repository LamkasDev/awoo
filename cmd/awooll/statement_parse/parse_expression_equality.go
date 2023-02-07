package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionEquality(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	op, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "==")
	if err != nil {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
			Type:  token.TokenOperatorEqEq,
			Start: op.Start - 1,
		}, leftNode.Node, rightNode.Node),
	}, nil
}

func ConstructExpressionNotEquality(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	op, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "!=")
	if err != nil {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
			Type:  token.TokenOperatorNotEq,
			Start: op.Start - 1,
		}, leftNode.Node, rightNode.Node),
	}, nil
}
