package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeIdentifier(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details compiler_context.CompileNodeValueDetails) ([]byte, error) {
	id := node.GetNodeIdentifierValue(&n)
	src, err := compiler_context.GetCompilerScopeMemory(context, id)
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToGetVariableFromScope, err)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[context.Parser.Lexer.Types.All[src.Type].Size],
		Destination: details.Register,
		Immediate:   uint32(src.Start),
	}, d)
}
