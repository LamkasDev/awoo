package cpu

// TODO: fixed sign bits

func ProcessADD(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] + cpu.Registers[ins.SourceTwo]
}

func ProcessSUB(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceTwo] - cpu.Registers[ins.SourceOne]
}

func ProcessADDI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] + ins.Immediate
}

func ProcessSLT(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] < cpu.Registers[ins.SourceTwo] {
		cpu.Registers[ins.Destination] = 1
	} else {
		cpu.Registers[ins.Destination] = 0
	}
}

func ProcessSLTI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	if cpu.Registers[ins.SourceOne] < ins.Immediate {
		cpu.Registers[ins.Destination] = 1
	} else {
		cpu.Registers[ins.Destination] = 0
	}
}

func ProcessLUI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = ins.Immediate
}

func ProcessAUIPC(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Counter += ins.Immediate
	cpu.Registers[ins.Destination] = cpu.Counter
}
