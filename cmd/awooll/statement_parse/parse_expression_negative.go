package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionNegative(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if t.Type == token.TokenOperatorSubstraction {
		n, err := ConstructExpressionReferenceFast(cparser, details)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		return node.CreateNodeNegative(t, n.Node), nil
	}
	return ConstructExpressionPriority(cparser, t, details)
}

func ConstructExpressionNegativeFast(cparser *parser.AwooParser, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenOperatorEq, token.TokenOperatorLT, token.TokenOperatorGT, token.TokenTypeBracketLeft, token.TokenOperatorSubstraction}, "primitive, identifier, =, <, >, ( or -")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructExpressionNegative(cparser, t, details)
}
