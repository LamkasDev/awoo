package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementAssignmentArrayIndex(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeArrayIndexIdentifier(&identifierNode))
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type]
	if err != nil {
		return d, err
	}
	valueNode := statement.GetStatementAssignmentValue(&s)
	if d, err = CompileNodeValue(ccompiler, valueNode, d, &details); err != nil {
		return d, err
	}

	addressRegister := cpu.GetNextTemporaryRegister(details.Register)
	if d, err = CompileArrayIndexAddress(ccompiler, identifierNode, d, addressRegister, variableType); err != nil {
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

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[variableType.Size],
		SourceOne:   addressRegister,
		SourceTwo:   details.Register,
		Immediate:   uint32(variableMemory.Start),
	}

	return encoder.Encode(saveInstruction, d)
}
