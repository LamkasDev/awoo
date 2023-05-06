package header

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction_helper"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

const AwooLinkerHeaderSize = 20

func ReserveHeader(clinker *linker.AwooLinker, celf *commonElf.AwooElf) {
	elf.ReserveSection(clinker, celf, commonElf.AwooElfSectionProgram, AwooLinkerHeaderSize)
}

func PopulateHeader(_ *linker.AwooLinker, celf *commonElf.AwooElf) error {
	// Makes sure stack starts after data section
	stackOffset := celf.SectionList.Sections[commonElf.AwooElfSectionData].Address + arch.AwooRegister(len(celf.SectionList.Sections[commonElf.AwooElfSectionData].Contents))
	if err := encoder.EncodeAt(celf, 0, instruction_helper.ConstructInstructionAdjustStack(stackOffset)); err != nil {
		return err
	}

	jumpGlobalInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   celf.SymbolTable.Internal[cc.AwooCompilerGlobalFunctionName].Start,
	}
	if err := encoder.EncodeAt(celf, 4, jumpGlobalInstruction); err != nil {
		return err
	}

	jumpMainInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   celf.SymbolTable.Internal[cc.AwooCompilerMainFunctionName].Start,
	}
	if err := encoder.EncodeAt(celf, 8, jumpMainInstruction); err != nil {
		return err
	}

	createMstatusInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionADDI,
		Destination: cpu.AwooRegisterTemporaryZero,
		Immediate:   cpu.AwooCrsMstatusHalt,
	}
	if err := encoder.EncodeAt(celf, 12, createMstatusInstruction); err != nil {
		return err
	}

	saveMstatusInstruction := instruction.AwooInstruction{
		Definition: instructions.AwooInstructionSW,
		SourceTwo:  cpu.AwooRegisterTemporaryZero,
		Immediate:  cpu.AwooCrsMstatus,
	}
	if err := encoder.EncodeAt(celf, 16, saveMstatusInstruction); err != nil {
		return err
	}

	return nil
}
