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

func GetCompilerMemoryEntry(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement) compiler_context.AwooCompilerMemoryEntry {
	variableTypeNode := statement.GetStatementDefinitionVariableType(&s)
	variableNameNode := statement.GetStatementDefinitionVariableIdentifier(&s)
	variableName := node.GetNodeIdentifierValue(&variableNameNode)
	switch variableTypeNode.Type {
	case node.ParserNodeTypePointer:
		// TODO: chaining pointers
		pointedTypeNode := node.GetNodeSingleValue(&variableTypeNode)
		pointedType := node.GetNodeTypeType(&pointedTypeNode)
		return compiler_context.AwooCompilerMemoryEntry{
			Name: variableName,
			Type: types.AwooTypePointer,
			Data: pointedType,
			Size: ccompiler.Context.Parser.Lexer.Types.All[types.AwooTypePointer].Size,
		}
	case node.ParserNodeTypeArray:
		arraySize := node.GetNodeArraySize(&variableTypeNode)
		arrayTypeNode := node.GetNodeArrayType(&variableTypeNode)
		arrayType := node.GetNodeTypeType(&arrayTypeNode)
		return compiler_context.AwooCompilerMemoryEntry{
			Name: variableName,
			Type: arrayType,
			Size: arraySize * ccompiler.Context.Parser.Lexer.Types.All[arrayType].Size,
		}
	}

	variableType := node.GetNodeTypeType(&variableTypeNode)
	return compiler_context.AwooCompilerMemoryEntry{
		Name: variableName,
		Type: variableType,
		Size: ccompiler.Context.Parser.Lexer.Types.All[variableType].Size,
	}
}

func CompileStatementDefinition(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	variableMemory, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, GetCompilerMemoryEntry(ccompiler, s))
	if err != nil {
		return d, err
	}

	variableValueNode := statement.GetStatementDefinitionVariableValue(&s)
	if variableValueNode == nil {
		return d, nil
	}
	if d, err = CompileNodeValue(ccompiler, *variableValueNode, d, &details); err != nil {
		return d, err
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type].Size],
		SourceTwo:   details.Register,
		Immediate:   uint32(variableMemory.Start),
	}
	if !variableMemory.Global {
		saveInstruction.SourceOne = cpu.AwooRegisterSavedZero
	}

	return encoder.Encode(saveInstruction, d)
}
