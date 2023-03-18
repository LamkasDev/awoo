package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeReference(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	// TODO: chaining references (only identifiers can be references anyways)
	variableNameNode := node.GetNodeSingleValue(&n)
	variableName := node.GetNodeIdentifierValue(&variableNameNode)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, variableName)
	if err != nil {
		return d, err
	}

	// TODO: merge this logic with primitives
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		Immediate:   uint32(variableMemory.Start),
		Destination: details.Register,
	}, d)
}

func CompileNodeDereference(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeValue(ccompiler, node.GetNodeSingleValue(&n), d, details)
	if err != nil {
		return d, err
	}

	// TODO: merge this logic with identifiers
	// TODO: fix da 4
	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsLoad[4],
		Destination: details.Register,
		SourceOne:   details.Register,
	}
	return encoder.Encode(loadInstruction, d)
}
