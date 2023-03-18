package instructions

import "github.com/LamkasDev/awoo-emu/cmd/common/instruction"

// Multiply / Divide extension
var AwooInstructionMUL = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x8, Format: instruction.AwooInstructionFormatR, Name: "MUL", Advance: true,
}

var AwooInstructionDIV = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x9, Format: instruction.AwooInstructionFormatR, Name: "DIV", Advance: true,
}
