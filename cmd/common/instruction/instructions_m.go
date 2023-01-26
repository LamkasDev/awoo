package instruction

// Multiply / Divide extension
var AwooInstructionMUL = AwooInstruction{
	Code: 0b0110011, Argument: 0x8, Format: AwooInstructionFormatR, Name: "MUL", Advance: true,
}

var AwooInstructionDIV = AwooInstruction{
	Code: 0b0110011, Argument: 0x9, Format: AwooInstructionFormatR, Name: "DIV", Advance: true,
}
