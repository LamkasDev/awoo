package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

/* func ProcessSD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessSW(cpu *AwooCPU, ins AwooDecodedInstruction) {
	memory.WriteMemoryWord(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWord(cpu.Registers[ins.SourceTwo]))
}

func ProcessSH(cpu *AwooCPU, ins AwooDecodedInstruction) {
	memory.WriteMemoryWordHalf(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, arch.AwooWordHalf(cpu.Registers[ins.SourceTwo]))
}

func ProcessSB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	memory.WriteMemory8(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, byte(cpu.Registers[ins.SourceTwo]))
}
