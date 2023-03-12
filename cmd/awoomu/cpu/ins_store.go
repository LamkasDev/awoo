package cpu

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/jwalton/gchalk"
)

func PrintSave(cpu *AwooCPU, ins AwooDecodedInstruction) {
	fmt.Printf("mem[%s] = %s",
		gchalk.BrightGreen(fmt.Sprintf("%#x", cpu.Registers[ins.SourceOne]+ins.Immediate)),
		gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
	)
}

/* func ProcessSD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessSW(cpu *AwooCPU, ins AwooDecodedInstruction) {
	memory.WriteMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWord(cpu.Registers[ins.SourceTwo]), memory.WriteMemoryWord)
	if arch.AwooDebug {
		PrintSave(cpu, ins)
	}
}

func ProcessSH(cpu *AwooCPU, ins AwooDecodedInstruction) {
	memory.WriteMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWordHalf(cpu.Registers[ins.SourceTwo]), memory.WriteMemoryWordHalf)
	if arch.AwooDebug {
		PrintSave(cpu, ins)
	}
}

func ProcessSB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	memory.WriteMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWordByte(cpu.Registers[ins.SourceTwo]), memory.WriteMemoryWordByte)
	if arch.AwooDebug {
		PrintSave(cpu, ins)
	}
}
