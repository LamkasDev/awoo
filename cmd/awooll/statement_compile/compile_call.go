package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeCall(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	functionName := node.GetNodeCallValue(&n)
	function, ok := compiler_context.GetCompilerFunction(&ccompiler.Context, functionName)
	if !ok {
		return d, awerrors.ErrorFailedToGetFunctionFromScope
	}

	stackOffset := uint32(compiler_context.GetCompilerScopeCurrentFunctionSize(&ccompiler.Context))
	d, err := encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   stackOffset,
	}, d)
	if err != nil {
		return d, err
	}

	arguments := node.GetNodeCallArguments(&n)
	argumentsOffset := uint32(0)
	for i := 0; i < len(function.Arguments); i++ {
		argDetails := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
		d, err = CompileNodeValue(ccompiler, arguments[i], d, &argDetails)
		if err != nil {
			return d, err
		}
		d, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionSW,
			SourceOne:   cpu.AwooRegisterSavedZero,
			SourceTwo:   details.Register,
			Immediate:   argumentsOffset,
		}, d)
		argumentsOffset += uint32(function.Arguments[i].Size)
		if err != nil {
			return d, err
		}
	}

	details.Register = cpu.AwooRegisterFunctionZero
	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		Destination: cpu.AwooRegisterStackPointer,
		Immediate:   uint32(function.Start),
	}, d)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   -stackOffset,
	}, d)
}
