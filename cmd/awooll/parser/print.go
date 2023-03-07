package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintStatementDefinitionVariable(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	t := statement.GetStatementDefinitionVariableType(s)
	id := statement.GetStatementDefinitionVariableIdentifier(s)
	prim := statement.GetStatementDefinitionVariableValue(s)
	return fmt.Sprintf("statement %s %s = %s (%s)",
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &t),
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &id),
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &prim),
		gchalk.Green(fmt.Sprintf("%#x", s.Type)),
	)
}

func PrintStatementAssignment(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	id := statement.GetStatementAssignmentIdentifier(s)
	prim := statement.GetStatementAssignmentValue(s)
	return fmt.Sprintf("statement %s = %s (%s)",
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &id),
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &prim),
		gchalk.Green(fmt.Sprintf("%#x", s.Type)),
	)
}

func PrintStatementDefinitionType(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	id := statement.GetStatementDefinitionTypeIdentifier(s)
	value := statement.GetStatementDefinitionTypeValue(s)
	return fmt.Sprintf("type %s = %s (%s)",
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &id),
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &value),
		gchalk.Green(fmt.Sprintf("%#x", s.Type)),
	)
}

func PrintStatementIf(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	value := statement.GetStatementIfValue(s)
	return fmt.Sprintf("if %s (%s)",
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &value),
		gchalk.Green(fmt.Sprintf("%#x", s.Type)),
	)
}

func PrintStatementFunc(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	id := statement.GetStatementFuncIdentifier(s)
	return fmt.Sprintf("func %s (%s)",
		lexer.PrintNode(&settings.Lexer, &context.Lexer, &id),
		gchalk.Green(fmt.Sprintf("%#x", s.Type)),
	)
}

func PrintNewStatement(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) {
	logger.Log("┗━━► %s \n", PrintStatement(settings, context, s))
}

func PrintStatement(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string {
	entry, ok := settings.Mappings.PrintStatement[s.Type]
	if ok {
		return entry(settings, context, s)
	}

	return "??"
}
