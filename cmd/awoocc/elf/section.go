package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/elf"

func PushSectionData(elf *elf.AwooElf, id elf.AwooElfSectionId, data []byte) {
	elf.SectionList.Sections[id].Contents = append(elf.SectionList.Sections[id].Contents, data...)
}
