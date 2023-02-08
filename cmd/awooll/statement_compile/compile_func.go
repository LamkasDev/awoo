package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

func CompileStatementFunc(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	// TODO: populate arguments into context
	d, err := CompileStatementGroup(context, statement.GetStatementFuncBody(&s), d)
	if err != nil {
		return d, err
	}
	idNode := statement.GetStatementFuncIdentifier(&s)
	compiler_context.PushCompilerFunction(context, compiler_context.AwooCompilerFunction{
		Name:      node.GetNodeIdentifierValue(&idNode),
		Start:     compiler_context.GetProgramHeaderSize() + context.Position,
		Size:      uint16(len(d)),
		Arguments: statement.GetStatementFuncArguments(&s),
	})

	return d, nil
}
