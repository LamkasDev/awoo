package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodePrimitiveFirst(context *compiler_context.AwooCompilerContext, prim interface{}, d []byte) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		Immediate:   uint32(prim.(int64)),
		Destination: cpu.AwooRegisterTemporaryZero,
	}, d)
}

func CompileNodePrimitive(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	prim := node.GetNodePrimitiveValue(&n)
	if details.First {
		return CompileNodePrimitiveFirst(context, prim, d)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(prim.(int64)),
		Destination: cpu.AwooRegisterTemporaryZero,
	}, d)
}
