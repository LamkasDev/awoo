package cpu

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/arch"
	"github.com/LamkasDev/awoo-emu/cmd/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/util"
	"github.com/jwalton/gchalk"
)

const AwooOpCodeMask = 0b01111111

type AwooDecodedInstructionProcess func(cpu *AwooCPU, ins AwooDecodedInstruction)
type AwooDecodedInstruction struct {
	Instruction instruction.AwooInstruction
	Process     AwooDecodedInstructionProcess
	SourceOne   arch.AwooRegister
	SourceTwo   arch.AwooRegister
	Destination arch.AwooRegister
	Immediate   arch.AwooRegister
}

func Decode(table AwooInstructionTable, raw uint32) (AwooDecodedInstruction, error) {
	code := (uint8)(raw) & AwooOpCodeMask
	subtable, ok := table[code]
	if !ok {
		return AwooDecodedInstruction{}, fmt.Errorf("unknown instruction %s", gchalk.Red(fmt.Sprintf("0x%x", code)))
	}
	format := instruction.AwooInstructionFormats[subtable.Format]
	argument := instruction.ProcessExtendedRange(raw, format.Argument)
	entry, ok := subtable.Subtable[(uint16)(argument)]
	if !ok {
		return AwooDecodedInstruction{}, fmt.Errorf("unknown instruction %s", gchalk.Red(fmt.Sprintf("0x%x", code)))
	}

	return AwooDecodedInstruction{
		Instruction: entry.Instruction,
		Process:     entry.Process,
		SourceOne:   util.SelectRangeRegister(raw, format.SourceOne.Start, format.SourceOne.Length),
		SourceTwo:   util.SelectRangeRegister(raw, format.SourceTwo.Start, format.SourceTwo.Length),
		Destination: util.SelectRangeRegister(raw, format.Destination.Start, format.Destination.Length),
		Immediate:   instruction.ProcessExtendedRange(raw, format.Immediate),
	}, nil
}
