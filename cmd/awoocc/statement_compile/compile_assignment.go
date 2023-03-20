package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func CompileStatementAssignment(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	switch identifierNode.Type {
	case node.ParserNodeTypePointer:
		return CompileStatementAssignmentPointer(ccompiler, elf, s)
	case node.ParserNodeTypeArrayIndex:
		return CompileStatementAssignmentArrayIndex(ccompiler, elf, s)
	}

	return CompileStatementAssignmentIdentifier(ccompiler, elf, s)
}
