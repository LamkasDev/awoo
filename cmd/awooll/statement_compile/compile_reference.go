package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeReference(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	// TODO: chaining references (only identifiers can be references anyways)
	idNode := node.GetNodeSingleValue(&n)
	id := node.GetNodeIdentifierValue(&idNode)
	src, err := compiler_context.GetCompilerScopeMemory(context, id)
	if err != nil {
		return d, err
	}

	// TODO: merge this logic with primitives
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADDI,
		Immediate:   uint32(src.Start),
		Destination: details.Register,
	}, d)
}

func CompileNodeDereference(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeValue(context, node.GetNodeSingleValue(&n), d, details)
	if err != nil {
		return d, err
	}

	// TODO: merge this logic with identifiers
	// TODO: fix da 4
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[4],
		Destination: details.Register,
		SourceOne:   details.Register,
	}, d)
}
