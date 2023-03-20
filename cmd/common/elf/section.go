package elf

type AwooElfSectionList struct {
	ProgramIndex AwooElfSectionId
	DataIndex    AwooElfSectionId
	Sections     []AwooElfSection
}

type AwooElfSectionId uint16

type AwooElfSection struct {
	Id       AwooElfSectionId
	Address  uint32
	Contents []byte
}
