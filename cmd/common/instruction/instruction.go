package instruction

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooInstruction struct {
	Definition  AwooInstructionDefinition
	Process     interface{}
	SourceOne   arch.AwooRegister
	SourceTwo   arch.AwooRegister
	Destination arch.AwooRegister
	Immediate   arch.AwooRegister
}
