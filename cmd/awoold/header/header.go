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
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

const AwooLinkerHeaderSize = 12

func ReserveHeader(clinker *linker.AwooLinker, celf *commonElf.AwooElf) {
	elf.ReserveSection(clinker, celf, commonElf.AwooElfSectionProgram, AwooLinkerHeaderSize)
}

func PopulateHeader(_ *linker.AwooLinker, celf *commonElf.AwooElf) error {
	stackOffset := celf.SectionList.Sections[commonElf.AwooElfSectionData].Address + arch.AwooRegister(len(celf.SectionList.Sections[commonElf.AwooElfSectionData].Contents))
	stackAdjustmentInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   stackOffset,
	}
	if err := encoder.EncodeAt(celf, 0, stackAdjustmentInstruction); err != nil {
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

	return nil
}
