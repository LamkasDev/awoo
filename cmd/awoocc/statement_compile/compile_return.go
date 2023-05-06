package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction_helper"
)

func CompileStatementReturn(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	currentScopeFunction := ccompiler.Context.Scopes.Functions[uint16(len(ccompiler.Context.Scopes.Functions)-1)]
	currentPrototypeFunction, ok := elf.GetSymbol(celf, currentScopeFunction.Name)
	if !ok {
		return awerrors.ErrorFailedToGetFunctionFromScope
	}

	currentReturnType := currentPrototypeFunction.Details.(elf.AwooElfSymbolTableEntryFunctionDetails).ReturnType
	if currentReturnType != nil {
		returnValueNode := statement.GetStatementReturnValue(&s)
		returnDetails := compiler_details.CompileNodeValueDetails{
			Type:     *currentReturnType,
			Register: cpu.AwooRegisterFunctionZero,
		}
		if err := CompileNodeValue(ccompiler, celf, *returnValueNode, &returnDetails); err != nil {
			return err
		}
	}

	loadReturnAddressInstruction := instruction_helper.ConstructInstructionLoadReturnAddress()
	if err := encoder.Encode(celf, loadReturnAddressInstruction); err != nil {
		return err
	}

	return encoder.Encode(celf, instruction_helper.ConstructInstructionJumpToReturnAddress())
}
