package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooElfSymbolTable map[string]AwooElfSymbolTableEntry

type AwooElfSymbolTableEntry struct {
	Name        string
	Type        types.AwooTypeId
	TypeDetails *types.AwooTypeId
	Start       arch.AwooRegister
	Size        arch.AwooRegister
}

func PushSymbol(elf *AwooElf, symbol AwooElfSymbolTableEntry) {
	elf.SymbolTable[symbol.Name] = symbol
}
