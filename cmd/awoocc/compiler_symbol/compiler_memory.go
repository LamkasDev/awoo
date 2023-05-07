package compiler_symbol

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

type AwooCompilerSymbolTable struct {
	Internal map[string]AwooCompilerSymbolTableEntry
	External map[string]AwooCompilerSymbolTableEntry
	Position arch.AwooRegister
}

func NewCompilerSymbolTable() AwooCompilerSymbolTable {
	return AwooCompilerSymbolTable{
		Internal: map[string]AwooCompilerSymbolTableEntry{},
		External: map[string]AwooCompilerSymbolTableEntry{},
	}
}

type AwooCompilerSymbolTableEntry struct {
	Symbol elf.AwooElfSymbolTableEntry
	Global bool
}
