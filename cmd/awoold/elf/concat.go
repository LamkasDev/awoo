package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction_helper"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func ConcatObjects(clinker *linker.AwooLinker, celf *elf.AwooElf) error {
	headerOffset := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))

	// Creates a new global function consisting of all global functions combined
	globalSize := arch.AwooRegister(0)
	globalOffset := headerOffset
	for _, libElf := range clinker.Contents {
		libGlobalFunction, ok := libElf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName]
		if ok {
			globalSize += libGlobalFunction.Size
		}
	}
	ReserveSection(clinker, celf, elf.AwooElfSectionProgram, int(globalSize))
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionLoadReturnAddress(0)); err != nil {
		return err
	}
	globalSize += 4
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionJumpToReturnAddress()); err != nil {
		return err
	}
	globalSize += 4
	celf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName] = elf.AwooElfSymbolTableEntry{
		Name: cc.AwooCompilerGlobalFunctionName,
		Type: types.AwooTypeFunction,
		Size: globalSize,
	}

	programOffset := headerOffset
	for _, libElf := range clinker.Contents {
		libProgramSize := arch.AwooRegister(len(libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))

		// Copies global function from dependency to executable
		libGlobalSize := arch.AwooRegister(0)
		libGlobalFunction, ok := libElf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName]
		if ok {
			libGlobalSize = libGlobalFunction.Size
			copy(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[globalOffset:], libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[:libGlobalSize])
			globalOffset += libGlobalSize
		}

		// Copies other functions from dependency to executable
		celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents = append(
			celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents,
			libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[libGlobalSize:]...,
		)

		// Adjusts position of relocation entries
		for _, relocEntry := range libElf.RelocationList {
			relocEntry.Offset += programOffset
			if relocEntry.Offset > libGlobalSize {
				// This is to ensure the created return is counted in, unless the entry exists before the return
				relocEntry.Offset += 8
			}
			celf.RelocationList = append(celf.RelocationList, relocEntry)
		}

		programOffset += libProgramSize
	}
	globalOffset += 8
	programOffset += 8

	// TODO: make sure return address and global offsets are counted in

	dataOffset := programOffset
	for _, libElf := range clinker.Contents {
		libDataSize := arch.AwooRegister(len(libElf.SectionList.Sections[elf.AwooElfSectionData].Contents))

		// Conjugates symbol table from dependency to executable
		elf.MergeSimpleSymbolTable(celf.SymbolTable.Internal, libElf.SymbolTable.Internal, dataOffset)

		// Copies data from dependency to executable
		celf.SectionList.Sections[elf.AwooElfSectionData].Contents = append(
			celf.SectionList.Sections[elf.AwooElfSectionData].Contents,
			libElf.SectionList.Sections[elf.AwooElfSectionData].Contents...,
		)

		dataOffset += libDataSize
	}

	return nil
}
