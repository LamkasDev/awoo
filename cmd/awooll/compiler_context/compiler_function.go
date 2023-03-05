package compiler_context

import "github.com/LamkasDev/awoo-emu/cmd/awooll/statement"

type AwooCompilerFunctionContainer struct {
	Start   string
	Entries map[string]AwooCompilerFunction
}

type AwooCompilerFunction struct {
	Name      string
	Start     uint16
	Size      uint16
	Arguments []statement.AwooParserStatementFuncArgument
}

func PushCompilerFunction(context *AwooCompilerContext, entry AwooCompilerFunction) {
	context.Functions.Entries[entry.Name] = entry
}

func GetCompilerFunction(context *AwooCompilerContext, name string) (AwooCompilerFunction, bool) {
	f, ok := context.Functions.Entries[name]
	return f, ok
}
