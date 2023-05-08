package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
)

func ConstructStatementImport(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	t, err := parser.ExpectToken(cparser, token.TokenTypePrimitive)
	if err != nil {
		return nil, err
	}
	if lexer_token.GetTokenPrimitiveType(t) != types.AwooTypeString {
		return nil, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red("string")),
			t.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
	}
	if _, err = parser.ExpectToken(cparser, details.EndToken); err != nil {
		return nil, err
	}

	return nil, nil
}
