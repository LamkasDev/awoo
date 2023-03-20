package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/types"

type AwooElfSymbolTable map[string]AwooElfSymbolTableEntry

type AwooElfSymbolTableEntry struct {
	Name        string
	Type        types.AwooTypeId
	TypeDetails *types.AwooTypeId
	Start       uint32
	Size        uint32
}
