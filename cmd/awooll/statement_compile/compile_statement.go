package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func CompileStatement(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	entry, ok := context.MappingsStatement[s.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileStatement, gchalk.Red(fmt.Sprintf("%#x", s.Type)))
	}
	d, err := entry(context, s, d)
	if err != nil {
		return d, err
	}

	return d, nil
}
