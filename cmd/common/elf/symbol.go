package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooElfSymbolTable struct {
	Internal map[string]AwooElfSymbolTableEntry
	External map[string]AwooElfSymbolTableEntry
}

type AwooElfSymbolTableEntry struct {
	Name    string
	Type    types.AwooTypeId
	Details interface{}
	Start   arch.AwooRegister
	Size    arch.AwooRegister
}

func SetSymbol(elf *AwooElf, symbol AwooElfSymbolTableEntry) {
	elf.SymbolTable.Internal[symbol.Name] = symbol
}

func SetSymbolExternal(elf *AwooElf, symbol AwooElfSymbolTableEntry) {
	elf.SymbolTable.External[symbol.Name] = symbol
}

func GetSymbol(elf *AwooElf, name string) (AwooElfSymbolTableEntry, bool) {
	symbol, ok := elf.SymbolTable.Internal[name]
	if ok {
		return symbol, true
	}
	symbol, ok = elf.SymbolTable.External[name]
	if ok {
		return symbol, true
	}

	return AwooElfSymbolTableEntry{}, false
}

func MergeSymbolTableFunction(target map[string]AwooElfSymbolTableEntry, source map[string]AwooElfSymbolTableEntry, offset arch.AwooRegister) {
	for name, symbol := range source {
		if IsSymbolFunction(symbol) && symbol.Name != cc.AwooCompilerGlobalFunctionName {
			_, ok := target[name]
			if ok {
				panic("already exists")
			}
			symbol.Start += offset
			target[name] = symbol
		}
	}
}

func MergeSymbolTableVariable(target map[string]AwooElfSymbolTableEntry, source map[string]AwooElfSymbolTableEntry, offset arch.AwooRegister) {
	for name, symbol := range source {
		if !IsSymbolFunction(symbol) {
			_, ok := target[name]
			if ok {
				panic("already exists")
			}
			symbol.Start += offset
			target[name] = symbol
		}
	}
}
