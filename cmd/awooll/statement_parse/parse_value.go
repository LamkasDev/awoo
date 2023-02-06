package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructNodeValue(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	switch t.Type {
	case token.TokenTypePrimitive:
		return node.CreateNodePrimitiveSafe(&cparser.Context, t)
	case token.TokenTypeIdentifier:
		return node.CreateNodeIdentifierSafe(&cparser.Context, t)
	}

	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red("primitive or identifier"))
}

func ConstructNodeValueFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier}, "primitive or identifier")
	if err != nil {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructNode, err)
	}
	return ConstructNodeValue(cparser, t)
}
