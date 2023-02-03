package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionPriority(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if t.Type == token.TokenTypeBracketLeft {
		details.PendingBrackets++
		return ConstructExpressionBracketFast(cparser, details)
	}
	return ConstructNodeValue(cparser, t, details)
}

func ConstructExpressionPriorityFast(cparser *parser.AwooParser, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketLeft}, "primitive, identifier or (")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructExpressionPriority(cparser, t, details)
}
