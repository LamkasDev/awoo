package instructions

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func ProcessBEQ(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if internal.CPU.Registers[ins.SourceOne] == internal.CPU.Registers[ins.SourceTwo] {
		internal.CPU.Counter += ins.Immediate
		internal.CPU.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s == %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBNE(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if internal.CPU.Registers[ins.SourceOne] != internal.CPU.Registers[ins.SourceTwo] {
		internal.CPU.Counter += ins.Immediate
		internal.CPU.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s != %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBGE(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if internal.CPU.Registers[ins.SourceOne] >= internal.CPU.Registers[ins.SourceTwo] {
		internal.CPU.Counter += ins.Immediate
		internal.CPU.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >= %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBGEU(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) >= (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceTwo]) {
		internal.CPU.Counter += ins.Immediate
		internal.CPU.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s >= %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceTwo]))),
		)
	}
}

func ProcessBLT(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if internal.CPU.Registers[ins.SourceOne] < internal.CPU.Registers[ins.SourceTwo] {
		internal.CPU.Counter += ins.Immediate
		internal.CPU.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceOne])),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.SourceTwo])),
		)
	}
}

func ProcessBLTU(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	if (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceTwo]) {
		internal.CPU.Counter += ins.Immediate
		internal.CPU.Advance = false
	}
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s < %s)",
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceOne]))),
			gchalk.Magenta(fmt.Sprint((arch.AwooRegisterU)(internal.CPU.Registers[ins.SourceTwo]))),
		)
	}
}

func ProcessJAL(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	internal.CPU.Registers[ins.Destination] = internal.CPU.Counter + 4
	internal.CPU.Counter += ins.Immediate
	internal.CPU.Advance = false
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s = %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
		)
	}
}

// TODO: no multiplied immediate.
func ProcessJALR(internal *internal.AwooEmulatorInternal, ins instruction.AwooInstruction) {
	t := internal.CPU.Counter + 4
	internal.CPU.Counter = (internal.CPU.Registers[ins.SourceOne] + ins.Immediate) &^ 1
	internal.CPU.Registers[ins.Destination] = t
	internal.CPU.Advance = false
	if arch.AwooDebug {
		fmt.Printf("%s = %s (%s = %s)",
			gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
			gchalk.Magenta(fmt.Sprint(internal.CPU.Registers[ins.Destination])),
			gchalk.Yellow("c"),
			gchalk.Magenta(fmt.Sprintf("%#x", internal.CPU.Counter)),
		)
	}
}
