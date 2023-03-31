package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooElfRelocationList []AwooElfRelocationListEntry

type AwooElfRelocationListEntry struct {
	Offset arch.AwooRegister
	Name   string
}

// TODO: create a global function where globals get initialized.
// TODO: make linker prepend calls to global, main and adjust stack to end of data section.

func PushRelocationEntry(elf *AwooElf, name string) {
	elf.RelocationList = append(elf.RelocationList, AwooElfRelocationListEntry{
		Offset: arch.AwooRegister(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents)),
		Name:   name,
	})
}
