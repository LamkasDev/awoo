package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func ReserveSection(_ *linker.AwooLinker, celf *elf.AwooElf, section elf.AwooElfSectionId, size int) {
	celf.SectionList.Sections[section].Contents = append(celf.SectionList.Sections[section].Contents, make([]byte, size)...)
}
