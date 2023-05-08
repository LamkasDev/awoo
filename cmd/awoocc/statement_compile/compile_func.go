package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction_helper"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func LoadExternalSymbols(ccompiler *compiler.AwooCompiler) error {
	for _, symbol := range ccompiler.Context.Parser.Scopes.Functions[scope.AwooScopeGlobalFunctionId].Blocks[scope.AwooScopeGlobalBlockId].SymbolTable.External {
		if _, err := scope.PushFunctionBlockSymbolExternal(&ccompiler.Context.Scopes, symbol.Symbol); err != nil {
			return err
		}
	}

	return nil
}

func CompileStatementFunc(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	functionBody := statement.GetStatementFuncBody(&s)
	functionNameNode := statement.GetStatementFuncIdentifier(&s)
	functionName := node.GetNodeIdentifierValue(&functionNameNode)

	compilerScopeFunction := scope.NewScopeFunction(functionName)
	scope.PushFunction(&ccompiler.Context.Scopes, compilerScopeFunction)
	if scope.IsFunctionGlobal(compilerScopeFunction) {
		if err := LoadExternalSymbols(ccompiler); err != nil {
			return err
		}
	}

	if !scope.IsFunctionGlobal(compilerScopeFunction) {
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
			if _, err := scope.PushCurrentFunctionSymbol(&ccompiler.Context.Scopes, argument); err != nil {
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

	if !scope.IsFunctionGlobal(compilerScopeFunction) {
		if err := encoder.Encode(celf, instruction_helper.ConstructInstructionSaveReturnAddress()); err != nil {
			return err
		}
	}

	if err := CompileStatementGroup(ccompiler, celf, functionBody); err != nil {
		return err
	}
	compilerFunction.Size = arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents)) - compilerFunction.Start
	elf.SetSymbol(celf, compilerFunction)

	if !scope.IsFunctionGlobal(compilerScopeFunction) {
		scope.PopCurrentFunction(&ccompiler.Context.Scopes)
	}
	return nil
}
