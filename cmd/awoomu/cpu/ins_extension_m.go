package cpu

func ProcessMUL(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] * cpu.Registers[ins.SourceTwo]
}

func ProcessDIV(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] / cpu.Registers[ins.SourceTwo]
}
