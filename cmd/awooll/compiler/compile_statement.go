package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func CompileStatement(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	switch s.Type {
	case statement.ParserStatementTypeDefinitionVariable:
		return CompileStatementDefinition(context, s, d)
	case statement.ParserStatementTypeAssignment:
		return CompileStatementAssignment(context, s, d)
	case statement.ParserStatementTypeDefinitionType:
		return []byte{}, nil
	case statement.ParserStatementTypeIf:
		return CompileStatementIf(context, s, d)
	case statement.ParserStatementTypeGroup:
		return CompileStatementGroup(context, s, d)
	}

	return []byte{}, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileStatement, gchalk.Red(fmt.Sprint(s.Type)))
}
