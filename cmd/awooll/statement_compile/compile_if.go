package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementIfNode(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, bodies [][]byte, jump uint32) ([][]byte, uint32, error) {
	details := compiler_details.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	nodeBody := []byte{}
	var err error

	switch s.Type {
	case statement.ParserStatementTypeIf:
		// We need to compile the body first to determine comparison jump address.
		compiler_context.PushCompilerScopeCurrentBlock(&ccompiler.Context, "if")
		nodeBody, err = CompileStatementGroup(ccompiler, statement.GetStatementIfBody(&s), nodeBody)
		if err != nil {
			return bodies, jump, err
		}
		compiler_context.PopCompilerScopeCurrentBlock(&ccompiler.Context)

		// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare).
		valueNode := statement.GetStatementIfValue(&s)
		ifHeader, err := CompileNodeValue(ccompiler, valueNode, []byte{}, &details)
		if err != nil {
			return bodies, jump, err
		}
		ifHeader, err = encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: instruction.AwooInstructionBEQ,
			SourceOne:   details.Register,
			Immediate:   uint32(len(nodeBody) + 8),
		}, ifHeader)
		if err != nil {
			return bodies, jump, err
		}
		nodeBody = append(ifHeader, nodeBody...)
	case statement.ParserStatementTypeGroup:
		compiler_context.PushCompilerScopeCurrentBlock(&ccompiler.Context, "else")
		nodeBody, err = CompileStatementGroup(ccompiler, s, []byte{})
		if err != nil {
			return bodies, jump, err
		}
		compiler_context.PopCompilerScopeCurrentBlock(&ccompiler.Context)
	}
	bodies = append(bodies, nodeBody)
	jump += uint32(len(nodeBody) + 4)

	return bodies, jump, nil
}

func CompileStatementIf(_ *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	bodies, jump, err := CompileStatementIfNode(&compiler.AwooCompiler{}, s, [][]byte{}, uint32(4))
	if err != nil {
		return d, err
	}
	elseGroups := statement.GetStatementIfElse(&s)
	for _, nextGroup := range elseGroups {
		bodies, jump, err = CompileStatementIfNode(&compiler.AwooCompiler{}, nextGroup, bodies, jump)
		if err != nil {
			return d, err
		}
	}

	// This is used for jumping over subsequent if statements.
	// An extra instruction is added on end of each block, except for the last.
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
