package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func CompileStatementFunc(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	functionNameNode := statement.GetStatementFuncIdentifier(&s)
	functionName := node.GetNodeIdentifierValue(&functionNameNode)
	compiler_context.PushCompilerScopeFunction(&ccompiler.Context, compiler_context.AwooCompilerScopeFunction{
		Name: functionName,
	})

	functionArguments := statement.GetStatementFuncArguments(&s)
	functionArgumentsOffset := uint32(0)
	for _, argument := range functionArguments {
		_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
			Symbol: commonElf.AwooElfSymbolTableEntry{
				Name:        argument.Name,
				Size:        argument.Size,
				Type:        argument.Type,
				TypeDetails: argument.TypeDetails,
			},
		})
		if err != nil {
			return err
		}
		functionArgumentsOffset += uint32(argument.Size)
	}

	_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
		Symbol: commonElf.AwooElfSymbolTableEntry{
			Name: "_returnAddress",
			Size: 4,
			Type: commonTypes.AwooTypeId(types.AwooTypePointer),
		},
	})
	if err != nil {
		return err
	}

	functionReturnTypeNode := statement.GetStatementFuncReturnType(&s)
	var functionReturnType *commonTypes.AwooTypeId
	if functionReturnTypeNode != nil {
		returnType := node.GetNodeTypeType(functionReturnTypeNode)
		functionReturnType = &returnType
	}

	stackAdjustmentInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionSW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		SourceTwo:   cpu.AwooRegisterReturnAddress,
		Immediate:   functionArgumentsOffset,
	}
	if err = encoder.Encode(elf, stackAdjustmentInstruction); err != nil {
		return err
	}

	compilerFunction := compiler_context.AwooCompilerFunction{
		Symbol: commonElf.AwooElfSymbolTableEntry{
			Name:        functionName,
			Type:        commonTypes.AwooTypeId(types.AwooTypeFunction),
			TypeDetails: functionReturnType,
			Start:       uint32(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents)),
		},
		Arguments: statement.GetStatementFuncArguments(&s),
	}
	compiler_context.PushCompilerFunction(&ccompiler.Context, compilerFunction)
	if err = CompileStatementGroup(ccompiler, elf, statement.GetStatementFuncBody(&s)); err != nil {
		return err
	}
	compiler_context.SetSizeOfCompilerFunction(&ccompiler.Context, functionName, uint32(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents))-compilerFunction.Symbol.Start)
	compiler_context.PopCompilerScopeCurrentFunction(&ccompiler.Context)

	return nil
}
