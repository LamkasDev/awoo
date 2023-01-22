package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/instruction"
)

type AwooInstructionTableEntry struct {
	Instruction instruction.AwooInstruction
	Process     AwooDecodedInstructionProcess
}
type AwooInstructionTableSubtable struct {
	Format   uint8
	Subtable map[uint16]AwooInstructionTableEntry
}
type AwooInstructionTable map[uint8]AwooInstructionTableSubtable

func SetupInstructionTableEntry(table AwooInstructionTable, code uint8, arg uint16, format uint8, name string, Process AwooDecodedInstructionProcess) {
	_, ok := table[code]
	if !ok {
		table[code] = AwooInstructionTableSubtable{
			Format:   format,
			Subtable: make(map[uint16]AwooInstructionTableEntry),
		}
	}

	table[code].Subtable[arg] = AwooInstructionTableEntry{
		Instruction: instruction.AwooInstruction{Code: code, Format: format, Name: name},
		Process:     Process,
	}
}

func SetupInstructionTable() AwooInstructionTable {
	table := AwooInstructionTable{}

	// Arithmetic (9 instructions)
	SetupInstructionTableEntry(table, 0b0110011, 0x0, instruction.AwooInstructionFormatR, "ADD", ProcessADD)
	SetupInstructionTableEntry(table, 0b0110011, 0x100, instruction.AwooInstructionFormatR, "SUB", ProcessSUB)
	SetupInstructionTableEntry(table, 0b0010011, 0x0, instruction.AwooInstructionFormatI, "ADDI", ProcessADDI)
	SetupInstructionTableEntry(table, 0b0110011, 0x2, instruction.AwooInstructionFormatR, "SLT", ProcessSLT)
	SetupInstructionTableEntry(table, 0b0010011, 0x2, instruction.AwooInstructionFormatI, "SLTI", ProcessSLTI)
	// SetupInstructionTableEntry(table, 0b0110011, 0x6, instruction.AwooInstructionFormatR, "SLTU", ProcessSLTU)
	// SetupInstructionTableEntry(table, 0b0010011, 0x6, instruction.AwooInstructionFormatI, "SLTIU", ProcessSLTIU)
	SetupInstructionTableEntry(table, 0b0110111, 0x0, instruction.AwooInstructionFormatU, "LUI", ProcessLUI)
	SetupInstructionTableEntry(table, 0b0010111, 0x0, instruction.AwooInstructionFormatU, "AUIPC", ProcessAUIPC)

	// Logical (12 instructions)
	SetupInstructionTableEntry(table, 0b0110011, 0x7, instruction.AwooInstructionFormatR, "AND", ProcessAND)
	SetupInstructionTableEntry(table, 0b0110011, 0x6, instruction.AwooInstructionFormatR, "OR", ProcessOR)
	SetupInstructionTableEntry(table, 0b0110011, 0x4, instruction.AwooInstructionFormatR, "XOR", ProcessXOR)
	SetupInstructionTableEntry(table, 0b0010011, 0x7, instruction.AwooInstructionFormatI, "ANDI", ProcessANDI)
	SetupInstructionTableEntry(table, 0b0110011, 0x6, instruction.AwooInstructionFormatI, "ORI", ProcessORI)
	SetupInstructionTableEntry(table, 0b0110011, 0x4, instruction.AwooInstructionFormatI, "XORI", ProcessXORI)
	SetupInstructionTableEntry(table, 0b0110011, 0x1, instruction.AwooInstructionFormatR, "SLL", ProcessSLL)
	SetupInstructionTableEntry(table, 0b0110011, 0x5, instruction.AwooInstructionFormatR, "SRL", ProcessSRL)
	SetupInstructionTableEntry(table, 0b0110011, 0x105, instruction.AwooInstructionFormatI, "SRA", ProcessSRA)
	SetupInstructionTableEntry(table, 0b0010011, 0x1, instruction.AwooInstructionFormatI, "SLLI", ProcessSLLI)
	SetupInstructionTableEntry(table, 0b0010011, 0x5, instruction.AwooInstructionFormatI, "SRLI", ProcessSRLI)
	SetupInstructionTableEntry(table, 0b0010011, 0x105, instruction.AwooInstructionFormatI, "SRAI", ProcessSRAI)

	// Load / Store (11 instructions)

	// Branching (8 instructions)

	return table
}
