package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/jwalton/gchalk"
)

func ConstructExpressionEndStatement(_ *parser.AwooParser, n node.AwooParserNodeResult, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if details.PendingBrackets > 0 {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(")"))
	}
	return node.AwooParserNodeResult{
		Node: n.Node,
		End:  &t.Type,
	}, nil
}

func ConstructExpressionEndBracket(_ *parser.AwooParser, n node.AwooParserNodeResult, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if details.PendingBrackets > 0 {
		details.PendingBrackets--
		return node.AwooParserNodeResult{
			Node: n.Node,
			End:  &t.Type,
		}, nil
	}
	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnexpectedToken, gchalk.Red(")"))
}