package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func PrependProgramData(clinker *linker.AwooLinker, elf *commonElf.AwooElf, data []byte) error {
	offset := arch.AwooRegister(len(data))
	for name, symbol := range elf.SymbolTable {
		symbol.Start += offset
		elf.SymbolTable[name] = symbol
	}
	for i, relocEntry := range elf.RelocationList {
		ins, err := decoder.Decode(clinker.Settings.Mappings.InstructionTable, arch.AwooInstruction(commonElf.ReadSectionData32(elf, elf.SectionList.ProgramIndex, relocEntry.Offset)))
		if err != nil {
			return err
		}
		ins.Immediate += offset
		encoder.EncodeAt(elf, relocEntry.Offset, ins)
		elf.RelocationList[i].Offset += offset
	}
	elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents = append(data, elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents...)

	return nil
}
