package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func PushSectionData(elf *elf.AwooElf, id elf.AwooElfSectionId, data []byte) {
	elf.SectionList.Sections[id].Contents = append(elf.SectionList.Sections[id].Contents, data...)
}

func ReadSectionData32(elf *elf.AwooElf, id elf.AwooElfSectionId, address arch.AwooRegister) uint32 {
	return uint32(elf.SectionList.Sections[id].Contents[address])<<24 |
		uint32(elf.SectionList.Sections[id].Contents[address+1])<<16 |
		uint32(elf.SectionList.Sections[id].Contents[address+2])<<8 |
		uint32(elf.SectionList.Sections[id].Contents[address+3])
}

func AlignSections(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf) {
	elf.SectionList.Sections[elf.SectionList.DataIndex].Address = arch.AwooRegister(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents))
}
