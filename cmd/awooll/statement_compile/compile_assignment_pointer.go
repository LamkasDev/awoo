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

func CompileStatementAssignmentPointer(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	identifierNode = node.GetNodeSingleValue(&identifierNode)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeIdentifierValue(&identifierNode))
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Data.(uint16)]
	if err != nil {
		return d, err
	}
	valueNode := statement.GetStatementAssignmentValue(&s)
	if d, err = CompileNodeValue(ccompiler, valueNode, d, &details); err != nil {
		return d, err
	}

	addressRegister := cpu.GetNextTemporaryRegister(details.Register)
	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type].Size],
		Destination: addressRegister,
		Immediate:   uint32(variableMemory.Start),
	}
	if !variableMemory.Global {
		loadInstruction.SourceOne = cpu.AwooRegisterSavedZero
	}
	if d, err = encoder.Encode(loadInstruction, d); err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[variableType.Size],
		SourceOne:   addressRegister,
		SourceTwo:   details.Register,
	}, d)
}
