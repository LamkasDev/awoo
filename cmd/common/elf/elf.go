package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooElf struct {
	Type           AwooElfTypeId
	Counter        arch.AwooRegister
	SectionList    AwooElfSectionList
	SymbolTable    AwooElfSymbolTable
	RelocationList AwooElfRelocationList
}

func NewAwooElf(elfType AwooElfTypeId) AwooElf {
	return AwooElf{
		Type: elfType,
		SectionList: AwooElfSectionList{
			ProgramIndex: 0,
			DataIndex:    1,
			Sections:     []AwooElfSection{{Id: 0, Contents: []byte{}}, {Id: 1, Contents: []byte{}}},
		},
		SymbolTable:    AwooElfSymbolTable{},
		RelocationList: AwooElfRelocationList{},
	}
}
