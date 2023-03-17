package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/jwalton/gchalk"
)

func ConstructNodeValue(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	entry, ok := cparser.Settings.Mappings.NodeValue[t.Type]
	if !ok {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorCantParseNode, gchalk.Red(fmt.Sprintf("%#x", t.Type)))
	}

	return entry(cparser, t, details)
}

func ConstructNodeValueFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchToken(cparser)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructNodeValue(cparser, t, details)
}
