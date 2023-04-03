package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/jwalton/gchalk"
)

func ConstructExpressionEndStatement(cparser *parser.AwooParser, n node.AwooParserNodeResult, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	if details.PendingBrackets > 0 {
		return node.AwooParserNodeResult{}, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red(")")),
			cparser.Position, 1, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
	}
	return node.AwooParserNodeResult{
		Node: n.Node,
		End:  &t.Type,
	}, nil
}

func ConstructExpressionEndBracket(cparser *parser.AwooParser, n node.AwooParserNodeResult, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	if details.PendingBrackets > 0 {
		details.PendingBrackets--
		return node.AwooParserNodeResult{
			Node: n.Node,
			End:  &t.Type,
		}, nil
	}
	return node.AwooParserNodeResult{}, parser_error.CreateParserErrorText(parser_error.AwooParserErrorUnexpectedToken,
		fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorUnexpectedToken], gchalk.Red(")")),
		cparser.Position, 1, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorUnexpectedToken])
}
