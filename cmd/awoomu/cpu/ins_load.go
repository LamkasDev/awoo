package cpu

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/jwalton/gchalk"
)

func PrintLoad(cpu *AwooCPU, ins AwooDecodedInstruction) {
	fmt.Printf("%s = %s",
		gchalk.Yellow(AwooRegisterNames[ins.Destination]),
		gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
	)
}

/* func ProcessLD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessLW(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWord))
	if arch.AwooDebug {
		PrintLoad(cpu, ins)
	}
}

func ProcessLH(cpu *AwooCPU, ins AwooDecodedInstruction) {
	// TODO: sign extend.
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordHalf))
	if arch.AwooDebug {
		PrintLoad(cpu, ins)
	}
}

func ProcessLB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	// TODO: sign extend.
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordByte))
	if arch.AwooDebug {
		PrintLoad(cpu, ins)
	}
}

/* func ProcessLWU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
} */

func ProcessLHU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordHalf))
	if arch.AwooDebug {
		PrintLoad(cpu, ins)
	}
}

func ProcessLBU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&cpu.Memory, cpu.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordByte))
	if arch.AwooDebug {
		PrintLoad(cpu, ins)
	}
}
