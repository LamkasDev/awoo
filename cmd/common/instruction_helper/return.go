package instruction_helper

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func ConstructInstructionLoadReturnAddress(immediate arch.AwooRegister) instruction.AwooInstruction {
	return instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionLW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Immediate:   immediate,
		Destination: cpu.AwooRegisterReturnAddress,
	}
}

func ConstructInstructionJumpToReturnAddress() instruction.AwooInstruction {
	return instruction.AwooInstruction{
		Definition: instructions.AwooInstructionJALR,
		SourceOne:  cpu.AwooRegisterReturnAddress,
	}
}
