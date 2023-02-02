package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementIf(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	valueNode := statement.GetStatementIfValue(&s)
	d, err := CompileNodeValueFast(context, valueNode, d)
	if err != nil {
		return d, err
	}

	body := []byte{}
	bodyNode := statement.GetStatementIfBody(&s)
	for _, statement := range bodyNode {
		body, err = CompileStatement(context, statement, body)
	}

	// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare 1 value)
	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionBEQ,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(len(body) + 4),
	}, d)
	if err != nil {
		return d, err
	}
	return append(d, body...), nil
}
