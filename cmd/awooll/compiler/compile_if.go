package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementIfNode(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, bodies [][]byte, jump uint32) ([][]byte, uint32, error) {
	var body []byte
	var err error
	switch s.Type {
	case statement.ParserStatementTypeIf:
		body, err = CompileStatementGroup(context, statement.GetStatementIfBody(&s), []byte{})
		if err != nil {
			return bodies, jump, err
		}
		// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare)
		valueNode := statement.GetStatementIfValue(&s)
		ifIns, err := CompileNodeValueFast(context, valueNode, []byte{})
		if err != nil {
			return bodies, jump, err
		}
		ifIns, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionBEQ,
			SourceOne:   cpu.AwooRegisterTemporaryZero,
			Immediate:   uint32(len(body) + 8),
		}, ifIns)
		if err != nil {
			return bodies, jump, err
		}
		body = append(ifIns, body...)
	case statement.ParserStatementTypeGroup:
		body, err = CompileStatementGroup(context, s, []byte{})
		if err != nil {
			return bodies, jump, err
		}
	}
	bodies = append(bodies, body)
	jump += uint32(len(body) + 4)

	return bodies, jump, nil
}

func CompileStatementIf(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	bodies, jump, err := CompileStatementIfNode(context, s, [][]byte{}, uint32(4))
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

	for i, body := range bodies {
		// TODO: skip jump if last in chain
		jump -= uint32(len(body) + 4)
		if jump <= 4 {
			continue
		}
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
