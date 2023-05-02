package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/jwalton/gchalk"
)

func ConstructStatement(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	entry, ok := cparser.Settings.Mappings.Statement[t.Type]
	if !ok {
		panic(fmt.Errorf("%w: %s", awerrors.ErrorCantParseStatement, gchalk.Red(fmt.Sprintf("%#x", t.Type))))
	}

	return entry(cparser, t, details)
}
