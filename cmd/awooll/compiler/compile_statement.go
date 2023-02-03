package compiler

import (
	"fmt"

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
		var err error = nil
		for _, n := range statement.GetStatementGroupBody(&s) {
			d, err = CompileStatement(context, n, d)
			if err != nil {
				return d, err
			}
		}
		return d, nil
	}

	return []byte{}, fmt.Errorf("no idea how to compile statement %s", gchalk.Red(fmt.Sprint(s.Type)))
}
