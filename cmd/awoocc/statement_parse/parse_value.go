package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/jwalton/gchalk"
)

func ConstructNodeValue(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	entry, ok := cparser.Settings.Mappings.NodeValue[t.Type]
	if !ok {
		panic(fmt.Errorf("%w: %s", awerrors.ErrorCantParseNode, gchalk.Red(fmt.Sprintf("%#x", t.Type))))
	}

	return entry(cparser, t, details)
}

func ConstructNodeValueFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.AdvanceParser(cparser)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructNodeValue(cparser, *t, details)
}
