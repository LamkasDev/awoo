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
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementCall(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	return CompileNodeCall(ccompiler, statement.GetStatementCallNode(&s), d, &details)
}

func CompileNodeCall(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	functionName := node.GetNodeCallValue(&n)
	function, ok := compiler_context.GetCompilerFunction(&ccompiler.Context, functionName)
	if !ok {
		return d, awerrors.ErrorFailedToGetFunctionFromScope
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
		if d, err = CompileNodeValue(ccompiler, functionArguments[i], d, &argumentDetails); err != nil {
			return d, err
		}
		saveInstruction := encoder.AwooEncodedInstruction{
			Instruction: *instructions.AwooInstructionsSave[function.Arguments[i].Size],
			SourceOne:   cpu.AwooRegisterSavedZero,
			SourceTwo:   argumentDetails.Register,
			Immediate:   functionArgumentsOffset,
		}
		if d, err = encoder.Encode(saveInstruction, d); err != nil {
			return d, err
		}
		functionArgumentsOffset += uint32(function.Arguments[i].Size)
	}

	stackAdjustmentInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   stackOffset,
	}
	if d, err = encoder.Encode(stackAdjustmentInstruction, d); err != nil {
		return d, err
	}

	details.Register = cpu.AwooRegisterFunctionZero
	jumpInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   uint32(function.Start),
	}
	if d, err = encoder.Encode(jumpInstruction, d); err != nil {
		return d, err
	}

	stackAdjustmentInstruction = encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   -stackOffset,
	}
	return encoder.Encode(stackAdjustmentInstruction, d)
}