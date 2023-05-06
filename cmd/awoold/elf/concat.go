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

const AwooGlobalPrefixSize = 8
const AwooGlobalSuffixSize = 12
const AwooGlobalEdgesSize = AwooGlobalPrefixSize + AwooGlobalSuffixSize

func GetGlobalFunctionSize(celf *elf.AwooElf) arch.AwooRegister {
	if globalFunction, ok := celf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName]; ok {
		return globalFunction.Size
	}

	return 0
}

func CalculateGlobalFunctionSize(clinker *linker.AwooLinker) arch.AwooRegister {
	size := arch.AwooRegister(0)
	for _, libElf := range clinker.Contents {
		libGlobalFunction, ok := libElf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName]
		if ok {
			size += libGlobalFunction.Size
		}
	}

	return size
}

func ConcatObjects(clinker *linker.AwooLinker, celf *elf.AwooElf) error {
	headerOffset := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))

	// Creates a new global function consisting of all global functions combined
	globalOffset := headerOffset
	globalSize := CalculateGlobalFunctionSize(clinker)
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionSaveReturnAddress()); err != nil {
		return err
	}
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionAdjustStack(4)); err != nil {
		return err
	}
	ReserveSection(clinker, celf, elf.AwooElfSectionProgram, int(globalSize))
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionAdjustStack(-4)); err != nil {
		return err
	}
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionLoadReturnAddress()); err != nil {
		return err
	}
	if err := encoder.Encode(celf, instruction_helper.ConstructInstructionJumpToReturnAddress()); err != nil {
		return err
	}
	globalSize += AwooGlobalEdgesSize
	celf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName] = elf.AwooElfSymbolTableEntry{
		Name:  cc.AwooCompilerGlobalFunctionName,
		Type:  types.AwooTypeFunction,
		Start: globalOffset,
		Size:  globalSize,
	}
	globalOffset += AwooGlobalPrefixSize

	for name, libElf := range clinker.Contents {
		libGlobalFunction, ok := libElf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName]
		if ok {
			// Copies global function from dependency to executable
			copy(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[globalOffset:], libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[:libGlobalFunction.Size])

			// Removes global section from contents so it's easier to process further
			libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents = libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[libGlobalFunction.Size:]
			clinker.Contents[name] = libElf

			// Adjusts position of relocation entries inside global function
			for _, relocEntry := range libElf.RelocationList {
				if relocEntry.Offset <= libGlobalFunction.Size {
					relocEntry.Offset += globalOffset
					celf.RelocationList = append(celf.RelocationList, relocEntry)
				}
			}

			globalOffset += libGlobalFunction.Size
		}
	}
	globalOffset += AwooGlobalSuffixSize

	programOffset := globalOffset
	programSize := arch.AwooRegister(0)
	for _, libElf := range clinker.Contents {
		libProgramSize := arch.AwooRegister(len(libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
		libGlobalSize := GetGlobalFunctionSize(&libElf)

		// Copies other functions from dependency to executable
		celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents = append(
			celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents,
			libElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents...,
		)

		// Conjugates function symbol table from dependency to executable
		elf.MergeSymbolTableFunction(celf.SymbolTable.Internal, libElf.SymbolTable.Internal, programOffset-libGlobalSize)

		// Adjusts position of relocation entries outside global function
		for _, relocEntry := range libElf.RelocationList {
			if relocEntry.Offset > libGlobalSize {
				relocEntry.Offset += programOffset - libGlobalSize
				celf.RelocationList = append(celf.RelocationList, relocEntry)
			}
		}

		programOffset += libProgramSize
		programSize += libProgramSize
	}

	dataOffset := programOffset
	dataSize := arch.AwooRegister(0)
	for _, libElf := range clinker.Contents {
		libDataSize := arch.AwooRegister(len(libElf.SectionList.Sections[elf.AwooElfSectionData].Contents))

		// Conjugates variable symbol table from dependency to executable
		elf.MergeSymbolTableVariable(celf.SymbolTable.Internal, libElf.SymbolTable.Internal, dataOffset)

		// Copies data from dependency to executable
		celf.SectionList.Sections[elf.AwooElfSectionData].Contents = append(
			celf.SectionList.Sections[elf.AwooElfSectionData].Contents,
			libElf.SectionList.Sections[elf.AwooElfSectionData].Contents...,
		)

		dataOffset += libDataSize
		dataSize += libDataSize
	}

	return nil
}
