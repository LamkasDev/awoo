package cpu

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/jwalton/gchalk"
)

func ProcessADD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] + cpu.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s + %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSUB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] - cpu.Registers[ins.SourceTwo]
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s - %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessADDI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] + ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s + %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSLT(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] < cpu.Registers[ins.SourceTwo] {
		cpu.Registers[ins.Destination] = 1
	} else {
		cpu.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessSLTU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if (arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]) {
		cpu.Registers[ins.Destination] = 1
	} else {
		cpu.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]))),
		)
	}
}

func ProcessSLTI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] < ins.Immediate {
		cpu.Registers[ins.Destination] = 1
	} else {
		cpu.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSLTIU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if (arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(ins.Immediate) {
		cpu.Registers[ins.Destination] = 1
	} else {
		cpu.Registers[ins.Destination] = 0
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(ins.Immediate))),
		)
	}
}

func ProcessLUI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = ins.Immediate
	if arch.AwooDebug {
		fmt.Printf("%s = %s",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessAUIPC(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Counter += ins.Immediate
	cpu.Registers[ins.Destination] = cpu.Counter
}
