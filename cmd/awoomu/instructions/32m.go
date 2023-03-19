package instructions

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func ProcessMUL(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] * internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s * %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessDIV(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Registers[ins.SourceOne] / internal.CPU.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s / %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}
