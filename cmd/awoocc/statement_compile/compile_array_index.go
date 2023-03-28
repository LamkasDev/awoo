package statement_compile

import (
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileArrayIndexAddress(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, addressDetails *compiler_details.CompileNodeValueDetails) error {
	if err := CompileNodeValue(ccompiler, elf, node.GetNodeArrayIndexIndex(&n), addressDetails); err != nil {
		return err
	}
	// TODO: add a method for sizes that are not power of 2
	return encoder.Encode(elf, instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionSLLI,
		SourceOne:   addressDetails.Register,
		Destination: addressDetails.Register,
		Immediate:   arch.AwooRegister(math.Log((float64)(ccompiler.Context.Parser.Lexer.Types.All[addressDetails.Type].Size)) / math.Log(2)),
	})
}

func CompileNodeArrayIndex(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeArrayIndexIdentifier(&n))
	if err != nil {
		return err
	}

	addressDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.GetNextTemporaryRegister(details.Register),
	}
	if err = CompileArrayIndexAddress(ccompiler, elf, n, &addressDetails); err != nil {
		return err
	}
	if !variableMemory.Global {
		addressAdjustmentInstruction := instruction.AwooInstruction{
			Definition:  instructions.AwooInstructionADD,
			SourceOne:   addressDetails.Register,
			SourceTwo:   cpu.AwooRegisterSavedZero,
			Destination: addressDetails.Register,
		}
		if err = encoder.Encode(elf, addressAdjustmentInstruction); err != nil {
			return err
		}
	}

	loadInstruction := instruction.AwooInstruction{
		Definition:  *instructions.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type].Size],
		SourceOne:   addressDetails.Register,
		Destination: details.Register,
		Immediate:   arch.AwooRegister(variableMemory.Symbol.Start),
	}
	return encoder.Encode(elf, loadInstruction)
}
