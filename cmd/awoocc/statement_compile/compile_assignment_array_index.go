package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementAssignmentArrayIndex(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeArrayIndexIdentifier(&identifierNode))
	if err != nil {
		return d, err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type]

	valueNode := statement.GetStatementAssignmentValue(&s)
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Type,
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if d, err = CompileNodeValue(ccompiler, valueNode, d, &valueDetails); err != nil {
		return d, err
	}

	addressDetails := compiler_details.CompileNodeValueDetails{
		Register: cpu.GetNextTemporaryRegister(valueDetails.Register),
	}
	if d, err = CompileArrayIndexAddress(ccompiler, identifierNode, d, &addressDetails); err != nil {
		return d, err
	}
	if !variableMemory.Global {
		addressAdjustmentInstruction := encoder.AwooEncodedInstruction{
			Instruction: instructions.AwooInstructionADD,
			SourceOne:   addressDetails.Register,
			SourceTwo:   cpu.AwooRegisterSavedZero,
			Destination: addressDetails.Register,
		}
		if d, err = encoder.Encode(addressAdjustmentInstruction, d); err != nil {
			return d, err
		}
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:   addressDetails.Register,
		SourceTwo:   valueDetails.Register,
		Immediate:   uint32(variableMemory.Start),
	}
	return encoder.Encode(saveInstruction, d)
}
