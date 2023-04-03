package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructExpressionUnary(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, err
	}
	if leftNode.Node.Type == node.ParserNodeTypeExpression &&
		!node.GetNodeExpressionIsBracket(&leftNode.Node) && token.DoesTokenTakePrecendence(op.Type, leftNode.Node.Token.Type) {
		n := node.CreateNodeExpression(op, node.GetNodeExpressionRight(&leftNode.Node), rightNode.Node)
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(leftNode.Node.Token, node.GetNodeExpressionLeft(&leftNode.Node), n),
		}, nil
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
	}, nil
}
