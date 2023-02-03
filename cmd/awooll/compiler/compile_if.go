package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementIfNode(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, bodies [][]byte, jump uint32) ([][]byte, uint32, error) {
	body, err := CompileStatement(context, s, []byte{})
	if err != nil {
		return bodies, jump, err
	}
	bodies = append(bodies, body)
	jump += uint32(len(body) + 4)

	return bodies, jump, nil
}

func CompileStatementIf(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	valueNode := statement.GetStatementIfValue(&s)
	d, err := CompileNodeValueFast(context, valueNode, d)
	if err != nil {
		return d, err
	}

	bodyGroup := statement.GetStatementIfBody(&s)
	bodies, jump, err := CompileStatementIfNode(context, bodyGroup, [][]byte{}, uint32(4))
	if err != nil {
		return d, err
	}
	nextGroups := statement.GetStatementIfNext(&s)
	for _, nextGroup := range nextGroups {
		bodies, jump, err = CompileStatementIfNode(context, nextGroup, bodies, jump)
		if err != nil {
			return d, err
		}
	}

	// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare 1 value)
	// TODO: skip to the next else
	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionBEQ,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(len(bodies[len(bodies)-1]) + 8),
	}, d)
	if err != nil {
		return d, err
	}

	for i, body := range bodies {
		// TODO: skip jump if last in chain
		jump -= uint32(len(body) + 4)
		bodies[i], err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionJAL,
			Destination: cpu.AwooRegisterZero,
			Immediate:   jump,
		}, bodies[i])
		if err != nil {
			return d, err
		}
	}
	for _, b := range bodies {
		d = append(d, b...)
	}

	return d, nil
}
