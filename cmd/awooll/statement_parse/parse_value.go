package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructNodeValue(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	var n node.AwooParserNodeResult
	switch t.Type {
	case token.TokenTypePrimitive:
		n = node.CreateNodePrimitiveSafe(&cparser.Context, t)
	case token.TokenTypeIdentifier:
		n = node.CreateNodeIdentifierSafe(&cparser.Context, t)
	default:
		return node.AwooParserNodeResult{
			Error: fmt.Errorf("expected a %s", gchalk.Red("primitive or identifier")),
		}
	}
	if n.Error != nil {
		return n
	}

	return n
}

func ConstructNodeValueFast(cparser *parser.AwooParser, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier}, "primitive or identifier")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructNodeValue(cparser, t, details)
}
