package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func CompileStatementFor(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	if err := CompileStatement(ccompiler, celf, statement.GetStatementForInitialization(&s)); err != nil {
		return err
	}

	// Reserve instruction for condition.
	conditionStart := arch.AwooRegister(len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents))
	if err := encoder.Encode(celf, instruction.AwooInstruction{}); err != nil {
		return err
	}

	conditionDetails := compiler_details.CompileNodeValueDetails{
		Type:     commonTypes.AwooTypeId(types.AwooTypeBoolean),
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if err := CompileNodeValue(ccompiler, celf, statement.GetStatementForCondition(&s), &conditionDetails); err != nil {
		return err
	}

	if err := CompileStatementGroup(ccompiler, celf, statement.GetStatementForBody(&s)); err != nil {
		return err
	}

	if err := CompileStatement(ccompiler, celf, statement.GetStatementForAdvancement(&s)); err != nil {
		return err
	}

	end := arch.AwooRegister(len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents))
	jumpToConditionInstruction := instruction.AwooInstruction{
		Definition: instructions.AwooInstructionJAL,
		Immediate:  conditionStart - end,
	}
	if err := encoder.Encode(celf, jumpToConditionInstruction); err != nil {
		return err
	}

	end = arch.AwooRegister(len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents))
	jumpBeyondEndInstruction := instruction.AwooInstruction{
		Definition: instructions.AwooInstructionBEQ,
		SourceOne:  cpu.AwooRegisterTemporaryZero,
		Immediate:  end - conditionStart,
	}
	if err := encoder.EncodeAt(celf, conditionStart, jumpBeyondEndInstruction); err != nil {
		return err
	}

	return nil
}
