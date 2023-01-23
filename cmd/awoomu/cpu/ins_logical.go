package cpu

import "github.com/LamkasDev/awoo-emu/cmd/awoomu/arch"

func ProcessAND(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & cpu.Registers[ins.SourceTwo]
}

func ProcessOR(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] | cpu.Registers[ins.SourceTwo]
}

func ProcessXOR(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] ^ cpu.Registers[ins.SourceTwo]
}

func ProcessANDI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] & ins.Immediate
}

func ProcessORI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] | ins.Immediate
}

func ProcessXORI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] ^ ins.Immediate
}

func ProcessSLL(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) << cpu.Registers[ins.SourceTwo])
}

func ProcessSRL(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) >> cpu.Registers[ins.SourceTwo])
}

func ProcessSRA(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] >> cpu.Registers[ins.SourceTwo]
}

func ProcessSLLI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) << ins.Immediate)
}

func ProcessSRLI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = (arch.AwooRegister)((arch.AwooRegisterU)(cpu.Registers[ins.SourceOne]) >> ins.Immediate)
}

func ProcessSRAI(cpu *AwooCPU, ins AwooDecodedInstruction) {
	cpu.Registers[ins.Destination] = cpu.Registers[ins.SourceOne] >> ins.Immediate
}
