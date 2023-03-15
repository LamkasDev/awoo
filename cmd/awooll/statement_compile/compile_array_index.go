package statement_compile

import (
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileArrayIndexAddress(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, register uint8, variableType types.AwooType) ([]byte, error) {
	d, err := CompileNodeValue(ccompiler, node.GetNodeArrayIndexIndex(&n), d, &compiler_details.CompileNodeValueDetails{
		Register: register,
	})
	if err != nil {
		return d, err
	}
	// TODO: add a method for sizes that are not power of 2
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLLI,
		SourceOne:   register,
		Destination: register,
		Immediate:   uint32(math.Log((float64)(variableType.Size)) / math.Log(2)),
	}, d)
}

func CompileNodeArrayIndex(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	identifier := node.GetNodeArrayIndexIdentifier(&n)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, identifier)
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type]
	if err != nil {
		return d, err
	}

	addressRegister := cpu.GetNextTemporaryRegister(details.Register)
	if d, err = CompileArrayIndexAddress(ccompiler, n, d, addressRegister, variableType); err != nil {
		return d, err
	}
	if !variableMemory.Global {
		d, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionADD,
			SourceOne:   addressRegister,
			SourceTwo:   cpu.AwooRegisterSavedZero,
			Destination: addressRegister,
		}, d)
		if err != nil {
			return d, err
		}
	}

	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type].Size],
		SourceOne:   addressRegister,
		Destination: details.Register,
		Immediate:   uint32(variableMemory.Start),
	}

	return encoder.Encode(loadInstruction, d)
}
