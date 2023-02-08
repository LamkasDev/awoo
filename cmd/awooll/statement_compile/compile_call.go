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
	compiler_context.PushCompilerScope(&context.Scopes, "call")
	arguments := node.GetNodeCallArguments(&n)
	for i := 0; i < len(f.Arguments); i++ {
		// TODO: merge this logic
		funcArg := f.Arguments[i]
		dest, err := compiler_context.PushCompilerScopeCurrentMemory(context, compiler_context.AwooCompilerContextMemoryEntry{
			Name: funcArg.Name,
			Size: funcArg.Size,
			Type: funcArg.Type,
			Data: funcArg.Data,
		})
		if err != nil {
			return d, err
		}
		d, err = CompileNodeValueFast(context, arguments[i], d, details)
		if err != nil {
			return d, err
		}
		d, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionSW,
			SourceTwo:   details.Register,
			Immediate:   uint32(dest),
		}, d)
		if err != nil {
			return d, err
		}
	}
	compiler_context.PopCompilerScope(&context.Scopes)
	details.Register = cpu.AwooRegisterFunctionOne

	// TODO: this is retarder.
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		Destination: cpu.AwooRegisterStackPointer,
		Immediate:   uint32(f.Start),
	}, d)
}
