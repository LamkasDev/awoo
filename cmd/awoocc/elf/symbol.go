package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func PopulateSymbolTable(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf) {
	for _, variable := range ccompiler.Context.Scopes.Global.Entries {
		elf.SymbolTable[variable.Symbol.Name] = variable.Symbol
	}
	for _, function := range ccompiler.Context.Functions.Entries {
		elf.SymbolTable[function.Symbol.Name] = function.Symbol
	}
}
