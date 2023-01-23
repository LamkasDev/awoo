package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/arch"
)

func ProcessBEQ(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] == cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
}

func ProcessBNE(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] != cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
}

func ProcessBGE(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] >= cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
}

func ProcessBGEU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if (arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) >= (arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]) {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
}

func ProcessBLT(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] < cpu.Registers[ins.SourceTwo] {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
}

func ProcessBLTU(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if (arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) < (arch.AwooRegisterU)(cpu.Registers[ins.SourceTwo]) {
		cpu.Counter += ins.Immediate
		cpu.Advance = false
	}
}

func ProcessJAL(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Counter + 4
	cpu.Counter += ins.Immediate
	cpu.Advance = false
}

// TODO: no multiplied immediate
func ProcessJALR(cpu *AwooCPU, ins AwooDecodedInstruction) {
	t := cpu.Counter + 4
	cpu.Counter = (cpu.Registers[ins.SourceOne] + ins.Immediate) &^ 1
	cpu.Registers[ins.Destination] = t
	cpu.Advance = false
}
