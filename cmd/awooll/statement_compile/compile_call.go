package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeCall(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	id := node.GetNodeCallValue(&n)
	f, ok := compiler_context.GetCompilerFunction(context, id)
	if !ok {
		return d, awerrors.ErrorFailedToGetFunctionFromScope
	}
	// compiler_context.PushCompilerScopeBlock(&context.Scopes, "call")
	// arguments := node.GetNodeCallArguments(&n)
	for i := 0; i < len(f.Arguments); i++ {
		// TODO: merge this logic
		// funcArg := f.Arguments[i]
		// TODO: push argument onto stack
	}
	// compiler_context.PopCompilerScopeFunction(&context.Scopes)
	details.Register = cpu.AwooRegisterFunctionZero

	// TODO: determine if stack adjustment is required (not if called function doesn't allocate variables)
	adj := uint32(compiler_context.GetCompilerScopeCurrentFunctionSize(context))
	d, err := encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   adj,
	}, d)
	if err != nil {
		return d, err
	}
	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		Destination: cpu.AwooRegisterStackPointer,
		Immediate:   uint32(f.Start),
	}, d)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   -adj,
	}, d)
}
