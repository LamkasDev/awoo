package instructions

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func ProcessADD(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] + internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSUB(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] - internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s - %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessADDI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] + ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSLT(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if internal.CPU.Registers[ins.SourceOne] < internal.CPU.Registers[ins.SourceTwo] {
		internal.CPU.Registers[ins.Destination] = 1
	} else {
		internal.CPU.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSLTU(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceTwo]) {
		internal.CPU.Registers[ins.Destination] = 1
	} else {
		internal.CPU.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceTwo]))),
		)
	}
}

func ProcessSLTI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if internal.CPU.Registers[ins.SourceOne] < ins.Immediate {
		internal.CPU.Registers[ins.Destination] = 1
	} else {
		internal.CPU.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSLTIU(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(ins.Immediate) {
		internal.CPU.Registers[ins.Destination] = 1
	} else {
		internal.CPU.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(ins.Immediate))),
		)
	}
}

func ProcessLUI(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessAUIPC(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Counter += ins.Immediate
	internal.CPU.Registers[ins.Destination] = internal.CPU.Counter
}
