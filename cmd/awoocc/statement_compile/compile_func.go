package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func CompileStatementFunc(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	functionNameNode := statement.GetStatementFuncIdentifier(&s)
	functionName := node.GetNodeIdentifierValue(&functionNameNode)

	compilerScopeFunction := compiler_context.AwooCompilerScopeFunction{
		Name:   functionName,
		Global: functionName == cc.AwooCompilerGlobalFunctionName,
	}
	compiler_context.PushCompilerScopeFunction(&ccompiler.Context, compilerScopeFunction)

	functionArguments := statement.GetStatementFuncArguments(&s)
	functionArgumentsOffset := arch.AwooRegister(0)
	for _, argument := range functionArguments {
		_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
			Symbol: commonElf.AwooElfSymbolTableEntry{
				Name:        argument.Name,
				Size:        argument.Size,
				Type:        argument.Type,
				TypeDetails: argument.TypeDetails,
			},
			Global: compilerScopeFunction.Global,
		})
		if err != nil {
			return err
		}
		functionArgumentsOffset += argument.Size
	}

	returnAddressMemory, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
		Symbol: commonElf.AwooElfSymbolTableEntry{
			Name: cc.AwooCompilerReturnAddressVariable,
			Size: 4,
			Type: commonTypes.AwooTypeId(types.AwooTypePointer),
		},
		Global: compilerScopeFunction.Global,
	})
	if compilerScopeFunction.Global {
		commonElf.PushSymbol(celf, returnAddressMemory.Symbol)
		elf.PushSectionData(celf, celf.SectionList.DataIndex, make([]byte, returnAddressMemory.Symbol.Size))
	}
	if err != nil {
		return err
	}

	functionReturnTypeNode := statement.GetStatementFuncReturnType(&s)
	var functionReturnType *commonTypes.AwooTypeId
	if functionReturnTypeNode != nil {
		returnType := node.GetNodeTypeType(functionReturnTypeNode)
		functionReturnType = &returnType
	}

	compilerFunction := compiler_context.AwooCompilerFunction{
		Symbol: commonElf.AwooElfSymbolTableEntry{
			Name:        functionName,
			Type:        commonTypes.AwooTypeId(types.AwooTypeFunction),
			TypeDetails: functionReturnType,
			Start:       arch.AwooRegister(len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents)),
		},
		Arguments: statement.GetStatementFuncArguments(&s),
	}
	compiler_context.PushCompilerFunction(&ccompiler.Context, compilerFunction)

	stackAdjustmentInstruction := instruction.AwooInstruction{
		Definition: instructions.AwooInstructionSW,
		SourceOne:  cpu.AwooRegisterSavedZero,
		SourceTwo:  cpu.AwooRegisterReturnAddress,
		Immediate:  functionArgumentsOffset,
	}
	if err = encoder.Encode(celf, stackAdjustmentInstruction); err != nil {
		return err
	}

	if err = CompileStatementGroup(ccompiler, celf, statement.GetStatementFuncBody(&s)); err != nil {
		return err
	}
	compilerFunction.Symbol.Size = arch.AwooRegister(len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents)) - compilerFunction.Symbol.Start
	compiler_context.SetSizeOfCompilerFunction(&ccompiler.Context, functionName, compilerFunction.Symbol.Size)
	commonElf.PushSymbol(celf, compilerFunction.Symbol)

	if !compilerScopeFunction.Global {
		compiler_context.PopCompilerScopeCurrentFunction(&ccompiler.Context)
	}
	return nil
}
