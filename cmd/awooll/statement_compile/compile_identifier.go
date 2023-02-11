package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeIdentifier(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	identifier := node.GetNodeIdentifierValue(&n)
	identifierMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, identifier)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[identifierMemory.Type].Size],
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: details.Register,
		Immediate:   uint32(identifierMemory.Start),
	}, d)
}
