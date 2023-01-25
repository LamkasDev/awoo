package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeNegative(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeValue(context, node.GetNodeNegativeValue(&n), d, details)
	if err != nil {
		return d, err
	}
	if details.Expression {
		return encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionSUB,
			Destination: cpu.AwooRegisterTemporaryOne,
			SourceTwo:   cpu.AwooRegisterTemporaryOne,
		}, d)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSUB,
		Destination: cpu.AwooRegisterTemporaryZero,
		SourceTwo:   cpu.AwooRegisterTemporaryZero,
	}, d)
}
