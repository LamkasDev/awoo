package instructions

import "github.com/LamkasDev/awoo-emu/cmd/common/instruction"

// Arithmetic (9 instructions).
var AwooInstructionADD = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x0, Format: instruction.AwooInstructionFormatR, Name: "ADD", Advance: true,
}
var AwooInstructionSUB = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x100, Format: instruction.AwooInstructionFormatR, Name: "SUB", Advance: true,
}
var AwooInstructionADDI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x0, Format: instruction.AwooInstructionFormatI, Name: "ADDI", Advance: true,
}
var AwooInstructionSLT = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x2, Format: instruction.AwooInstructionFormatR, Name: "SLT", Advance: true,
}
var AwooInstructionSLTI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x2, Format: instruction.AwooInstructionFormatI, Name: "SLTI", Advance: true,
}
var AwooInstructionSLTU = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x3, Format: instruction.AwooInstructionFormatR, Name: "SLTU", Advance: true,
}
var AwooInstructionSLTIU = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x3, Format: instruction.AwooInstructionFormatI, Name: "SLTIU", Advance: true,
}
var AwooInstructionLUI = instruction.AwooInstructionDefinition{
	Code: 0b0110111, Argument: 0x0, Format: instruction.AwooInstructionFormatU, Name: "LUI", Advance: true,
}
var AwooInstructionAUIPC = instruction.AwooInstructionDefinition{
	Code: 0b0010111, Argument: 0x0, Format: instruction.AwooInstructionFormatU, Name: "AUIPC", Advance: true,
}

// Logical (12 instructions).
var AwooInstructionAND = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x7, Format: instruction.AwooInstructionFormatR, Name: "AND", Advance: true,
}
var AwooInstructionOR = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x6, Format: instruction.AwooInstructionFormatR, Name: "OR", Advance: true,
}
var AwooInstructionXOR = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x4, Format: instruction.AwooInstructionFormatR, Name: "XOR", Advance: true,
}
var AwooInstructionANDI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x7, Format: instruction.AwooInstructionFormatI, Name: "ANDI", Advance: true,
}
var AwooInstructionORI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x6, Format: instruction.AwooInstructionFormatI, Name: "ORI", Advance: true,
}
var AwooInstructionXORI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x4, Format: instruction.AwooInstructionFormatI, Name: "XORI", Advance: true,
}
var AwooInstructionSLL = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x1, Format: instruction.AwooInstructionFormatR, Name: "SLL", Advance: true,
}
var AwooInstructionSRL = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x5, Format: instruction.AwooInstructionFormatR, Name: "SRL", Advance: true,
}
var AwooInstructionSRA = instruction.AwooInstructionDefinition{
	Code: 0b0110011, Argument: 0x105, Format: instruction.AwooInstructionFormatI, Name: "SRA", Advance: true,
}
var AwooInstructionSLLI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x1, Format: instruction.AwooInstructionFormatI, Name: "SLLI", Advance: true,
}
var AwooInstructionSRLI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x5, Format: instruction.AwooInstructionFormatI, Name: "SRLI", Advance: true,
}
var AwooInstructionSRAI = instruction.AwooInstructionDefinition{
	Code: 0b0010011, Argument: 0x105, Format: instruction.AwooInstructionFormatI, Name: "SRAI", Advance: true,
}

// Load / Store (11 instructions).
// SetupInstructionTableEntry(table, 0b0000011, 0x3, instruction.AwooInstructionFormatI, "LD", nil).
var AwooInstructionLW = instruction.AwooInstructionDefinition{
	Code: 0b0000011, Argument: 0x2, Format: instruction.AwooInstructionFormatI, Name: "LW", Advance: true,
}
var AwooInstructionLH = instruction.AwooInstructionDefinition{
	Code: 0b0000011, Argument: 0x1, Format: instruction.AwooInstructionFormatI, Name: "LH", Advance: true,
}
var AwooInstructionLB = instruction.AwooInstructionDefinition{
	Code: 0b0000011, Argument: 0x0, Format: instruction.AwooInstructionFormatI, Name: "LB", Advance: true,
}

// SetupInstructionTableEntry(table, 0b0000011, 0x6, instruction.AwooInstructionFormatI, "LWU", nil).
var AwooInstructionLHU = instruction.AwooInstructionDefinition{
	Code: 0b0000011, Argument: 0x5, Format: instruction.AwooInstructionFormatI, Name: "LHU", Advance: true,
}
var AwooInstructionLBU = instruction.AwooInstructionDefinition{
	Code: 0b0000011, Argument: 0x4, Format: instruction.AwooInstructionFormatI, Name: "LBU", Advance: true,
}

// SetupInstructionTableEntry(table, 0b0100011, 0x3, instruction.AwooInstructionFormatS, "SD", nil).
var AwooInstructionSW = instruction.AwooInstructionDefinition{
	Code: 0b0100011, Argument: 0x2, Format: instruction.AwooInstructionFormatS, Name: "SW", Advance: true,
}
var AwooInstructionSH = instruction.AwooInstructionDefinition{
	Code: 0b0100011, Argument: 0x1, Format: instruction.AwooInstructionFormatS, Name: "SH", Advance: true,
}
var AwooInstructionSB = instruction.AwooInstructionDefinition{
	Code: 0b0100011, Argument: 0x0, Format: instruction.AwooInstructionFormatS, Name: "SB", Advance: true,
}

// Branching (8 instructions).
var AwooInstructionBEQ = instruction.AwooInstructionDefinition{
	Code: 0b1100011, Argument: 0x0, Format: instruction.AwooInstructionFormatB, Name: "BEQ", Advance: true,
}
var AwooInstructionBNE = instruction.AwooInstructionDefinition{
	Code: 0b1100011, Argument: 0x1, Format: instruction.AwooInstructionFormatB, Name: "BNE", Advance: true,
}
var AwooInstructionBGE = instruction.AwooInstructionDefinition{
	Code: 0b1100011, Argument: 0x5, Format: instruction.AwooInstructionFormatB, Name: "BGE", Advance: true,
}
var AwooInstructionBGEU = instruction.AwooInstructionDefinition{
	Code: 0b1100011, Argument: 0x7, Format: instruction.AwooInstructionFormatB, Name: "BGEU", Advance: true,
}
var AwooInstructionBLT = instruction.AwooInstructionDefinition{
	Code: 0b1100011, Argument: 0x4, Format: instruction.AwooInstructionFormatB, Name: "BLT", Advance: true,
}
var AwooInstructionBLTU = instruction.AwooInstructionDefinition{
	Code: 0b1100011, Argument: 0x6, Format: instruction.AwooInstructionFormatB, Name: "BLTU", Advance: true,
}
var AwooInstructionJAL = instruction.AwooInstructionDefinition{
	Code: 0b1101111, Argument: 0x0, Format: instruction.AwooInstructionFormatJ, Name: "JAL", Advance: true,
}
var AwooInstructionJALR = instruction.AwooInstructionDefinition{
	Code: 0b1100111, Argument: 0x0, Format: instruction.AwooInstructionFormatI, Name: "JALR", Advance: true,
}
