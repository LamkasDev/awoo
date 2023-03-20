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
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
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
			Name:        variableName,
			Type:        commonTypes.AwooTypeId(types.AwooTypePointer),
			TypeDetails: &pointedType,
			Size:        ccompiler.Context.Parser.Lexer.Types.All[commonTypes.AwooTypeId(types.AwooTypePointer)].Size,
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

func CompileStatementDefinition(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	variableMemory, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, GetCompilerMemoryEntry(ccompiler, s))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Type]

	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	if valueNode == nil {
		return nil
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
	if err = CompileNodeValue(ccompiler, elf, *valueNode, &valueDetails); err != nil {
		return err
	}
	if valueDetails.Address.Used {
		return nil
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:   valueDetails.Address.Register,
		SourceTwo:   valueDetails.Register,
		Immediate:   valueDetails.Address.Immediate,
	}
	return encoder.Encode(elf, saveInstruction)
}
