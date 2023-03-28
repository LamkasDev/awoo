package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	awooElf "github.com/LamkasDev/awoo-emu/cmd/awoocc/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
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
			Symbol: elf.AwooElfSymbolTableEntry{
				Name:        variableName,
				Type:        commonTypes.AwooTypeId(types.AwooTypePointer),
				TypeDetails: &pointedType,
				Size:        ccompiler.Context.Parser.Lexer.Types.All[commonTypes.AwooTypeId(types.AwooTypePointer)].Size,
			},
		}
	case node.ParserNodeTypeTypeArray:
		arraySize := node.GetNodeTypeArraySize(&variableTypeNode)
		arrayTypeNode := node.GetNodeTypeArrayType(&variableTypeNode)
		arrayType := node.GetNodeTypeType(&arrayTypeNode)
		return compiler_context.AwooCompilerMemoryEntry{
			Symbol: elf.AwooElfSymbolTableEntry{
				Name: variableName,
				Type: arrayType,
				Size: arraySize * ccompiler.Context.Parser.Lexer.Types.All[arrayType].Size,
			},
		}
	}

	variableType := node.GetNodeTypeType(&variableTypeNode)
	return compiler_context.AwooCompilerMemoryEntry{
		Symbol: elf.AwooElfSymbolTableEntry{
			Name: variableName,
			Type: variableType,
			Size: ccompiler.Context.Parser.Lexer.Types.All[variableType].Size,
		},
	}
}

func CompileStatementDefinition(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	variableMemory, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, GetCompilerMemoryEntry(ccompiler, s))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type]

	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	if valueNode == nil {
		return nil
	}
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.AwooRegisterTemporaryZero,
		Address: compiler_details.CompileNodeValueDetailsAddress{
			Immediate: variableMemory.Symbol.Start,
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

	saveInstruction := instruction.AwooInstruction{
		Definition: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:  valueDetails.Address.Register,
		SourceTwo:  valueDetails.Register,
		Immediate:  valueDetails.Address.Immediate,
	}
	if variableMemory.Global {
		awooElf.PushRelocationEntry(elf, variableMemory.Symbol.Name)
	}
	return encoder.Encode(elf, saveInstruction)
}
