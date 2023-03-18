package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementIfNode(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, bodies [][]byte, jump uint32) ([][]byte, uint32, error) {
	nodeBody := []byte{}
	var err error

	switch s.Type {
	case statement.ParserStatementTypeIf:
		// We need to compile the body first to determine comparison jump address.
		compiler_context.PushCompilerScopeCurrentBlock(&ccompiler.Context, compiler_context.AwooCompilerScopeBlock{
			Name: "if",
		})
		nodeBody, err = CompileStatementGroup(ccompiler, statement.GetStatementIfBody(&s), nodeBody)
		if err != nil {
			return bodies, jump, err
		}
		compiler_context.PopCompilerScopeCurrentBlock(&ccompiler.Context)

		// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare).
		valueNode := statement.GetStatementIfValue(&s)
		valueDetails := compiler_details.CompileNodeValueDetails{
			Type:     types.AwooTypeBoolean,
			Register: cpu.AwooRegisterTemporaryZero,
		}
		ifHeader, err := CompileNodeValue(ccompiler, valueNode, []byte{}, &valueDetails)
		if err != nil {
			return bodies, jump, err
		}

		jumpBeyondEndInstruction := encoder.AwooEncodedInstruction{
			Instruction: instructions.AwooInstructionBEQ,
			SourceOne:   valueDetails.Register,
			Immediate:   uint32(len(nodeBody) + 8),
		}
		ifHeader, err = encoder.Encode(jumpBeyondEndInstruction, ifHeader)
		if err != nil {
			return bodies, jump, err
		}

		nodeBody = append(ifHeader, nodeBody...)
	case statement.ParserStatementTypeGroup:
		compiler_context.PushCompilerScopeCurrentBlock(&ccompiler.Context, compiler_context.AwooCompilerScopeBlock{
			Name: "else",
		})
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
		jumpToNextBlockInstruction := encoder.AwooEncodedInstruction{
			Instruction: instructions.AwooInstructionJAL,
			Immediate:   jump,
		}
		if bodies[i], err = encoder.Encode(jumpToNextBlockInstruction, bodies[i]); err != nil {
			return d, err
		}
	}

	for _, b := range bodies {
		d = append(d, b...)
	}
	return d, nil
}
