package statement_compile

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
		compiler_context.PushCompilerScopeCurrentBlock(context, "if")
		body, err = CompileStatementGroup(context, statement.GetStatementIfBody(&s), []byte{})
		if err != nil {
			return bodies, jump, err
		}
		compiler_context.PopCompilerScopeCurrentBlock(context)
		// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare).
		valueNode := statement.GetStatementIfValue(&s)
		details := compiler_context.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
		ifIns, err := CompileNodeValueFast(context, valueNode, []byte{}, &details)
		if err != nil {
			return bodies, jump, err
		}
		ifIns, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionBEQ,
			SourceOne:   details.Register,
			Immediate:   uint32(len(body) + 8),
		}, ifIns)
		if err != nil {
			return bodies, jump, err
		}
		body = append(ifIns, body...)
	case statement.ParserStatementTypeGroup:
		compiler_context.PushCompilerScopeCurrentBlock(context, "else")
		body, err = CompileStatementGroup(context, s, []byte{})
		if err != nil {
			return bodies, jump, err
		}
		compiler_context.PopCompilerScopeCurrentBlock(context)
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
