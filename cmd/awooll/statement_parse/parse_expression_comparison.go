package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionComparison(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchToken(cparser)
	if err != nil {
		return leftNode, err
	}
	if t.Type == token.TokenOperatorEq {
		if op.Type == token.TokenOperatorLT {
			op.Type = token.TokenOperatorLTEQ
		} else {
			op.Type = token.TokenOperatorGTEQ
		}
		t, err = parser.FetchToken(cparser)
		if err != nil {
			return leftNode, err
		}
	}
	rightNode, err := ConstructExpressionReference(cparser, t, details)
	if err != nil {
		return rightNode, err
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
	}, nil
}
