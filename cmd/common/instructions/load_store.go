package instructions

import "github.com/LamkasDev/awoo-emu/cmd/common/instruction"

var AwooInstructionsLoad = map[uint32]*instruction.AwooInstructionDefinition{
	1: &AwooInstructionLB,
	2: &AwooInstructionLH,
	4: &AwooInstructionLW,
}

var AwooInstructionsSave = map[uint32]*instruction.AwooInstructionDefinition{
	1: &AwooInstructionSB,
	2: &AwooInstructionSH,
	4: &AwooInstructionSW,
}
