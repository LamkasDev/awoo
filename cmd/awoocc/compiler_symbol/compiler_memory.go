package compiler_symbol

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

type AwooCompilerSymbolTable struct {
	Entries  map[string]AwooCompilerSymbolTableEntry
	Position arch.AwooRegister
}

type AwooCompilerSymbolTableEntry struct {
	Symbol elf.AwooElfSymbolTableEntry
	Global bool
}
