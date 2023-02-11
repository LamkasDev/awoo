package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodePrimitive(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	primitiveValue := node.GetNodePrimitiveValue(&n).(int64)
	if primitiveValue > arch.AwooImmediateSmallMax {
		// TODO: this will get fucked by sign extension.
		nextRegister := cpu.GetNextTemporaryRegister(details.Register)
		d, err := encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionLUI,
			Immediate:   uint32(primitiveValue),
			Destination: nextRegister,
		}, d)
		if err != nil {
			return d, err
		}
		return encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionADDI,
			Immediate:   uint32(primitiveValue),
			SourceOne:   nextRegister,
			Destination: details.Register,
		}, d)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		Immediate:   uint32(primitiveValue),
		Destination: details.Register,
	}, d)
}
