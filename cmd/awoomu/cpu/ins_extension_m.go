package cpu

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/jwalton/gchalk"
)

func ProcessMUL(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] * cpu.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s * %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessDIV(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] / cpu.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s / %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}
