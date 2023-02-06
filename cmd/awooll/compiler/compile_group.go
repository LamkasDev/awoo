package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

func CompileStatementGroup(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	var err error
	for _, n := range statement.GetStatementGroupBody(&s) {
		d, err = CompileStatement(context, n, d)
		if err != nil {
			return d, err
		}
	}
	return d, nil
}
