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
	d := make([]byte, 8)
	nameNode := statement.GetStatementAssignmentIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	primNode := statement.GetStatementAssignmentValue(&s)
	prim := node.GetNodePrimitiveValue(&primNode)
	dest, _ := compiler_context.GetContextMemory(context, name)
	err := encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		Destination: cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(prim.(int)),
	}, d)
	if err != nil {
		return d, err
	}
	err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSW,
		SourceTwo:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(dest),
	}, d[4:])
	return d, err
}
