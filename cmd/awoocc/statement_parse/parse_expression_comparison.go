package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructExpressionComparison(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if t, _ := parser.ExpectTokensOptional(cparser, []uint16{token.TokenOperatorEq, op.Type}); t != nil {
		switch op.Type {
		case token.TokenOperatorLT:
			switch t.Type {
			case token.TokenOperatorEq:
				op.Type = token.TokenOperatorLTEQ
			case token.TokenOperatorLT:
				op.Type = token.TokenOperatorLS
			}
		case token.TokenOperatorGT:
			switch t.Type {
			case token.TokenOperatorEq:
				op.Type = token.TokenOperatorGTEQ
			case token.TokenOperatorGT:
				op.Type = token.TokenOperatorRS
			}
		}
	}
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, err
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
	}, nil
}
