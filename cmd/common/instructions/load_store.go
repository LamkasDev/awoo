package instructions

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

var AwooInstructionsLoad = map[arch.AwooRegister]*instruction.AwooInstructionDefinition{
	1: &AwooInstructionLB,
	2: &AwooInstructionLH,
	4: &AwooInstructionLW,
}

var AwooInstructionsSave = map[arch.AwooRegister]*instruction.AwooInstructionDefinition{
	1: &AwooInstructionSB,
	2: &AwooInstructionSH,
	4: &AwooInstructionSW,
}
