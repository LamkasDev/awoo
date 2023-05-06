package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_symbol"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction_helper"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func CompileStatementFunc(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	functionBody := statement.GetStatementFuncBody(&s)
	if len(statement.GetStatementGroupBody(&functionBody)) == 0 {
		return nil
	}

	functionNameNode := statement.GetStatementFuncIdentifier(&s)
	functionName := node.GetNodeIdentifierValue(&functionNameNode)

	compilerScopeFunction := compiler_context.NewCompilerScopeFunction(functionName)
	compiler_context.PushCompilerScopeFunction(&ccompiler.Context, compilerScopeFunction)

	if !compiler_context.IsCompilerScopeFunctionGlobal(compilerScopeFunction) {
		functionArguments := []elf.AwooElfSymbolTableEntry{
			{
				Name: cc.AwooCompilerReturnAddressVariable,
				Size: cc.AwooCompilerReturnAddressSize,
				Type: types.AwooTypePointer,
			},
		}
		functionArguments = append(functionArguments, statement.GetStatementFuncArguments(&s)...)
		functionArgumentsOffset := arch.AwooRegister(0)
		for _, argument := range functionArguments {
			_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_symbol.AwooCompilerSymbolTableEntry{
				Symbol: argument,
				Global: false,
			})
			if err != nil {
				return err
			}
			functionArgumentsOffset += argument.Size
		}
	}

	compilerFunction := elf.AwooElfSymbolTableEntry{
		Name: functionName,
		Type: types.AwooTypeFunction,
		Details: elf.AwooElfSymbolTableEntryFunctionDetails{
			ReturnType: statement.GetStatementFuncReturnTypePrecise(&s),
			Arguments:  statement.GetStatementFuncArguments(&s),
		},
		Start: arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents)),
	}
	elf.SetSymbol(celf, compilerFunction)

	if !compiler_context.IsCompilerScopeFunctionGlobal(compilerScopeFunction) {
		if err := encoder.Encode(celf, instruction_helper.ConstructInstructionSaveReturnAddress()); err != nil {
			return err
		}
	}

	if err := CompileStatementGroup(ccompiler, celf, functionBody); err != nil {
		return err
	}
	compilerFunction.Size = arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents)) - compilerFunction.Start
	elf.SetSymbol(celf, compilerFunction)

	if !compiler_context.IsCompilerScopeFunctionGlobal(compilerScopeFunction) {
		compiler_context.PopCompilerScopeCurrentFunction(&ccompiler.Context)
	}
	return nil
}
