package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementAssignment(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	nameNode := statement.GetStatementAssignmentIdentifier(&s)
	name := node.GetNodeIdentifierValue(&nameNode)
	dest, err := compiler_context.GetCompilerScopeMemory(context, name)
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToGetVariableFromScope, err)
	}
	valueNode := statement.GetStatementAssignmentValue(&s)
	d, err = CompileNodeValueFast(context, valueNode, d)
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCompileNode, err)
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: *instruction.AwooInstructionsSave[context.Parser.Lexer.Types.All[dest.Type].Size],
		SourceTwo:   cpu.AwooRegisterTemporaryZero,
		Immediate:   uint32(dest.Start),
	}, d)
}
