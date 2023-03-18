package instructions

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func PrintSave(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	fmt.Printf("mem[%s] = %s",
		gchalk.BrightGreen(fmt.Sprintf("%#x", internal.CPU.Registers[ins.SourceOne]+ins.Immediate)),
		gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
	)
}

/* func ProcessSD(cpu *AwooCPU, ins instruction.AwooDecodedInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] & internal.CPU.Registers[ins.SourceTwo]
} */

func ProcessSW(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	memory.WriteMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWord(internal.CPU.Registers[ins.SourceTwo]), memory.WriteMemoryWord)
	if arch.AwooDebug {
		PrintSave(internal, ins)
	}
}

func ProcessSH(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	memory.WriteMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWordHalf(internal.CPU.Registers[ins.SourceTwo]), memory.WriteMemoryWordHalf)
	if arch.AwooDebug {
		PrintSave(internal, ins)
	}
}

func ProcessSB(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	memory.WriteMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWordByte(internal.CPU.Registers[ins.SourceTwo]), memory.WriteMemoryWordByte)
	if arch.AwooDebug {
		PrintSave(internal, ins)
	}
}
