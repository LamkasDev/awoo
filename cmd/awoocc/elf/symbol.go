package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func AlignSymbols(ccompiler *compiler.AwooCompiler, elf *commonElf.AwooElf) error {
	for _, symbol := range elf.SymbolTable {
		if symbol.Type == commonTypes.AwooTypeId(types.AwooTypeFunction) {
			continue
		}
		symbol.Start += elf.SectionList.Sections[elf.SectionList.DataIndex].Address
		elf.SymbolTable[symbol.Name] = symbol
	}
	for _, relocEntry := range elf.RelocationList {
		ins, err := decoder.Decode(ccompiler.Settings.Mappings.InstructionTable, arch.AwooInstruction(commonElf.ReadSectionData32(elf, elf.SectionList.ProgramIndex, relocEntry.Offset)))
		if err != nil {
			return err
		}
		ins.Immediate = elf.SymbolTable[relocEntry.Name].Start
		encoder.EncodeAt(elf, relocEntry.Offset, ins)
	}

	return nil
}
