package parser_context

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"

type AwooParserFunctionContainer struct {
	Entries map[string]AwooParserFunction
	Start   string
}

type AwooParserFunction struct {
	Name       string
	ReturnType *uint16
	Arguments  []statement.AwooParserStatementFuncArgument
}

func PushParserFunction(context *AwooParserContext, entry AwooParserFunction) {
	context.Functions.Entries[entry.Name] = entry
}

func GetParserFunction(context *AwooParserContext, name string) (AwooParserFunction, bool) {
	f, ok := context.Functions.Entries[name]
	return f, ok
}
