package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodePrimitive(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	// TODO: handle unsigned.
	primitiveValue := node.GetNodePrimitiveValueFormat[int64](ccompiler.Context.Parser.Lexer, &n)
	if primitiveValue > arch.AwooImmediateSmallMax {
		luiValue := (primitiveValue >> 12) << 12
		addiValue := primitiveValue - luiValue
		if addiValue&0b1000_0000_0000 > 0 {
			luiValue += addiValue * 2
		}

		// Load upper 20-bits, if primitive value is over 12-bits
		err := encoder.Encode(elf, instruction.AwooInstruction{
			Definition:  instructions.AwooInstructionLUI,
			Immediate:   arch.AwooRegister(luiValue),
			Destination: details.Register,
		})
		if err != nil {
			return err
		}

		if addiValue != 0 {
			// Load lower 12-bits
			return encoder.Encode(elf, instruction.AwooInstruction{
				Definition:  instructions.AwooInstructionADDI,
				Immediate:   arch.AwooRegister(addiValue),
				SourceOne:   details.Register,
				Destination: details.Register,
			})
		}

		return nil
	}

	return encoder.Encode(elf, instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionADDI,
		Immediate:   arch.AwooRegister(primitiveValue),
		Destination: details.Register,
	})
}
