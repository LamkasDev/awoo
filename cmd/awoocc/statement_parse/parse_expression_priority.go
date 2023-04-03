package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructExpressionPriority(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	if t.Type == token.TokenTypeBracketLeft {
		details.PendingBrackets++
		return ConstructExpressionBracketFast(cparser, details)
	}
	return ConstructNodeValue(cparser, t, details)
}

func ConstructExpressionPriorityFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketCurlyLeft, token.TokenTypeBracketLeft})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructExpressionPriority(cparser, t, details)
}
