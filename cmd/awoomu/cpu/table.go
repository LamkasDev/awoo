package cpu

import "github.com/LamkasDev/awoo-emu/cmd/common/instruction"

func SetupDecoderInstructionTable() instruction.AwooInstructionTable {
	table := instruction.SetupInstructionTable()

	// Arithmetic (9 instructions)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionADD.Code, instruction.AwooInstructionADD.Argument, ProcessADD)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSUB.Code, instruction.AwooInstructionSUB.Argument, ProcessSUB)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionADDI.Code, instruction.AwooInstructionADDI.Argument, ProcessADDI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSLT.Code, instruction.AwooInstructionSLT.Argument, ProcessSLT)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSLTI.Code, instruction.AwooInstructionSLTI.Argument, ProcessSLTI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSLTU.Code, instruction.AwooInstructionSLTU.Argument, ProcessSLTU)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSLTIU.Code, instruction.AwooInstructionSLTIU.Argument, ProcessSLTIU)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLUI.Code, instruction.AwooInstructionLUI.Argument, ProcessLUI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionAUIPC.Code, instruction.AwooInstructionAUIPC.Argument, ProcessAUIPC)

	// Logical (12 instructions)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionAND.Code, instruction.AwooInstructionAND.Argument, ProcessAND)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionOR.Code, instruction.AwooInstructionOR.Argument, ProcessOR)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionXOR.Code, instruction.AwooInstructionXOR.Argument, ProcessXOR)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionANDI.Code, instruction.AwooInstructionANDI.Argument, ProcessANDI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionORI.Code, instruction.AwooInstructionORI.Argument, ProcessORI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionXORI.Code, instruction.AwooInstructionXORI.Argument, ProcessXORI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSLL.Code, instruction.AwooInstructionSLL.Argument, ProcessSLL)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSRL.Code, instruction.AwooInstructionSRL.Argument, ProcessSRL)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSRA.Code, instruction.AwooInstructionSRA.Argument, ProcessSRA)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSLLI.Code, instruction.AwooInstructionSLLI.Argument, ProcessSLLI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSRLI.Code, instruction.AwooInstructionSRLI.Argument, ProcessSRLI)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSRAI.Code, instruction.AwooInstructionSRAI.Argument, ProcessSRAI)

	// Load / Store (11 instructions)
	// instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLD.Code, instruction.AwooInstructionLD.Argument, ProcessLD)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLW.Code, instruction.AwooInstructionLW.Argument, ProcessLW)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLH.Code, instruction.AwooInstructionLH.Argument, ProcessLH)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLB.Code, instruction.AwooInstructionLB.Argument, ProcessLB)
	// instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLWU.Code, instruction.AwooInstructionLWU.Argument, ProcessLWU)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLHU.Code, instruction.AwooInstructionLHU.Argument, ProcessLHU)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionLBU.Code, instruction.AwooInstructionLBU.Argument, ProcessLBU)
	// instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSD.Code, instruction.AwooInstructionSD.Argument, ProcessSD)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSW.Code, instruction.AwooInstructionSW.Argument, ProcessSW)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSH.Code, instruction.AwooInstructionSH.Argument, ProcessSH)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionSB.Code, instruction.AwooInstructionSB.Argument, ProcessSB)

	// Branching (8 instructions)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionBEQ.Code, instruction.AwooInstructionBEQ.Argument, ProcessBEQ)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionBNE.Code, instruction.AwooInstructionBNE.Argument, ProcessBNE)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionBGE.Code, instruction.AwooInstructionBGE.Argument, ProcessBGE)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionBGEU.Code, instruction.AwooInstructionBGEU.Argument, ProcessBGEU)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionBLT.Code, instruction.AwooInstructionBLT.Argument, ProcessBLT)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionBLTU.Code, instruction.AwooInstructionBLTU.Argument, ProcessBLTU)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionJAL.Code, instruction.AwooInstructionJAL.Argument, ProcessJAL)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionJALR.Code, instruction.AwooInstructionJALR.Argument, ProcessJALR)

	// Multiply / Divide extension
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionMUL.Code, instruction.AwooInstructionMUL.Argument, ProcessMUL)
	instruction.DecorateInstructionTableEntry(table, instruction.AwooInstructionDIV.Code, instruction.AwooInstructionDIV.Argument, ProcessDIV)

	return table
}
