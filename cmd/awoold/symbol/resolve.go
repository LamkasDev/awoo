package symbol

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func ShiftImmediate(clinker *linker.AwooLinker, elf *commonElf.AwooElf, offset arch.AwooRegister, shift arch.AwooRegister) error {
	ins, err := decoder.Decode(clinker.Settings.Mappings.InstructionTable, arch.AwooInstruction(commonElf.ReadSectionData32(elf, commonElf.AwooElfSectionProgram, offset)))
	if err != nil {
		return err
	}
	// TODO: support shifts over 2047
	ins.Immediate = shift
	if err = encoder.EncodeAt(elf, offset, ins); err != nil {
		return err
	}

	return nil
}

func ResolveSymbols(clinker *linker.AwooLinker, elf *commonElf.AwooElf) error {
	for _, relocEntry := range elf.RelocationList {
		symbol, ok := elf.SymbolTable.Internal[relocEntry.Name]
		if !ok {
			return awerrors.ErrorUnknownVariable
		}
		if err := ShiftImmediate(clinker, elf, relocEntry.Offset, symbol.Start); err != nil {
			return err
		}
	}

	return nil
}
