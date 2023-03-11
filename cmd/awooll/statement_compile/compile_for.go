package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementFor(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	d, err := CompileStatement(ccompiler, statement.GetStatementForInitialization(&s), d)
	if err != nil {
		return d, err
	}
	initSize := len(d)
	d, err = CompileNodeValue(ccompiler, statement.GetStatementForCondition(&s), d, &compiler_details.CompileNodeValueDetails{
		Register: cpu.AwooRegisterTemporaryZero,
	})
	if err != nil {
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
	body, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJAL,
		Immediate:   uint32((-len(d) + initSize) - len(body)),
	}, body)
	if err != nil {
		return d, err
	}

	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionBEQ,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32((len(d) - initSize) + len(body)),
	}, d)
	if err != nil {
		return d, err
	}
	d = append(d, body...)

	return d, nil
}
