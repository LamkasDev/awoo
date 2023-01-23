package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/instruction"
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
		Instruction: instruction.AwooInstruction{Code: code, Format: format, Name: name, Advance: true},
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
	SetupInstructionTableEntry(table, 0b0110011, 0x6, instruction.AwooInstructionFormatR, "SLTU", ProcessSLTU)
	SetupInstructionTableEntry(table, 0b0010011, 0x6, instruction.AwooInstructionFormatI, "SLTIU", ProcessSLTIU)
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
	// SetupInstructionTableEntry(table, 0b0000011, 0x3, instruction.AwooInstructionFormatI, "LD", ProcessLD)
	SetupInstructionTableEntry(table, 0b0000011, 0x2, instruction.AwooInstructionFormatI, "LW", ProcessLW)
	SetupInstructionTableEntry(table, 0b0000011, 0x1, instruction.AwooInstructionFormatI, "LH", ProcessLH)
	SetupInstructionTableEntry(table, 0b0000011, 0x0, instruction.AwooInstructionFormatI, "LB", ProcessLB)
	// SetupInstructionTableEntry(table, 0b0000011, 0x6, instruction.AwooInstructionFormatI, "LWU", ProcessLWU)
	SetupInstructionTableEntry(table, 0b0000011, 0x5, instruction.AwooInstructionFormatI, "LHU", ProcessLHU)
	SetupInstructionTableEntry(table, 0b0000011, 0x4, instruction.AwooInstructionFormatI, "LBU", ProcessLBU)
	// SetupInstructionTableEntry(table, 0b0100011, 0x3, instruction.AwooInstructionFormatS, "SD", ProcessSD)
	SetupInstructionTableEntry(table, 0b0100011, 0x2, instruction.AwooInstructionFormatS, "SW", ProcessSW)
	SetupInstructionTableEntry(table, 0b0100011, 0x1, instruction.AwooInstructionFormatS, "SH", ProcessSH)
	SetupInstructionTableEntry(table, 0b0100011, 0x0, instruction.AwooInstructionFormatS, "SB", ProcessSB)

	// Branching (8 instructions)
	SetupInstructionTableEntry(table, 0b1100011, 0x0, instruction.AwooInstructionFormatB, "BEQ", ProcessBEQ)
	SetupInstructionTableEntry(table, 0b1100011, 0x1, instruction.AwooInstructionFormatB, "BNE", ProcessBNE)
	SetupInstructionTableEntry(table, 0b1100011, 0x5, instruction.AwooInstructionFormatB, "BGE", ProcessBGE)
	SetupInstructionTableEntry(table, 0b1100011, 0x7, instruction.AwooInstructionFormatB, "BGEU", ProcessBGEU)
	SetupInstructionTableEntry(table, 0b1100011, 0x4, instruction.AwooInstructionFormatB, "BLT", ProcessBLT)
	SetupInstructionTableEntry(table, 0b1100011, 0x6, instruction.AwooInstructionFormatB, "BLTU", ProcessBLTU)
	SetupInstructionTableEntry(table, 0b1101111, 0x0, instruction.AwooInstructionFormatJ, "JAL", ProcessJAL)
	SetupInstructionTableEntry(table, 0b1100111, 0x0, instruction.AwooInstructionFormatI, "JALR", ProcessJALR)

	return table
}
