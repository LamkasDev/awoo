package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeIdentifier(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	id := node.GetNodeIdentifierValue(&n)
	src, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(context, id)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[context.Parser.Lexer.Types.All[src.Type].Size],
		SourceOne:   cpu.AwooRegisterSavedOne,
		Destination: details.Register,
		Immediate:   uint32(src.Start),
	}, d)
}
