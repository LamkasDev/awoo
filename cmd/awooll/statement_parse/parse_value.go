package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/jwalton/gchalk"
)

func ConstructNodeValue(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	entry, ok := cparser.Settings.Mappings.NodeValue[t.Type]
	if !ok {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red("primitive or identifier"))
	}

	return entry(cparser, t)
}

func ConstructNodeValueFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchToken(cparser)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructNodeValue(cparser, t)
}
