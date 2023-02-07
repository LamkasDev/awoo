package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeNegative(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeValue(context, node.GetNodeSingleValue(&n), d, details)
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCompileNode, err)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSUB,
		Destination: details.Register,
		SourceTwo:   details.Register,
	}, d)
}
