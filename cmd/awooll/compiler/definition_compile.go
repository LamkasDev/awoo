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
	d := []byte{}
	tNode := statement.GetStatementDefinitionVariableType(&s)
	t := node.GetNodeTypeType(&tNode)
	nameNode := statement.GetStatementDefinitionVariableIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	dest, err := compiler_context.SetCompilerScopeCurrentMemory(context, name, t)
	if err != nil {
		return d, err
	}
	d, err = CompileNodeValueFast(context, valueNode, d)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSW,
		SourceTwo:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(dest),
	}, d)
}
