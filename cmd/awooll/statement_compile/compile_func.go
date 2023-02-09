package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

func CompileStatementFunc(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	// TODO: populate arguments into context
	// -> there will need to be a special class for handling stack variables
	idNode := statement.GetStatementFuncIdentifier(&s)
	identifier := node.GetNodeIdentifierValue(&idNode)
	compiler_context.PushCompilerScopeFunction(context, identifier)
	d, err := CompileStatementGroup(context, statement.GetStatementFuncBody(&s), d)
	if err != nil {
		return d, err
	}
	compiler_context.PopCompilerScopeCurrentFunction(context)
	compiler_context.PushCompilerFunction(context, compiler_context.AwooCompilerFunction{
		Name:      identifier,
		Start:     compiler_context.GetProgramHeaderSize() + context.Position,
		Size:      uint16(len(d)),
		Arguments: statement.GetStatementFuncArguments(&s),
	})

	return d, nil
}
