package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeArray(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	addressAdjustmentInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionADDI,
		SourceOne:   details.Address.Register,
		Immediate:   details.Address.Immediate,
		Destination: details.Register,
	}
	if err := encoder.Encode(celf, addressAdjustmentInstruction); err != nil {
		return err
	}

	for i, elementNode := range node.GetNodeArrayElements(&n) {
		if i != 0 {
			addressAdjustmentInstruction = instruction.AwooInstruction{
				Definition:  instructions.AwooInstructionADDI,
				SourceOne:   details.Register,
				Immediate:   arch.AwooRegister(ccompiler.Context.Parser.Lexer.Types.All[details.Type].Size),
				Destination: details.Register,
			}
			if err := encoder.Encode(celf, addressAdjustmentInstruction); err != nil {
				return err
			}
		}

		elementType := ccompiler.Context.Parser.Lexer.Types.All[details.Type]
		elementDetails := compiler_details.CompileNodeValueDetails{
			Type:     details.Type,
			Register: cpu.GetNextTemporaryRegister(details.Register),
		}
		if err := CompileNodeValue(ccompiler, celf, elementNode, &elementDetails); err != nil {
			return err
		}

		saveInstruction := instruction.AwooInstruction{
			Definition: *instructions.AwooInstructionsSave[elementType.Size],
			SourceOne:  details.Register,
			SourceTwo:  elementDetails.Register,
		}
		if err := encoder.Encode(celf, saveInstruction); err != nil {
			return err
		}
	}

	details.Address.Used = true
	return nil
}
