package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func CompileStatement(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement) ([]byte, error) {
	switch s.Type {
	case statement.ParserStatementTypeDefinitionVariable:
		return CompileStatementDefinition(context, s)
	case statement.ParserStatementTypeAssignment:
		return CompileStatementAssignment(context, s)
	case statement.ParserStatementTypeDefinitionType:
		return []byte{}, nil
	}

	return []byte{}, fmt.Errorf("no idea how to compile statement %s", gchalk.Red(fmt.Sprint(s.Type)))
}
