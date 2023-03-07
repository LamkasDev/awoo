package internal

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintInternalInstruction(internal *AwooEmulatorInternal, raw []byte, ins cpu.AwooDecodedInstruction) {
	baseDetails := cpu.PrintDecodedInstruction(ins)
	if logger.AwooLoggerExtraEnabled {
		extraDetails := PrintInternalInstructionExtra(internal, ins)
		if extraDetails != "" {
			baseDetails = fmt.Sprintf("%s %s", baseDetails, extraDetails)
		}
	}

	logger.Log(
		"c: %s; r: %s; %s\n",
		gchalk.Red(fmt.Sprintf("%#6x", internal.CPU.Counter)),
		gchalk.Cyan(fmt.Sprintf("%#x %#x %#x %#x", raw[0:1], raw[1:2], raw[2:3], raw[3:4])),
		baseDetails,
	)
}

// TODO: refactor to a map.
func PrintInternalInstructionExtra(internal *AwooEmulatorInternal, ins cpu.AwooDecodedInstruction) string {
	switch ins.Instruction.Code {
	case instruction.AwooInstructionADD.Code:
		return fmt.Sprintf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne]+internal.CPU.Registers[ins.SourceTwo])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	case instruction.AwooInstructionSUB.Code:
		return fmt.Sprintf("%s = %s (%s - %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne]-internal.CPU.Registers[ins.SourceTwo])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	case instruction.AwooInstructionLUI.Code:
		return fmt.Sprintf("%s = %s",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	case instruction.AwooInstructionADDI.Code:
		return fmt.Sprintf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne]+ins.Immediate)),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	case instruction.AwooInstructionLW.Code:
		return fmt.Sprintf("%s = %s",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(memory.ReadMemoryWord(&internal.CPU.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate))),
		)
	case instruction.AwooInstructionLH.Code:
		return fmt.Sprintf("%s = %s",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(memory.ReadMemoryWordHalf(&internal.CPU.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate))),
		)
	case instruction.AwooInstructionLB.Code:
		return fmt.Sprintf("%s = %s",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(memory.ReadMemoryWordByte(&internal.CPU.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate))),
		)
	case instruction.AwooInstructionSW.Code,
		instruction.AwooInstructionSH.Code,
		instruction.AwooInstructionSB.Code:
		return fmt.Sprintf("mem[%s] = %s",
			gchalk.BrightGreen(fmt.Sprintf("%#x", internal.CPU.Registers[ins.SourceOne]+ins.Immediate)),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	case instruction.AwooInstructionJAL.Code:
		return fmt.Sprintf("%s = %s (c = %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.BrightGreen(fmt.Sprintf("%#x", internal.CPU.Counter+4)),
			gchalk.BrightGreen(fmt.Sprintf("%#x", internal.CPU.Counter+ins.Immediate)),
		)
	case instruction.AwooInstructionJALR.Code:
		return fmt.Sprintf("%s = %s (c = %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.BrightGreen(fmt.Sprintf("%#x", internal.CPU.Counter+4)),
			gchalk.BrightGreen(fmt.Sprintf("%#x", (internal.CPU.Registers[ins.SourceOne]+ins.Immediate)&^1)),
		)
	}

	return ""
}
