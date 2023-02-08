package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

type ConstructStatementDetails struct {
	CanReturn bool
}

// TODO: redo using map
func ConstructStatement(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructStatementDetails) (statement.AwooParserStatement, error) {
	statement, err := statement.AwooParserStatement{}, awerrors.ErrorExpectedStatement
	switch t.Type {
	case token.TokenTypeType:
		statement, err = ConstructStatementDefinitionVariable(cparser, t)
	case token.TokenTypeIdentifier:
		statement, err = ConstructStatementAssignment(cparser, t)
	case token.TokenTypeTypeDefinition:
		statement, err = ConstructStatementDefinitionType(cparser)
	case token.TokenTypeIf:
		statement, err = ConstructStatementIf(cparser, details)
	case token.TokenTypeFunc:
		statement, err = ConstructStatementFunc(cparser)
	case token.TokenTypeReturn:
		if !details.CanReturn {
			return statement, fmt.Errorf("%w: %s", awerrors.ErrorUnexpectedStatement, gchalk.Red("return"))
		}
		statement, err = ConstructStatementReturn(cparser)
		details.CanReturn = false
	}

	return statement, err
}
