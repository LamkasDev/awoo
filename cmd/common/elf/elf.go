package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/types"

type AwooElf struct {
	Sections          AwooElfSections
	SymbolTable       AwooElfSymbolTable
	RelocationDetails AwooElfRelocationTable
}

type AwooElfSections struct {
	Text AwooElfSection
	Data AwooElfSection
}

type AwooElfSection []byte

type AwooElfSymbolTable map[string]AwooElfSymbolTableEntry

type AwooElfSymbolTableEntry struct {
	Name        string
	Type        types.AwooTypeId
	TypeDetails *types.AwooTypeId
	Start       uint32
	Size        uint32
}

type AwooElfRelocationTable []AwooElfRelocationTableEntry

type AwooElfRelocationTableEntry struct {
	Offset uint32
	Name   string
}
