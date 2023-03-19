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

func CompileStatementAssignmentPointer(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	identifierNode = node.GetNodeSingleValue(&identifierNode)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeIdentifierValue(&identifierNode))
	if err != nil {
		return d, err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[*variableMemory.TypeDetails]

	valueNode := statement.GetStatementAssignmentValue(&s)
	details := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Type,
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if d, err = CompileNodeValue(ccompiler, valueNode, d, &details); err != nil {
		return d, err
	}

	addressRegister := cpu.GetNextTemporaryRegister(details.Register)
	loadInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type].Size],
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
		Instruction: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:   addressRegister,
		SourceTwo:   details.Register,
	}, d)
}
