package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructExpressionNegative(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	if t.Type == token.TokenOperatorSubstraction {
		n, err := ConstructExpressionReferenceFast(cparser, details)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		return node.CreateNodeNegative(t, n.Node), nil
	}
	return ConstructExpressionPriority(cparser, t, details)
}

func ConstructExpressionNegativeFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketCurlyLeft, token.TokenTypeBracketLeft, token.TokenOperatorSubstraction})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructExpressionNegative(cparser, *t, details)
}
