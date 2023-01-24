package print

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func PrintNewStatement(context *lexer_context.AwooLexerContext, s *statement.AwooParserStatement) {
	fmt.Printf("┗━━► %s \n", PrintStatement(context, s))
}

func PrintStatement(context *lexer_context.AwooLexerContext, s *statement.AwooParserStatement) string {
	switch s.Type {
	case statement.ParserStatementTypeDefinition:
		t := statement.GetStatementDefinitionType(s)
		id := statement.GetStatementDefinitionIdentifier(s)
		prim := statement.GetStatementDefinitionValue(s)
		return fmt.Sprintf("statement %s %s = %s (%s)", GetNodeDataText(context, &t), GetNodeDataText(context, &id), GetNodeDataText(context, &prim), gchalk.Green(fmt.Sprintf("%#x", s.Type)))
	case statement.ParserStatementTypeAssignment:
		id := statement.GetStatementAssignmentIdentifier(s)
		prim := statement.GetStatementAssignmentValue(s)
		return fmt.Sprintf("statement %s = %s (%s)", GetNodeDataText(context, &id), GetNodeDataText(context, &prim), gchalk.Green(fmt.Sprintf("%#x", s.Type)))
	}

	return "??"
}
