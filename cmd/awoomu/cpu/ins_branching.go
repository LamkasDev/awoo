package cpu

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/jwalton/gchalk"
)

func ProcessBEQ(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] == cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s == %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBNE(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] != cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s != %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBGE(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] >= cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >= %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBGEU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if (arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) >= (arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]) {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >= %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]))),
		)
	}
}

func ProcessBLT(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] < cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBLTU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if (arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]) {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]))),
		)
	}
}

func ProcessJAL(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Counter + 4
	cpu.Counter += ins.Immediate
	cpu.Advance = false
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s = %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
		)
	}
}

// TODO: no multiplied immediate.
func ProcessJALR(cpu *AwooCPU, ins AwooDecodedInstruction) {
	t := cpu.Counter + 4
	cpu.Counter = (cpu.Registers[ins.SourceOne] + ins.Immediate) &^ 1
	cpu.Registers[ins.Destination] = t
	cpu.Advance = false
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s = %s)",
			gchalk.Yellow(AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprint(cpu.Registers[ins.Destination])),
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", cpu.Counter)),
		)
	}
}
