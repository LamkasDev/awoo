package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionReference(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	switch t.Type {
	case token.TokenOperatorMultiplication:
		n, err := CreateNodeIdentifierVariableSafeFast(cparser)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		return node.CreateNodeDereference(t, n.Node), nil
	case token.TokenTypeReference:
		n, err := CreateNodeIdentifierVariableSafeFast(cparser)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		return node.CreateNodeReference(t, n.Node), nil
	}
	return ConstructExpressionNegative(cparser, t, details)
}

func ConstructExpressionReferenceFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokensParser(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenOperatorEq, token.TokenOperatorLT, token.TokenOperatorGT, token.TokenTypeBracketLeft, token.TokenOperatorSubstraction, token.TokenOperatorMultiplication, token.TokenTypeReference}, "primitive, identifier, =, <, >, (, -, * or &")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructExpressionReference(cparser, t, details)
}
