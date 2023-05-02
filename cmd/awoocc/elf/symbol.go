package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func AlignSymbols(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf) error {
	for _, symbol := range celf.SymbolTable.Internal {
		if symbol.Type == types.AwooTypeFunction {
			continue
		}
		symbol.Start += celf.SectionList.Sections[elf.AwooElfSectionData].Address
		celf.SymbolTable.Internal[symbol.Name] = symbol
	}
	for _, relocEntry := range celf.RelocationList {
		ins, err := decoder.Decode(ccompiler.Settings.Mappings.InstructionTable, arch.AwooInstruction(elf.ReadSectionData32(celf, elf.AwooElfSectionProgram, relocEntry.Offset)))
		if err != nil {
			return err
		}
		ins.Immediate = celf.SymbolTable.Internal[relocEntry.Name].Start
		encoder.EncodeAt(celf, relocEntry.Offset, ins)
	}

	return nil
}
