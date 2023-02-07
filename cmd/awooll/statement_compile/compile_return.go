package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementReturn(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	d, err := CompileNodeValue(context, statement.GetStatementReturnValue(&s), d, &compiler_context.CompileNodeValueDetails{
		Register: cpu.AwooRegisterFunctionOne,
	})
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCompileStatement, err)
	}

	// TODO: this is retarder.
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJALR,
		SourceOne:   cpu.AwooRegisterStackPointer,
	}, d)
}
