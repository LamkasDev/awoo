package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementReturn(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	currentScopeFunction := ccompiler.Context.Scopes.Functions[uint16(len(ccompiler.Context.Scopes.Functions)-1)]
	currentFunction, _ := compiler_context.GetCompilerFunction(&ccompiler.Context, currentScopeFunction.Name)
	if currentFunction.ReturnType != nil {
		returnValue := statement.GetStatementReturnValue(&s)
		d, err := CompileNodeValue(ccompiler, *returnValue, d, &compiler_details.CompileNodeValueDetails{
			Register: cpu.AwooRegisterFunctionZero,
		})
		if err != nil {
			return d, err
		}
	}

	d, err := encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionLW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Immediate:   uint32(compiler_context.GetCompilerFunctionArgumentsSize(currentFunction)),
		Destination: cpu.AwooRegisterReturnAddress,
	}, d)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		SourceOne:   cpu.AwooRegisterReturnAddress,
	}, d)
}
