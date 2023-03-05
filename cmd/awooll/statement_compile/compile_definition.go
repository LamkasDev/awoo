package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementDefinition(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}

	variableTypeNode := statement.GetStatementDefinitionVariableType(&s)
	variableNameNode := statement.GetStatementDefinitionVariableIdentifier(&s)
	entry := compiler_context.AwooCompilerContextMemoryEntry{}
	switch variableTypeNode.Type {
	case node.ParserNodeTypeType:
		entry.Type = node.GetNodeTypeType(&variableTypeNode)
	case node.ParserNodeTypePointer:
		entry.Type = types.AwooTypePointer
		// TODO: chaining pointers
		variableTypeNode = node.GetNodeSingleValue(&variableTypeNode)
		entry.Data = node.GetNodeTypeType(&variableTypeNode)
	}
	entry.Size = ccompiler.Context.Parser.Lexer.Types.All[entry.Type].Size
	entry.Name = node.GetNodeIdentifierValue(&variableNameNode)

	variableValueNode := statement.GetStatementDefinitionVariableValue(&s)
	variableMemory, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, entry)
	if err != nil {
		return d, err
	}
	d, err = CompileNodeValue(ccompiler, variableValueNode, d, &details)
	if err != nil {
		return d, err
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[ccompiler.Context.Parser.Lexer.Types.All[entry.Type].Size],
		SourceTwo:   details.Register,
		Immediate:   uint32(variableMemory.Start),
	}
	if !variableMemory.Global {
		saveInstruction.SourceOne = cpu.AwooRegisterSavedZero
	}
	return encoder.Encode(saveInstruction, d)
}
