package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction_helper"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementCall(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	return CompileNodeCall(ccompiler, celf, statement.GetStatementCallNode(&s), &details)
}

func CompileNodeCall(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	functionName := node.GetNodeCallValue(&n)
	function, ok := elf.GetSymbol(celf, functionName)
	if !ok {
		return awerrors.ErrorFailedToGetFunctionFromScope
	}
	stackOffset := scope.GetCurrentFunctionSize(&ccompiler.Context.Scopes)

	functionPrototypeArguments := function.Details.(elf.AwooElfSymbolTableEntryFunctionDetails).Arguments
	functionArguments := node.GetNodeCallArguments(&n)
	functionArgumentsOffset := stackOffset + cc.AwooCompilerReturnAddressSize
	for i := 0; i < len(functionPrototypeArguments); i++ {
		argumentDetails := compiler_details.CompileNodeValueDetails{
			Type:     functionPrototypeArguments[i].Type,
			Register: cpu.AwooRegisterTemporaryZero,
		}
		if err := CompileNodeValue(ccompiler, celf, functionArguments[i], &argumentDetails); err != nil {
			return err
		}
		saveInstruction := instruction.AwooInstruction{
			Definition: *instructions.AwooInstructionsSave[functionPrototypeArguments[i].Size],
			SourceOne:  cpu.AwooRegisterSavedZero,
			SourceTwo:  argumentDetails.Register,
			Immediate:  functionArgumentsOffset,
		}
		if err := encoder.Encode(celf, saveInstruction); err != nil {
			return err
		}
		functionArgumentsOffset += functionPrototypeArguments[i].Size
	}

	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionAdjustStack(stackOffset)); err != nil {
		return err
	}

	details.Register = cpu.AwooRegisterFunctionZero
	jumpInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
	}
	elf.PushRelocationEntry(celf, function.Name)
	if err := encoder.Encode(celf, jumpInstruction); err != nil {
		return err
	}

	return encoder.Encode(celf, instruction_helper.ConstructInstructionAdjustStack(-stackOffset))
}
