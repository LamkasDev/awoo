package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementReturn(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	currentScopeFunction := ccompiler.Context.Scopes.Functions[uint16(len(ccompiler.Context.Scopes.Functions)-1)]
	currentFunction, _ := compiler_context.GetCompilerFunction(&ccompiler.Context, currentScopeFunction.Name)
	var err error
	if currentFunction.ReturnType != nil {
		returnValueNode := statement.GetStatementReturnValue(&s)
		returnDetails := compiler_details.CompileNodeValueDetails{
			Type:     *currentFunction.ReturnType,
			Register: cpu.AwooRegisterFunctionZero,
		}
		if d, err = CompileNodeValue(ccompiler, *returnValueNode, d, &returnDetails); err != nil {
			return d, err
		}
	}

	loadReturnAddressInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionLW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Immediate:   uint32(compiler_context.GetCompilerFunctionArgumentsSize(currentFunction)),
		Destination: cpu.AwooRegisterReturnAddress,
	}
	if d, err = encoder.Encode(loadReturnAddressInstruction, d); err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJALR,
		SourceOne:   cpu.AwooRegisterReturnAddress,
	}, d)
}
