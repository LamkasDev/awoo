package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodePrimitive(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	// TODO: handle unsigned.
	primitiveValue := node.GetNodePrimitiveValueFormat[int64](ccompiler.Context.Parser.Lexer, &n)
	if primitiveValue > arch.AwooImmediateSmallMax {
		luiValue := (primitiveValue >> 12) << 12
		addiValue := primitiveValue - luiValue
		if addiValue&0b1000_0000_0000 > 0 {
			luiValue += addiValue * 2
		}

		// Load upper 20-bits, if primitive value is over 12-bits
		d, err := encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionLUI,
			Immediate:   uint32(luiValue),
			Destination: details.Register,
		}, d)
		if err != nil {
			return d, err
		}

		if addiValue != 0 {
			// Load lower 12-bits
			return encoder.Encode(encoder.AwooEncodedInstruction{
				Instruction: instruction.AwooInstructionADDI,
				Immediate:   uint32(addiValue),
				SourceOne:   details.Register,
				Destination: details.Register,
			}, d)
		}

		return d, nil
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		Immediate:   uint32(primitiveValue),
		Destination: details.Register,
	}, d)
}
