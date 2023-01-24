package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementDefinition(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement) ([]byte, error) {
	d := make([]byte, 8)
	tNode := statement.GetStatementDefinitionType(&s)
	t := node.GetNodeTypeType(&tNode)
	nameNode := statement.GetStatementDefinitionIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	primNode := statement.GetStatementDefinitionValue(&s)
	prim := node.GetNodePrimitiveValue(&primNode)
	dest := uint8(context.Memory.Position)
	if err := compiler_context.SetContextMemory(context, name, t); err != nil {
		return []byte{}, err
	}
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
