package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementAssignment(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement) ([]byte, error) {
	d := []byte{}
	nameNode := statement.GetStatementAssignmentIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	dest, _ := compiler_context.GetCompilerScopeMemory(context, name)
	valueNode := statement.GetStatementAssignmentValue(&s)
	d, err := CompileNodeValueFast(context, valueNode, d)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSW,
		SourceTwo:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(dest),
	}, d)
}
