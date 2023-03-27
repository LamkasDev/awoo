package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementReturn(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	currentScopeFunction := ccompiler.Context.Scopes.Functions[uint16(len(ccompiler.Context.Scopes.Functions)-1)]
	currentFunction, _ := compiler_context.GetCompilerFunction(&ccompiler.Context, currentScopeFunction.Name)
	if currentFunction.Symbol.TypeDetails != nil {
		returnValueNode := statement.GetStatementReturnValue(&s)
		returnDetails := compiler_details.CompileNodeValueDetails{
			Type:     *currentFunction.Symbol.TypeDetails,
			Register: cpu.AwooRegisterFunctionZero,
		}
		if err := CompileNodeValue(ccompiler, elf, *returnValueNode, &returnDetails); err != nil {
			return err
		}
	}

	loadReturnAddressInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionLW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Immediate:   uint32(compiler_context.GetCompilerFunctionArgumentsSize(currentFunction)),
		Destination: cpu.AwooRegisterReturnAddress,
	}
	if err := encoder.Encode(elf, loadReturnAddressInstruction); err != nil {
		return err
	}

	return encoder.Encode(elf, encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJALR,
		SourceOne:   cpu.AwooRegisterReturnAddress,
	})
}
