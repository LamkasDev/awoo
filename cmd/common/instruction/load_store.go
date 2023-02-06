package instruction

var AwooInstructionsLoad = map[uint16]*AwooInstruction{
	1: &AwooInstructionLB,
	2: &AwooInstructionLH,
	4: &AwooInstructionLW,
}

var AwooInstructionsSave = map[uint16]*AwooInstruction{
	1: &AwooInstructionSB,
	2: &AwooInstructionSH,
	4: &AwooInstructionSW,
}
