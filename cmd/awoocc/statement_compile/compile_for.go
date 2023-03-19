package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func CompileStatementFor(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	var err error
	if d, err = CompileStatement(ccompiler, statement.GetStatementForInitialization(&s), d); err != nil {
		return d, err
	}
	initSize := len(d)

	conditionDetails := compiler_details.CompileNodeValueDetails{
		Type:     commonTypes.AwooTypeId(types.AwooTypeBoolean),
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if d, err = CompileNodeValue(ccompiler, statement.GetStatementForCondition(&s), d, &conditionDetails); err != nil {
		return d, err
	}

	body, err := CompileStatementGroup(ccompiler, statement.GetStatementForBody(&s), []byte{})
	if err != nil {
		return d, err
	}

	body, err = CompileStatement(ccompiler, statement.GetStatementForAdvancement(&s), body)
	if err != nil {
		return d, err
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
	if d, err = encoder.Encode(jumpBeyondEndInstruction, d); err != nil {
		return d, err
	}

	d = append(d, body...)
	return d, nil
}
