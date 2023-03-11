package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
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
		argDetails := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
		d, err = CompileNodeValue(ccompiler, functionArguments[i], d, &argDetails)
		if err != nil {
			return d, err
		}
		d, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: *instruction.AwooInstructionsSave[function.Arguments[i].Size],
			SourceOne:   cpu.AwooRegisterSavedZero,
			SourceTwo:   details.Register,
			Immediate:   functionArgumentsOffset,
		}, d)
		if err != nil {
			return d, err
		}
		functionArgumentsOffset += uint32(function.Arguments[i].Size)
	}

	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   stackOffset,
	}, d)
	if err != nil {
		return d, err
	}

	details.Register = cpu.AwooRegisterFunctionZero
	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
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
