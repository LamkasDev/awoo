package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementAssignment(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	nameNode := statement.GetStatementAssignmentIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	dest, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(context, name)
	if err != nil {
		return d, err
	}
	valueNode := statement.GetStatementAssignmentValue(&s)
	details := compiler_context.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	d, err = CompileNodeValueFast(context, valueNode, d, &details)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[context.Parser.Lexer.Types.All[dest.Type].Size],
		SourceOne:   cpu.AwooRegisterSavedZero,
		SourceTwo:   details.Register,
		Immediate:   uint32(dest.Start),
	}, d)
}
