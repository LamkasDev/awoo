package statement_compile

import (
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileArrayIndexAddress(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, addressDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeValue(ccompiler, node.GetNodeArrayIndexIndex(&n), d, addressDetails)
	if err != nil {
		return d, err
	}
	// TODO: add a method for sizes that are not power of 2
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLLI,
		SourceOne:   addressDetails.Register,
		Destination: addressDetails.Register,
		Immediate:   uint32(math.Log((float64)(ccompiler.Context.Parser.Lexer.Types.All[addressDetails.Type].Size)) / math.Log(2)),
	}, d)
}

func CompileNodeArrayIndex(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeArrayIndexIdentifier(&n))
	if err != nil {
		return d, err
	}

	addressDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Type,
		Register: cpu.GetNextTemporaryRegister(details.Register),
	}
	if d, err = CompileArrayIndexAddress(ccompiler, n, d, &addressDetails); err != nil {
		return d, err
	}
	if !variableMemory.Global {
		addressAdjustmentInstruction := encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionADD,
			SourceOne:   addressDetails.Register,
			SourceTwo:   cpu.AwooRegisterSavedZero,
			Destination: addressDetails.Register,
		}
		if d, err = encoder.Encode(addressAdjustmentInstruction, d); err != nil {
			return d, err
		}
	}

	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type].Size],
		SourceOne:   addressDetails.Register,
		Destination: details.Register,
		Immediate:   uint32(variableMemory.Start),
	}
	return encoder.Encode(loadInstruction, d)
}
