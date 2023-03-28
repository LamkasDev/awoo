package decoder

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

func Decode(table instructions.AwooInstructionTable, raw arch.AwooInstruction) (instruction.AwooInstruction, error) {
	code := (uint8)(raw) & instruction.AwooInstructionCodeMask
	subtable, ok := table[code]
	if !ok {
		return instruction.AwooInstruction{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownInstruction, gchalk.Red(fmt.Sprintf("%#x", code)))
	}
	format := instruction.AwooInstructionFormats[subtable.Format]
	argument := instruction.ProcessExtendedRange(raw, format.Argument, false)
	entry, ok := subtable.Subtable[(uint16)(argument)]
	if !ok {
		return instruction.AwooInstruction{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownInstruction, gchalk.Red(fmt.Sprintf("%#x", code)))
	}

	return instruction.AwooInstruction{
		Definition:  entry.Instruction,
		Process:     entry.Process,
		SourceOne:   cpu.AwooRegisterId(util.SelectRangeRegister(raw, format.SourceOne.Start, format.SourceOne.Length)),
		SourceTwo:   cpu.AwooRegisterId(util.SelectRangeRegister(raw, format.SourceTwo.Start, format.SourceTwo.Length)),
		Destination: cpu.AwooRegisterId(util.SelectRangeRegister(raw, format.Destination.Start, format.Destination.Length)),
		Immediate:   instruction.ProcessExtendedRange(raw, format.Immediate, true),
	}, nil
}
