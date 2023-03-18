package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeIdentifier(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	identifier := node.GetNodeIdentifierValue(&n)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, identifier)
	if err != nil {
		return d, err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type]

	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsLoad[variableType.Size],
		Destination: details.Register,
		Immediate:   uint32(variableMemory.Start),
	}
	if !variableMemory.Global {
		loadInstruction.SourceOne = cpu.AwooRegisterSavedZero
	}
	return encoder.Encode(loadInstruction, d)
}
