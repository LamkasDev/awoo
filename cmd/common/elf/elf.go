package elf

type AwooElf struct {
	Type           AwooElfTypeId
	Counter        uint32
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
