package compiler_symbol

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

type AwooCompilerSymbolTable struct {
	Entries  map[string]AwooCompilerSymbolTableEntry
	Position arch.AwooRegister
}

func NewCompilerSymbolTable() AwooCompilerSymbolTable {
	return AwooCompilerSymbolTable{
		Entries: map[string]AwooCompilerSymbolTableEntry{},
	}
}

type AwooCompilerSymbolTableEntry struct {
	Symbol elf.AwooElfSymbolTableEntry
	Global bool
}
