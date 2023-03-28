package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

// TODO: create a global function where globals get initialized.
// TODO: make linker prepend calls to global, main and adjust stack to end of data section.

func PushRelocationEntry(elf *commonElf.AwooElf, name string) {
	elf.RelocationList = append(elf.RelocationList, commonElf.AwooElfRelocationListEntry{
		Offset: arch.AwooRegister(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents)),
		Name:   name,
	})
}
