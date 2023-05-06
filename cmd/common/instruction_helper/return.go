package instruction_helper

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func ConstructInstructionSaveReturnAddress() instruction.AwooInstruction {
	return instruction.AwooInstruction{
		Definition: instructions.AwooInstructionSW,
		SourceOne:  cpu.AwooRegisterSavedZero,
		SourceTwo:  cpu.AwooRegisterReturnAddress,
	}
}

func ConstructInstructionLoadReturnAddress() instruction.AwooInstruction {
	return instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionLW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterReturnAddress,
	}
}

func ConstructInstructionJumpToReturnAddress() instruction.AwooInstruction {
	return instruction.AwooInstruction{
		Definition: instructions.AwooInstructionJALR,
		SourceOne:  cpu.AwooRegisterReturnAddress,
	}
}
