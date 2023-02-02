package instruction

// Arithmetic (9 instructions)
var AwooInstructionADD = AwooInstruction{
	Code: 0b0110011, Argument: 0x0, Format: AwooInstructionFormatR, Name: "ADD", Advance: true,
}
var AwooInstructionSUB = AwooInstruction{
	Code: 0b0110011, Argument: 0x100, Format: AwooInstructionFormatR, Name: "SUB", Advance: true,
}
var AwooInstructionADDI = AwooInstruction{
	Code: 0b0010011, Argument: 0x0, Format: AwooInstructionFormatI, Name: "ADDI", Advance: true,
}
var AwooInstructionSLT = AwooInstruction{
	Code: 0b0110011, Argument: 0x2, Format: AwooInstructionFormatR, Name: "SLT", Advance: true,
}
var AwooInstructionSLTI = AwooInstruction{
	Code: 0b0010011, Argument: 0x2, Format: AwooInstructionFormatI, Name: "SLTI", Advance: true,
}
var AwooInstructionSLTU = AwooInstruction{
	Code: 0b0110011, Argument: 0x3, Format: AwooInstructionFormatR, Name: "SLTU", Advance: true,
}
var AwooInstructionSLTIU = AwooInstruction{
	Code: 0b0010011, Argument: 0x6, Format: AwooInstructionFormatI, Name: "SLTIU", Advance: true,
}
var AwooInstructionLUI = AwooInstruction{
	Code: 0b0110111, Argument: 0x0, Format: AwooInstructionFormatU, Name: "LUI", Advance: true,
}
var AwooInstructionAUIPC = AwooInstruction{
	Code: 0b0010111, Argument: 0x0, Format: AwooInstructionFormatU, Name: "AUIPC", Advance: true,
}

// Logical (12 instructions)
var AwooInstructionAND = AwooInstruction{
	Code: 0b0110011, Argument: 0x7, Format: AwooInstructionFormatR, Name: "AND", Advance: true,
}
var AwooInstructionOR = AwooInstruction{
	Code: 0b0110011, Argument: 0x6, Format: AwooInstructionFormatR, Name: "OR", Advance: true,
}
var AwooInstructionXOR = AwooInstruction{
	Code: 0b0110011, Argument: 0x4, Format: AwooInstructionFormatR, Name: "XOR", Advance: true,
}
var AwooInstructionANDI = AwooInstruction{
	Code: 0b0010011, Argument: 0x7, Format: AwooInstructionFormatI, Name: "ANDI", Advance: true,
}
var AwooInstructionORI = AwooInstruction{
	Code: 0b0110011, Argument: 0x6, Format: AwooInstructionFormatI, Name: "ORI", Advance: true,
}
var AwooInstructionXORI = AwooInstruction{
	Code: 0b0110011, Argument: 0x4, Format: AwooInstructionFormatI, Name: "XORI", Advance: true,
}
var AwooInstructionSLL = AwooInstruction{
	Code: 0b0110011, Argument: 0x1, Format: AwooInstructionFormatR, Name: "SLL", Advance: true,
}
var AwooInstructionSRL = AwooInstruction{
	Code: 0b0110011, Argument: 0x5, Format: AwooInstructionFormatR, Name: "SRL", Advance: true,
}
var AwooInstructionSRA = AwooInstruction{
	Code: 0b0110011, Argument: 0x105, Format: AwooInstructionFormatI, Name: "SRA", Advance: true,
}
var AwooInstructionSLLI = AwooInstruction{
	Code: 0b0010011, Argument: 0x1, Format: AwooInstructionFormatI, Name: "SLLI", Advance: true,
}
var AwooInstructionSRLI = AwooInstruction{
	Code: 0b0010011, Argument: 0x5, Format: AwooInstructionFormatI, Name: "SRLI", Advance: true,
}
var AwooInstructionSRAI = AwooInstruction{
	Code: 0b0010011, Argument: 0x105, Format: AwooInstructionFormatI, Name: "SRAI", Advance: true,
}

// Load / Store (11 instructions)
// SetupInstructionTableEntry(table, 0b0000011, 0x3, AwooInstructionFormatI, "LD", nil)
var AwooInstructionLW = AwooInstruction{
	Code: 0b0000011, Argument: 0x2, Format: AwooInstructionFormatI, Name: "LW", Advance: true,
}
var AwooInstructionLH = AwooInstruction{
	Code: 0b0000011, Argument: 0x1, Format: AwooInstructionFormatI, Name: "LH", Advance: true,
}
var AwooInstructionLB = AwooInstruction{
	Code: 0b0000011, Argument: 0x0, Format: AwooInstructionFormatI, Name: "LB", Advance: true,
}

// SetupInstructionTableEntry(table, 0b0000011, 0x6, AwooInstructionFormatI, "LWU", nil)
var AwooInstructionLHU = AwooInstruction{
	Code: 0b0000011, Argument: 0x5, Format: AwooInstructionFormatI, Name: "LHU", Advance: true,
}
var AwooInstructionLBU = AwooInstruction{
	Code: 0b0000011, Argument: 0x4, Format: AwooInstructionFormatI, Name: "LBU", Advance: true,
}

// SetupInstructionTableEntry(table, 0b0100011, 0x3, AwooInstructionFormatS, "SD", nil)
var AwooInstructionSW = AwooInstruction{
	Code: 0b0100011, Argument: 0x2, Format: AwooInstructionFormatS, Name: "SW", Advance: true,
}
var AwooInstructionSH = AwooInstruction{
	Code: 0b0100011, Argument: 0x1, Format: AwooInstructionFormatS, Name: "SH", Advance: true,
}
var AwooInstructionSB = AwooInstruction{
	Code: 0b0100011, Argument: 0x0, Format: AwooInstructionFormatS, Name: "SB", Advance: true,
}

// Branching (8 instructions)
var AwooInstructionBEQ = AwooInstruction{
	Code: 0b1100011, Argument: 0x0, Format: AwooInstructionFormatB, Name: "BEQ", Advance: true,
}
var AwooInstructionBNE = AwooInstruction{
	Code: 0b1100011, Argument: 0x1, Format: AwooInstructionFormatB, Name: "BNE", Advance: true,
}
var AwooInstructionBGE = AwooInstruction{
	Code: 0b1100011, Argument: 0x5, Format: AwooInstructionFormatB, Name: "BGE", Advance: true,
}
var AwooInstructionBGEU = AwooInstruction{
	Code: 0b1100011, Argument: 0x7, Format: AwooInstructionFormatB, Name: "BGEU", Advance: true,
}
var AwooInstructionBLT = AwooInstruction{
	Code: 0b1100011, Argument: 0x4, Format: AwooInstructionFormatB, Name: "BLT", Advance: true,
}
var AwooInstructionBLTU = AwooInstruction{
	Code: 0b1100011, Argument: 0x6, Format: AwooInstructionFormatB, Name: "BLTU", Advance: true,
}
var AwooInstructionJAL = AwooInstruction{
	Code: 0b1101111, Argument: 0x0, Format: AwooInstructionFormatJ, Name: "JAL", Advance: true,
}
var AwooInstructionJALR = AwooInstruction{
	Code: 0b1100111, Argument: 0x0, Format: AwooInstructionFormatI, Name: "JALR", Advance: true,
}
