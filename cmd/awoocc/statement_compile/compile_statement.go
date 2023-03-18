package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/jwalton/gchalk"
)

func CompileStatement(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	entry, ok := ccompiler.Settings.Mappings.Statement[s.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileStatement, gchalk.Red(fmt.Sprintf("%#x", s.Type)))
	}

	return entry(ccompiler, s, d)
}
