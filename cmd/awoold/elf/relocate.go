package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func PrependProgramData(_ *linker.AwooLinker, elf *commonElf.AwooElf, data []byte) error {
	offset := arch.AwooRegister(len(data))
	for name, symbol := range elf.SymbolTable {
		symbol.Start += offset
		elf.SymbolTable[name] = symbol
	}
	for i, relocEntry := range elf.RelocationList {
		relocEntry.Offset += offset
		elf.RelocationList[i] = relocEntry
	}
	elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents = append(data, elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents...)

	return nil
}
