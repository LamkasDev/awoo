package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func ConstructStatement(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	entry, ok := cparser.Settings.Mappings.Statement[t.Type]
	if !ok {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorCantParseStatement, gchalk.Red(fmt.Sprintf("%#x", t.Type)))
	}

	return entry(cparser, t, details)
}
