package instruction

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
)

type AwooInstruction struct {
	Definition  AwooInstructionDefinition
	Process     interface{}
	SourceOne   cpu.AwooRegisterId
	SourceTwo   cpu.AwooRegisterId
	Destination cpu.AwooRegisterId
	Immediate   arch.AwooRegister
}
