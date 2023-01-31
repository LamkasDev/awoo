package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/jwalton/gchalk"
)

func PrintNewStatement(context *lexer_context.AwooLexerContext, s *AwooParserStatement) {
	fmt.Printf("┗━━► %s \n", PrintStatement(context, s))
}

func PrintStatement(context *lexer_context.AwooLexerContext, s *AwooParserStatement) string {
	switch s.Type {
	case ParserStatementTypeDefinition:
		t := GetStatementDefinitionType(s)
		id := GetStatementDefinitionIdentifier(s)
		prim := GetStatementDefinitionValue(s)
		return fmt.Sprintf("statement %s %s = %s (%s)", node.GetNodeDataText(context, &t), node.GetNodeDataText(context, &id), node.GetNodeDataText(context, &prim), gchalk.Green(fmt.Sprintf("%#x", s.Type)))
	case ParserStatementTypeAssignment:
		id := GetStatementAssignmentIdentifier(s)
		prim := GetStatementAssignmentValue(s)
		return fmt.Sprintf("statement %s = %s (%s)", node.GetNodeDataText(context, &id), node.GetNodeDataText(context, &prim), gchalk.Green(fmt.Sprintf("%#x", s.Type)))
	}

	return "??"
}
