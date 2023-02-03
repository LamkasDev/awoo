package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatement(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (statement.AwooParserStatement, error) {
	statement, err := statement.AwooParserStatement{}, fmt.Errorf("expected a %s", gchalk.Red("statement"))
	switch t.Type {
	case token.TokenTypeType:
		statement, err = ConstructStatementDefinitionVariable(cparser, t)
	case token.TokenTypeIdentifier:
		statement, err = ConstructStatementAssignment(cparser, t)
	case token.TokenTypeTypeDefinition:
		statement, err = ConstructStatementDefinitionType(cparser, t)
	case token.TokenTypeIf:
		statement, err = ConstructStatementIf(cparser, t)
	}

	return statement, err
}
