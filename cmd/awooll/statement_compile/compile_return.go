package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementReturn(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	currentScopeFunction := ccompiler.Context.Scopes.Functions[uint16(len(ccompiler.Context.Scopes.Functions)-1)]
	currentFunction := ccompiler.Context.Functions.Entries[currentScopeFunction.Name]
	if currentFunction.ReturnType != nil {
		returnValue := statement.GetStatementReturnValue(&s)
		d, err := CompileNodeValue(ccompiler, *returnValue, d, &compiler_details.CompileNodeValueDetails{
			Register: cpu.AwooRegisterFunctionZero,
		})
		if err != nil {
			return d, err
		}
	}

	// TODO: setup a proper return address stack.
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		SourceOne:   cpu.AwooRegisterStackPointer,
	}, d)
}
