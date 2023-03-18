package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeArray(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	addressAdjustmentInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		SourceOne:   details.Address.Register,
		Immediate:   details.Address.Immediate,
		Destination: details.Register,
	}
	var err error
	if d, err = encoder.Encode(addressAdjustmentInstruction, d); err != nil {
		return d, err
	}

	for i, elementNode := range node.GetNodeArrayElements(&n) {
		if i != 0 {
			addressAdjustmentInstruction = encoder.AwooEncodedInstruction{
				Instruction: instructions.AwooInstructionADDI,
				SourceOne:   details.Register,
				Immediate:   uint32(ccompiler.Context.Parser.Lexer.Types.All[details.Type].Size),
				Destination: details.Register,
			}
			if d, err = encoder.Encode(addressAdjustmentInstruction, d); err != nil {
				return d, err
			}
		}

		elementType := ccompiler.Context.Parser.Lexer.Types.All[details.Type]
		elementDetails := compiler_details.CompileNodeValueDetails{
			Type:     details.Type,
			Register: cpu.GetNextTemporaryRegister(details.Register),
		}
		if d, err = CompileNodeValue(ccompiler, elementNode, d, &elementDetails); err != nil {
			return d, err
		}

		saveInstruction := encoder.AwooEncodedInstruction{
			Instruction: *instructions.AwooInstructionsSave[elementType.Size],
			SourceOne:   details.Register,
			SourceTwo:   elementDetails.Register,
		}
		if d, err = encoder.Encode(saveInstruction, d); err != nil {
			return d, err
		}
	}

	details.Address.Used = true
	return d, nil
}
