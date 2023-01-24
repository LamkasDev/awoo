package parser_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
)

type AwooParserContext struct {
	Lexer     lexer_context.AwooLexerContext
	Variables map[string]AwooParserContextVariable
}

type AwooParserContextVariable struct {
	Name string
}

func GetContextVariable(context *AwooParserContext, name string) (AwooParserContextVariable, bool) {
	v, ok := context.Variables[name]
	return v, ok
}

func SetContextVariable(context *AwooParserContext, variable AwooParserContextVariable) {
	context.Variables[variable.Name] = variable
}
