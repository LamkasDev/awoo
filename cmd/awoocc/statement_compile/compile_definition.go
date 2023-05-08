package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func GetCompilerMemoryEntry(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement) elf.AwooElfSymbolTableEntry {
	variableTypeNode := statement.GetStatementDefinitionVariableType(&s)
	variableNameNode := statement.GetStatementDefinitionVariableIdentifier(&s)
	variableName := node.GetNodeIdentifierValue(&variableNameNode)
	switch variableTypeNode.Type {
	case node.ParserNodeTypePointer:
		// TODO: chaining pointers
		pointedTypeNode := node.GetNodeSingleValue(&variableTypeNode)
		pointedType := node.GetNodeTypeType(&pointedTypeNode)
		return elf.AwooElfSymbolTableEntry{
			Name:    variableName,
			Type:    types.AwooTypePointer,
			Details: pointedType,
			Size:    ccompiler.Context.Parser.Lexer.Types.All[types.AwooTypePointer].Size,
		}
	case node.ParserNodeTypeTypeArray:
		arraySize := node.GetNodeTypeArraySize(&variableTypeNode)
		arrayTypeNode := node.GetNodeTypeArrayType(&variableTypeNode)
		arrayType := node.GetNodeTypeType(&arrayTypeNode)
		return elf.AwooElfSymbolTableEntry{
			Name: variableName,
			Type: arrayType,
			Size: arraySize * ccompiler.Context.Parser.Lexer.Types.All[arrayType].Size,
		}
	}

	variableType := node.GetNodeTypeType(&variableTypeNode)
	return elf.AwooElfSymbolTableEntry{
		Name: variableName,
		Type: variableType,
		Size: ccompiler.Context.Parser.Lexer.Types.All[variableType].Size,
	}
}

func CompileStatementDefinition(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	variableMemory, err := scope.PushCurrentFunctionSymbol(&ccompiler.Context.Scopes, GetCompilerMemoryEntry(ccompiler, s))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type]

	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	if valueNode == nil {
		return nil
	}
	// TODO: immediate should change to symbol name for resolving
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.AwooRegisterTemporaryZero,
		Address: compiler_details.CompileNodeValueDetailsAddress{
			Memory: variableMemory,
		},
	}
	if !variableMemory.Global {
		valueDetails.Address.Register = cpu.AwooRegisterSavedZero
	}
	if err = CompileNodeValue(ccompiler, celf, *valueNode, &valueDetails); err != nil {
		return err
	}
	if variableMemory.Global {
		elf.SetSymbol(celf, variableMemory.Symbol)
		elf.PushSectionData(celf, elf.AwooElfSectionData, make([]byte, variableMemory.Symbol.Size))
	}
	if valueDetails.Address.Used {
		return nil
	}

	saveInstruction := instruction.AwooInstruction{
		Definition: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:  valueDetails.Address.Register,
		SourceTwo:  valueDetails.Register,
		Immediate:  variableMemory.Symbol.Start,
	}
	if variableMemory.Global {
		elf.PushRelocationEntry(celf, variableMemory.Symbol.Name)
	}
	return encoder.Encode(celf, saveInstruction)
}
