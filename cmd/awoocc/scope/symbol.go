package scope

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

type AwooScopeSymbolTable struct {
	Internal map[string]AwooScopeSymbolTableEntry
	External map[string]AwooScopeSymbolTableEntry
	Position arch.AwooRegister
}

func NewScopeSymbolTable(position arch.AwooRegister) AwooScopeSymbolTable {
	return AwooScopeSymbolTable{
		Internal: map[string]AwooScopeSymbolTableEntry{},
		External: map[string]AwooScopeSymbolTableEntry{},
		Position: position,
	}
}

type AwooScopeSymbolTableEntry struct {
	Symbol elf.AwooElfSymbolTableEntry
	Global bool
}
