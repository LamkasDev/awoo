package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func CompileStatementFor(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	if err := CompileStatement(ccompiler, elf, statement.GetStatementForInitialization(&s)); err != nil {
		return err
	}

	// Reserve instruction for condition.
	conditionStart := uint32(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents))
	if err := encoder.Encode(elf, encoder.AwooEncodedInstruction{}); err != nil {
		return err
	}

	conditionDetails := compiler_details.CompileNodeValueDetails{
		Type:     commonTypes.AwooTypeId(types.AwooTypeBoolean),
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if err := CompileNodeValue(ccompiler, elf, statement.GetStatementForCondition(&s), &conditionDetails); err != nil {
		return err
	}

	if err := CompileStatementGroup(ccompiler, elf, statement.GetStatementForBody(&s)); err != nil {
		return err
	}

	if err := CompileStatement(ccompiler, elf, statement.GetStatementForAdvancement(&s)); err != nil {
		return err
	}

	end := uint32(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents))
	jumpToConditionInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJAL,
		Immediate:   uint32(conditionStart - end),
	}
	if err := encoder.Encode(elf, jumpToConditionInstruction); err != nil {
		return err
	}

	end = uint32(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents))
	jumpBeyondEndInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionBEQ,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(end - conditionStart),
	}
	if err := encoder.EncodeAt(elf, conditionStart, jumpBeyondEndInstruction); err != nil {
		return err
	}

	return nil
}
