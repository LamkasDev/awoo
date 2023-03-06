package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintNewStatement(context *parser_context.AwooParserContext, s *statement.AwooParserStatement) {
	logger.Log("┗━━► %s \n", PrintStatement(context, s))
}

func PrintStatement(context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	tText := gchalk.Green(fmt.Sprintf("%#x", s.Type))
	switch s.Type {
	case statement.ParserStatementTypeDefinitionVariable:
		t := statement.GetStatementDefinitionVariableType(s)
		id := statement.GetStatementDefinitionVariableIdentifier(s)
		prim := statement.GetStatementDefinitionVariableValue(s)
		return fmt.Sprintf("statement %s %s = %s (%s)", lexer.PrintNode(context.Lexer, &t), lexer.PrintNode(context.Lexer, &id), lexer.PrintNode(context.Lexer, &prim), tText)
	case statement.ParserStatementTypeAssignment:
		id := statement.GetStatementAssignmentIdentifier(s)
		prim := statement.GetStatementAssignmentValue(s)
		return fmt.Sprintf("statement %s = %s (%s)", lexer.PrintNode(context.Lexer, &id), lexer.PrintNode(context.Lexer, &prim), tText)
	case statement.ParserStatementTypeDefinitionType:
		id := statement.GetStatementDefinitionTypeIdentifier(s)
		value := statement.GetStatementDefinitionTypeValue(s)
		return fmt.Sprintf("type %s = %s (%s)", lexer.PrintNode(context.Lexer, &id), lexer.PrintNode(context.Lexer, &value), tText)
	case statement.ParserStatementTypeIf:
		value := statement.GetStatementIfValue(s)
		return fmt.Sprintf("if %s (%s)", lexer.PrintNode(context.Lexer, &value), tText)
	case statement.ParserStatementTypeFunc:
		id := statement.GetStatementFuncIdentifier(s)
		return fmt.Sprintf("func %s (%s)", lexer.PrintNode(context.Lexer, &id), tText)
	}

	return "??"
}
