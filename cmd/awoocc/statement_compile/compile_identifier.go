package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeIdentifier(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	identifier := node.GetNodeIdentifierValue(&n)
	variableMemory, err := compiler_context.GetCompilerScopeFunctionMemory(&ccompiler.Context, identifier)
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type]

	loadInstruction := instruction.AwooInstruction{
		Definition:  *instructions.AwooInstructionsLoad[variableType.Size],
		Destination: details.Register,
		Immediate:   variableMemory.Symbol.Start,
	}
	if !variableMemory.Global {
		loadInstruction.SourceOne = cpu.AwooRegisterSavedZero
	} else {
		elf.PushRelocationEntry(celf, variableMemory.Symbol.Name)
	}
	return encoder.Encode(celf, loadInstruction)
}
