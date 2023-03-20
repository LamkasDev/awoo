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
	initSize := len(d)

	conditionDetails := compiler_details.CompileNodeValueDetails{
		Type:     commonTypes.AwooTypeId(types.AwooTypeBoolean),
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if err := CompileNodeValue(ccompiler, elf, statement.GetStatementForCondition(&s), &conditionDetails); err != nil {
		return err
	}

	body, err := CompileStatementGroup(ccompiler, statement.GetStatementForBody(&s), []byte{})
	if err != nil {
		return err
	}

	body, err = CompileStatement(ccompiler, elf, statement.GetStatementForAdvancement(&s))
	if err != nil {
		return err
	}

	jumpToConditionInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJAL,
		Immediate:   uint32((-len(d) + initSize) - len(body)),
	}
	body, err = encoder.Encode(jumpToConditionInstruction, body)
	if err != nil {
		return d, err
	}

	jumpBeyondEndInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionBEQ,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32((len(d) - initSize) + len(body)),
	}
	if err = encoder.Encode(elf, jumpBeyondEndInstruction); err != nil {
		return err
	}

	return nil
}
