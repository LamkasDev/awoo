package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

/* func ProcessLD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessLW(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWord))
}

func ProcessLH(cpu *AwooCPU, ins AwooDecodedInstruction) {
	// TODO: sign extend.
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordHalf))
}

func ProcessLB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	// TODO: sign extend.
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordByte))
}

/* func ProcessLWU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessLHU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordHalf))
}

func ProcessLBU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordByte))
}
