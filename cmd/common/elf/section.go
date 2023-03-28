package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooElfSectionList struct {
	ProgramIndex AwooElfSectionId
	DataIndex    AwooElfSectionId
	Sections     []AwooElfSection
}

type AwooElfSectionId uint16

type AwooElfSection struct {
	Id       AwooElfSectionId
	Address  arch.AwooRegister
	Contents []byte
}
