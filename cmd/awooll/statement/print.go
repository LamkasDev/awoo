package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintNewStatement(context *lexer_context.AwooLexerContext, s *AwooParserStatement) {
	logger.Log("┗━━► %s \n", PrintStatement(context, s))
}

func PrintStatement(context *lexer_context.AwooLexerContext, s *AwooParserStatement) string {
	tText := gchalk.Green(fmt.Sprintf("%#x", s.Type))
	switch s.Type {
	case ParserStatementTypeDefinitionVariable:
		t := GetStatementDefinitionVariableType(s)
		id := GetStatementDefinitionVariableIdentifier(s)
		prim := GetStatementDefinitionVariableValue(s)
		return fmt.Sprintf("statement %s %s = %s (%s)", node.GetNodeDataText(context, &t), node.GetNodeDataText(context, &id), node.GetNodeDataText(context, &prim), tText)
	case ParserStatementTypeAssignment:
		id := GetStatementAssignmentIdentifier(s)
		prim := GetStatementAssignmentValue(s)
		return fmt.Sprintf("statement %s = %s (%s)", node.GetNodeDataText(context, &id), node.GetNodeDataText(context, &prim), tText)
	case ParserStatementTypeDefinitionType:
		id := GetStatementDefinitionTypeIdentifier(s)
		value := GetStatementDefinitionTypeValue(s)
		return fmt.Sprintf("type %s = %s (%s)", node.GetNodeDataText(context, &id), node.GetNodeDataText(context, &value), tText)
	case ParserStatementTypeIf:
		value := GetStatementIfValue(s)
		return fmt.Sprintf("if %s (%s)", node.GetNodeDataText(context, &value), tText)
	case ParserStatementTypeFunc:
		id := GetStatementFuncIdentifier(s)
		return fmt.Sprintf("func %s (%s)", node.GetNodeDataText(context, &id), tText)
	}

	return "??"
}
