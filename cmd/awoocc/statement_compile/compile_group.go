package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func CompileStatementGroup(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	for _, n := range statement.GetStatementGroupBody(&s) {
		if err := CompileStatement(ccompiler, celf, n); err != nil {
			return err
		}
	}

	return nil
}
