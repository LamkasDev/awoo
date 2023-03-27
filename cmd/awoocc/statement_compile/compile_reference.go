package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeReference(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	// TODO: chaining references (only identifiers can be references anyways)
	variableNameNode := node.GetNodeSingleValue(&n)
	variableName := node.GetNodeIdentifierValue(&variableNameNode)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, variableName)
	if err != nil {
		return err
	}

	// TODO: merge this logic with primitives
	return encoder.Encode(elf, encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		Immediate:   uint32(variableMemory.Symbol.Start),
		Destination: details.Register,
	})
}

func CompileNodeDereference(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	err := CompileNodeValue(ccompiler, elf, node.GetNodeSingleValue(&n), details)
	if err != nil {
		return err
	}

	// TODO: merge this logic with identifiers
	// TODO: fix da 4
	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsLoad[4],
		Destination: details.Register,
		SourceOne:   details.Register,
	}
	return encoder.Encode(elf, loadInstruction)
}
