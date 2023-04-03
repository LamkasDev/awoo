package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructExpressionReference(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	switch t.Type {
	case token.TokenOperatorDereference:
		n, err := CreateNodeIdentifierVariableSafeFast(cparser)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		return node.CreateNodeDereference(t, n.Node), nil
	case token.TokenOperatorReference:
		n, err := CreateNodeIdentifierVariableSafeFast(cparser)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		return node.CreateNodeReference(t, n.Node), nil
	}
	return ConstructExpressionNegative(cparser, t, details)
}

func ConstructExpressionReferenceFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketCurlyLeft, token.TokenTypeBracketLeft, token.TokenOperatorSubstraction, token.TokenOperatorDereference, token.TokenOperatorReference})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructExpressionReference(cparser, t, details)
}
