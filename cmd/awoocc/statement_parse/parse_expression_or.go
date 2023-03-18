package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
)

func ConstructExpressionOr(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, err
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
	}, nil
}
