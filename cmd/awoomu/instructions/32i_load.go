package instructions

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func PrintLoad(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	fmt.Printf("%s = %s",
		gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
		gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
	)
}

/* func ProcessLD(cpu *AwooCPU, ins instruction.AwooDecodedInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] & internal.CPU.Registers[ins.SourceTwo]
} */

func ProcessLW(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWord))
	if arch.AwooDebug {
		PrintLoad(internal, ins)
	}
}

func ProcessLH(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	// TODO: sign extend.
	internal.CPU.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordHalf))
	if arch.AwooDebug {
		PrintLoad(internal, ins)
	}
}

func ProcessLB(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	// TODO: sign extend.
	internal.CPU.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordByte))
	if arch.AwooDebug {
		PrintLoad(internal, ins)
	}
}

/* func ProcessLWU(cpu *AwooCPU, ins instruction.AwooDecodedInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] & internal.CPU.Registers[ins.SourceTwo]
} */

func ProcessLHU(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordHalf))
	if arch.AwooDebug {
		PrintLoad(internal, ins)
	}
}

func ProcessLBU(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = arch.AwooRegister(memory.ReadMemorySafe(&internal.Memory, internal.CPU.Registers[ins.SourceOne]+ins.Immediate, memory.ReadMemoryWordByte))
	if arch.AwooDebug {
		PrintLoad(internal, ins)
	}
}
