package symbol

import (
	"errors"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func ShiftImmediate(clinker *linker.AwooLinker, elf *commonElf.AwooElf, offset arch.AwooRegister, shift arch.AwooRegister) error {
	ins, err := decoder.Decode(clinker.Settings.Mappings.InstructionTable, arch.AwooInstruction(commonElf.ReadSectionData32(elf, elf.SectionList.ProgramIndex, offset)))
	if err != nil {
		return err
	}
	ins.Immediate = shift
	encoder.EncodeAt(elf, offset, ins)

	return nil
}

func ResolveSymbols(clinker *linker.AwooLinker, elf *commonElf.AwooElf) error {
	for _, relocEntry := range elf.RelocationList {
		symbol, ok := clinker.Contents.SymbolTable[relocEntry.Name]
		if !ok {
			return errors.New("fuck")
		}
		if err := ShiftImmediate(clinker, elf, relocEntry.Offset, symbol.Start); err != nil {
			return err
		}
	}

	return nil
}
