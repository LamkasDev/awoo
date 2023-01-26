package instruction

type AwooInstructionTableEntry struct {
	Instruction AwooInstruction
	Data        interface{}
}
type AwooInstructionTableSubtable struct {
	Format   uint8
	Subtable map[uint16]AwooInstructionTableEntry
}
type AwooInstructionTable map[uint8]AwooInstructionTableSubtable

func SetupInstructionTableEntry(table AwooInstructionTable, instruction AwooInstruction, data interface{}) {
	_, ok := table[instruction.Code]
	if !ok {
		table[instruction.Code] = AwooInstructionTableSubtable{
			Format:   instruction.Format,
			Subtable: make(map[uint16]AwooInstructionTableEntry),
		}
	}

	table[instruction.Code].Subtable[instruction.Argument] = AwooInstructionTableEntry{
		Instruction: instruction,
		Data:        data,
	}
}

func DecorateInstructionTableEntry(table AwooInstructionTable, code uint8, arg uint16, data interface{}) {
	e := table[code].Subtable[arg]
	e.Data = data
	table[code].Subtable[arg] = e
}

func SetupInstructionTable() AwooInstructionTable {
	table := AwooInstructionTable{}

	// Arithmetic (9 instructions)
	SetupInstructionTableEntry(table, AwooInstructionADD, nil)
	SetupInstructionTableEntry(table, AwooInstructionSUB, nil)
	SetupInstructionTableEntry(table, AwooInstructionADDI, nil)
	SetupInstructionTableEntry(table, AwooInstructionSLT, nil)
	SetupInstructionTableEntry(table, AwooInstructionSLTI, nil)
	SetupInstructionTableEntry(table, AwooInstructionSLTU, nil)
	SetupInstructionTableEntry(table, AwooInstructionSLTIU, nil)
	SetupInstructionTableEntry(table, AwooInstructionLUI, nil)
	SetupInstructionTableEntry(table, AwooInstructionAUIPC, nil)

	// Logical (12 instructions)
	SetupInstructionTableEntry(table, AwooInstructionAND, nil)
	SetupInstructionTableEntry(table, AwooInstructionOR, nil)
	SetupInstructionTableEntry(table, AwooInstructionXOR, nil)
	SetupInstructionTableEntry(table, AwooInstructionANDI, nil)
	SetupInstructionTableEntry(table, AwooInstructionORI, nil)
	SetupInstructionTableEntry(table, AwooInstructionXORI, nil)
	SetupInstructionTableEntry(table, AwooInstructionSLL, nil)
	SetupInstructionTableEntry(table, AwooInstructionSRL, nil)
	SetupInstructionTableEntry(table, AwooInstructionSRA, nil)
	SetupInstructionTableEntry(table, AwooInstructionSLLI, nil)
	SetupInstructionTableEntry(table, AwooInstructionSRLI, nil)
	SetupInstructionTableEntry(table, AwooInstructionSRAI, nil)

	// Load / Store (11 instructions)
	// SetupInstructionTableEntry(table, AwooInstructionLD, nil)
	SetupInstructionTableEntry(table, AwooInstructionLW, nil)
	SetupInstructionTableEntry(table, AwooInstructionLH, nil)
	SetupInstructionTableEntry(table, AwooInstructionLB, nil)
	// SetupInstructionTableEntry(table, AwooInstructionLWU, nil)
	SetupInstructionTableEntry(table, AwooInstructionLHU, nil)
	SetupInstructionTableEntry(table, AwooInstructionLBU, nil)
	// SetupInstructionTableEntry(table, AwooInstructionSD, nil)
	SetupInstructionTableEntry(table, AwooInstructionSW, nil)
	SetupInstructionTableEntry(table, AwooInstructionSH, nil)
	SetupInstructionTableEntry(table, AwooInstructionSB, nil)

	// Branching (8 instructions)
	SetupInstructionTableEntry(table, AwooInstructionBEQ, nil)
	SetupInstructionTableEntry(table, AwooInstructionBNE, nil)
	SetupInstructionTableEntry(table, AwooInstructionBGE, nil)
	SetupInstructionTableEntry(table, AwooInstructionBGEU, nil)
	SetupInstructionTableEntry(table, AwooInstructionBLT, nil)
	SetupInstructionTableEntry(table, AwooInstructionBLTU, nil)
	SetupInstructionTableEntry(table, AwooInstructionJAL, nil)
	SetupInstructionTableEntry(table, AwooInstructionJALR, nil)

	// Multiply / Divide extension
	SetupInstructionTableEntry(table, AwooInstructionMUL, nil)
	SetupInstructionTableEntry(table, AwooInstructionDIV, nil)

	return table
}
