package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooElf struct {
	Name           string
	Type           AwooElfTypeId
	Counter        arch.AwooRegister
	SectionList    AwooElfSectionList
	SymbolTable    AwooElfSymbolTable
	RelocationList AwooElfRelocationList
}

func NewAwooElf(name string, elfType AwooElfTypeId) AwooElf {
	return AwooElf{
		Name: name,
		Type: elfType,
		SectionList: AwooElfSectionList{
			Sections: []AwooElfSection{{Id: AwooElfSectionProgram, Contents: []byte{}}, {Id: AwooElfSectionData, Contents: []byte{}}},
		},
		SymbolTable: AwooElfSymbolTable{
			Internal: map[string]AwooElfSymbolTableEntry{},
			External: map[string]AwooElfSymbolTableEntry{},
		},
		RelocationList: AwooElfRelocationList{},
	}
}
