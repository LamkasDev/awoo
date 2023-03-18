package instructions

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func SetupDecoderInstructionTable() instructions.AwooInstructionTable {
	table := instructions.SetupInstructionTable()

	// Arithmetic (9 instructions)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionADD.Code, instructions.AwooInstructionADD.Argument, ProcessADD)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSUB.Code, instructions.AwooInstructionSUB.Argument, ProcessSUB)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionADDI.Code, instructions.AwooInstructionADDI.Argument, ProcessADDI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSLT.Code, instructions.AwooInstructionSLT.Argument, ProcessSLT)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSLTI.Code, instructions.AwooInstructionSLTI.Argument, ProcessSLTI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSLTU.Code, instructions.AwooInstructionSLTU.Argument, ProcessSLTU)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSLTIU.Code, instructions.AwooInstructionSLTIU.Argument, ProcessSLTIU)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLUI.Code, instructions.AwooInstructionLUI.Argument, ProcessLUI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionAUIPC.Code, instructions.AwooInstructionAUIPC.Argument, ProcessAUIPC)

	// Logical (12 instructions)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionAND.Code, instructions.AwooInstructionAND.Argument, ProcessAND)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionOR.Code, instructions.AwooInstructionOR.Argument, ProcessOR)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionXOR.Code, instructions.AwooInstructionXOR.Argument, ProcessXOR)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionANDI.Code, instructions.AwooInstructionANDI.Argument, ProcessANDI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionORI.Code, instructions.AwooInstructionORI.Argument, ProcessORI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionXORI.Code, instructions.AwooInstructionXORI.Argument, ProcessXORI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSLL.Code, instructions.AwooInstructionSLL.Argument, ProcessSLL)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSRL.Code, instructions.AwooInstructionSRL.Argument, ProcessSRL)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSRA.Code, instructions.AwooInstructionSRA.Argument, ProcessSRA)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSLLI.Code, instructions.AwooInstructionSLLI.Argument, ProcessSLLI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSRLI.Code, instructions.AwooInstructionSRLI.Argument, ProcessSRLI)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSRAI.Code, instructions.AwooInstructionSRAI.Argument, ProcessSRAI)

	// Load / Store (11 instructions)
	// instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLD.Code, instructions.AwooInstructionLD.Argument, ProcessLD)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLW.Code, instructions.AwooInstructionLW.Argument, ProcessLW)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLH.Code, instructions.AwooInstructionLH.Argument, ProcessLH)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLB.Code, instructions.AwooInstructionLB.Argument, ProcessLB)
	// instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLWU.Code, instructions.AwooInstructionLWU.Argument, ProcessLWU)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLHU.Code, instructions.AwooInstructionLHU.Argument, ProcessLHU)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionLBU.Code, instructions.AwooInstructionLBU.Argument, ProcessLBU)
	// instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSD.Code, instructions.AwooInstructionSD.Argument, ProcessSD)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSW.Code, instructions.AwooInstructionSW.Argument, ProcessSW)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSH.Code, instructions.AwooInstructionSH.Argument, ProcessSH)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionSB.Code, instructions.AwooInstructionSB.Argument, ProcessSB)

	// Branching (8 instructions)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionBEQ.Code, instructions.AwooInstructionBEQ.Argument, ProcessBEQ)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionBNE.Code, instructions.AwooInstructionBNE.Argument, ProcessBNE)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionBGE.Code, instructions.AwooInstructionBGE.Argument, ProcessBGE)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionBGEU.Code, instructions.AwooInstructionBGEU.Argument, ProcessBGEU)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionBLT.Code, instructions.AwooInstructionBLT.Argument, ProcessBLT)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionBLTU.Code, instructions.AwooInstructionBLTU.Argument, ProcessBLTU)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionJAL.Code, instructions.AwooInstructionJAL.Argument, ProcessJAL)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionJALR.Code, instructions.AwooInstructionJALR.Argument, ProcessJALR)

	// Multiply / Divide extension
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionMUL.Code, instructions.AwooInstructionMUL.Argument, ProcessMUL)
	instructions.DecorateInstructionTableEntry(table, instructions.AwooInstructionDIV.Code, instructions.AwooInstructionDIV.Argument, ProcessDIV)

	return table
}
