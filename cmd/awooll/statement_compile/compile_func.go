package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

func CompileStatementFunc(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	functionNameNode := statement.GetStatementFuncIdentifier(&s)
	functionName := node.GetNodeIdentifierValue(&functionNameNode)
	functionArguments := statement.GetStatementFuncArguments(&s)

	compiler_context.PushCompilerScopeFunction(&ccompiler.Context, functionName)
	for _, argument := range functionArguments {
		compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerContextMemoryEntry{
			Name: argument.Name,
			Size: argument.Size,
			Type: argument.Type,
			Data: argument.Data,
		})
	}
	d, err := CompileStatementGroup(ccompiler, statement.GetStatementFuncBody(&s), d)
	if err != nil {
		return d, err
	}
	compiler_context.PopCompilerScopeCurrentFunction(&ccompiler.Context)
	compiler_context.PushCompilerFunction(&ccompiler.Context, compiler_context.AwooCompilerFunction{
		Name:      functionName,
		Start:     compiler_context.GetProgramHeaderSize() + ccompiler.Context.CurrentAddress,
		Size:      uint16(len(d)),
		Arguments: statement.GetStatementFuncArguments(&s),
	})

	return d, nil
}
