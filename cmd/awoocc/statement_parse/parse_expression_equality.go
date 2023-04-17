package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructExpressionEquality(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	op, err := parser.ExpectToken(cparser, token.TokenOperatorEq)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, err
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
			Type:     token.TokenOperatorEqEq,
			Position: lexer_token.ExtendAwooLexerTokenPosition(t.Position, op.Position),
		}, leftNode.Node, rightNode.Node),
	}, nil
}

func ConstructExpressionNotEquality(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	op, err := parser.ExpectToken(cparser, token.TokenOperatorEq)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	rightNode, err := ConstructExpressionReferenceFast(cparser, details)
	if err != nil {
		return rightNode, err
	}
	return node.AwooParserNodeResult{
		Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
			Type:     token.TokenOperatorNotEq,
			Position: lexer_token.ExtendAwooLexerTokenPosition(t.Position, op.Position),
		}, leftNode.Node, rightNode.Node),
	}, nil
}
