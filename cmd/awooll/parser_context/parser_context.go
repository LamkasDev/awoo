package parser_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
)

type AwooParserContext struct {
	Lexer     lexer_context.AwooLexerContext
	Variables map[string]AwooParserContextVariable
	Functions map[string]AwooParserContextFunction
}

type AwooParserContextVariable struct {
	Name string
	Type uint16
}

type AwooParserContextFunction struct {
	Name string
}

func GetContextVariable(context *AwooParserContext, name string) (AwooParserContextVariable, bool) {
	variable, ok := context.Variables[name]
	return variable, ok
}

func SetContextVariable(context *AwooParserContext, variable AwooParserContextVariable) {
	context.Variables[variable.Name] = variable
}

func GetContextFunction(context *AwooParserContext, name string) (AwooParserContextFunction, bool) {
	function, ok := context.Functions[name]
	return function, ok
}

func SetContextFunction(context *AwooParserContext, function AwooParserContextFunction) {
	context.Functions[function.Name] = function
}
