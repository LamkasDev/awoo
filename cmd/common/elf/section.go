package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooElfSectionList struct {
	Sections []AwooElfSection
}

type AwooElfSectionId uint16

type AwooElfSection struct {
	Id       AwooElfSectionId
	Address  arch.AwooRegister
	Contents []byte
}

const AwooElfSectionProgram = AwooElfSectionId(0x00)
const AwooElfSectionData = AwooElfSectionId(0x01)

func AlignSections(elf *AwooElf) {
	elf.SectionList.Sections[AwooElfSectionData].Address = arch.AwooRegister(len(elf.SectionList.Sections[AwooElfSectionProgram].Contents))
}

func PushSectionData(elf *AwooElf, id AwooElfSectionId, data []byte) {
	elf.SectionList.Sections[id].Contents = append(elf.SectionList.Sections[id].Contents, data...)
}

func ReadSectionData32(elf *AwooElf, id AwooElfSectionId, address arch.AwooRegister) uint32 {
	return uint32(elf.SectionList.Sections[id].Contents[address])<<24 |
		uint32(elf.SectionList.Sections[id].Contents[address+1])<<16 |
		uint32(elf.SectionList.Sections[id].Contents[address+2])<<8 |
		uint32(elf.SectionList.Sections[id].Contents[address+3])
}
