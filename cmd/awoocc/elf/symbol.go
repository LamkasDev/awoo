package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func PopulateSymbols(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf) {
	for _, variable := range ccompiler.Context.Scopes.Global.Entries {
		elf.SymbolTable[variable.Symbol.Name] = variable.Symbol
	}
	for _, function := range ccompiler.Context.Functions.Entries {
		elf.SymbolTable[function.Symbol.Name] = function.Symbol
	}
}

func AlignSymbols(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf) error {
	for _, symbol := range elf.SymbolTable {
		symbol.Start += elf.SectionList.Sections[elf.SectionList.DataIndex].Address
		elf.SymbolTable[symbol.Name] = symbol
	}
	for _, relocEntry := range elf.RelocationList {
		ins, err := decoder.Decode(ccompiler.Settings.Mappings.InstructionTable, arch.AwooInstruction(ReadSectionData32(elf, elf.SectionList.ProgramIndex, relocEntry.Offset)))
		if err != nil {
			return err
		}
		ins.Immediate += elf.SymbolTable[relocEntry.Name].Start
		encoder.EncodeAt(elf, relocEntry.Offset, ins)
	}

	return nil
}
