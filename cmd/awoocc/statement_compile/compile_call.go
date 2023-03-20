package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementCall(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	return CompileNodeCall(ccompiler, elf, statement.GetStatementCallNode(&s), &details)
}

func CompileNodeCall(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	functionName := node.GetNodeCallValue(&n)
	function, ok := compiler_context.GetCompilerFunction(&ccompiler.Context, functionName)
	if !ok {
		return awerrors.ErrorFailedToGetFunctionFromScope
	}
	stackOffset := uint32(compiler_context.GetCompilerScopeCurrentFunctionSize(&ccompiler.Context))

	functionArguments := node.GetNodeCallArguments(&n)
	functionArgumentsOffset := stackOffset
	var err error
	for i := 0; i < len(function.Arguments); i++ {
		argumentDetails := compiler_details.CompileNodeValueDetails{
			Type:     function.Arguments[i].Type,
			Register: cpu.AwooRegisterTemporaryZero,
		}
		if err = CompileNodeValue(ccompiler, elf, functionArguments[i], &argumentDetails); err != nil {
			return err
		}
		saveInstruction := encoder.AwooEncodedInstruction{
			Instruction: *instructions.AwooInstructionsSave[function.Arguments[i].Size],
			SourceOne:   cpu.AwooRegisterSavedZero,
			SourceTwo:   argumentDetails.Register,
			Immediate:   functionArgumentsOffset,
		}
		if err = encoder.Encode(elf, saveInstruction); err != nil {
			return err
		}
		functionArgumentsOffset += uint32(function.Arguments[i].Size)
	}

	stackAdjustmentInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   stackOffset,
	}
	if err = encoder.Encode(elf, stackAdjustmentInstruction); err != nil {
		return err
	}

	details.Register = cpu.AwooRegisterFunctionZero
	jumpInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   uint32(function.Start),
	}
	if err = encoder.Encode(elf, jumpInstruction); err != nil {
		return err
	}

	stackAdjustmentInstruction = encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   -stackOffset,
	}
	return encoder.Encode(elf, stackAdjustmentInstruction)
}
