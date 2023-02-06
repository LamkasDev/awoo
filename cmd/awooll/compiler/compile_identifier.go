package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeIdentifier(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	id := node.GetNodeIdentifierValue(&n)
	src, err := compiler_context.GetCompilerScopeMemory(context, id)
	if err != nil {
		return d, err
	}

	// TODO: pick instruction based on src size in bytes
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionLW,
		Destination: details.Register,
		Immediate:   uint32(src.Start),
	}, d)
}
