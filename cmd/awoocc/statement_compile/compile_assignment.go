package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

func CompileStatementAssignment(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	switch identifierNode.Type {
	case node.ParserNodeTypePointer:
		return CompileStatementAssignmentPointer(ccompiler, s, d)
	case node.ParserNodeTypeArrayIndex:
		return CompileStatementAssignmentArrayIndex(ccompiler, s, d)
	}

	return CompileStatementAssignmentIdentifier(ccompiler, s, d)
}
