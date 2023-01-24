package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func CompileStatement(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement) ([]byte, error) {
	d, err := []byte{}, fmt.Errorf("no idea how to compile %s", gchalk.Red(fmt.Sprint(s.Type)))
	switch s.Type {
	case statement.ParserStatementTypeDefinition:
		d, err = CompileStatementDefinition(context, s)
	case statement.ParserStatementTypeAssignment:
		d, err = CompileStatementAssignment(context, s)
	}

	return d, err
}
