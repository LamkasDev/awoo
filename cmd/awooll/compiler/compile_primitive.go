package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodePrimitive(n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	prim := node.GetNodePrimitiveValue(&n).(int64)
	if prim > arch.AwooImmediateSmallMax {
		// TODO: this will get fucked by sign extension.
		r := cpu.GetNextTemporaryRegister(details.Register)
		d, err := encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionLUI,
			Immediate:   uint32(prim),
			Destination: r,
		}, d)
		if err != nil {
			return d, fmt.Errorf("%w: %w", awerrors.ErrorCantCompileNode, err)
		}
		return encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionADDI,
			Immediate:   uint32(prim),
			SourceOne:   r,
			Destination: details.Register,
		}, d)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		Immediate:   uint32(prim),
		Destination: details.Register,
	}, d)
}
