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
	arguments := node.GetNodeCallArguments(&n)
	var err error
	for _, arg := range arguments {
		d, err = CompileNodeValueFast(context, arg, d, &compiler_context.CompileNodeValueDetails{
			Register: cpu.AwooRegisterFunctionOne,
		})
		if err != nil {
			return d, err
		}
	}
	details.Register = cpu.AwooRegisterFunctionOne

	// TODO: this is retarder.
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		Destination: cpu.AwooRegisterStackPointer,
		Immediate:   uint32(f.Start),
	}, d)
}
