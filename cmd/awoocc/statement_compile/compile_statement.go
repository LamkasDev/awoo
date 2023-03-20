package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/jwalton/gchalk"
)

func CompileStatement(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	entry, ok := ccompiler.Settings.Mappings.Statement[s.Type]
	if !ok {
		return fmt.Errorf("%w: %s", awerrors.ErrorCantCompileStatement, gchalk.Red(fmt.Sprintf("%#x", s.Type)))
	}

	return entry(ccompiler, elf, s)
}
