package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/arch"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
)

/* func ProcessLD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessLW(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemoryWord(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate))
}

func ProcessLH(cpu *AwooCPU, ins AwooDecodedInstruction) {
	// TODO: sign extend
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemoryWordHalf(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate))
}

func ProcessLB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	// TODO: sign extend
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemory8(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate))
}

/* func ProcessLWU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessLHU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemoryWordHalf(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate))
}

func ProcessLBU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemory8(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate))
}

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
