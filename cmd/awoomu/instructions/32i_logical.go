package instructions

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func ProcessAND(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] & internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s & %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessOR(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] | internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s | %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessXOR(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] ^ internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s ^ %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessANDI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] & ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s & %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessORI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] | ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s | %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessXORI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] ^ ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s ^ %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSLL(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) << internal.CPU.Registers[ins.SourceTwo])
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s << %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSRL(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) >> internal.CPU.Registers[ins.SourceTwo])
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >> %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSRA(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] >> internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >> %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSLLI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) << ins.Immediate)
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s << %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSRLI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) >> ins.Immediate)
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >> %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSRAI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] >> ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >> %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}
