package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
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
	case node.ParserNodeTypeTypeArray:
		arraySize := node.GetNodeTypeArraySize(&variableTypeNode)
		arrayTypeNode := node.GetNodeTypeArrayType(&variableTypeNode)
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
	variableMemory, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, GetCompilerMemoryEntry(ccompiler, s))
	if err != nil {
		return d, err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type]

	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	if valueNode == nil {
		return d, nil
	}
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Type,
		Register: cpu.AwooRegisterTemporaryZero,
		Address: compiler_details.CompileNodeValueDetailsAddress{
			Immediate: variableMemory.Start,
		},
	}
	if !variableMemory.Global {
		valueDetails.Address.Register = cpu.AwooRegisterSavedZero
	}
	if d, err = CompileNodeValue(ccompiler, *valueNode, d, &valueDetails); err != nil {
		return d, err
	}
	if valueDetails.Address.Used {
		return d, nil
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:   valueDetails.Address.Register,
		SourceTwo:   valueDetails.Register,
		Immediate:   valueDetails.Address.Immediate,
	}
	return encoder.Encode(saveInstruction, d)
}
