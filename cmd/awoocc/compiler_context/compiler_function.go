package compiler_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooCompilerFunctionContainer struct {
	Entries map[string]AwooCompilerFunction
}

type AwooCompilerFunction struct {
	Name       string
	ReturnType *types.AwooTypeId
	Arguments  []statement.AwooParserStatementFuncArgument
	Start      uint32
}

func PushCompilerFunction(context *AwooCompilerContext, entry AwooCompilerFunction) {
	context.Functions.Entries[entry.Name] = entry
}

func GetCompilerFunction(context *AwooCompilerContext, name string) (AwooCompilerFunction, bool) {
	f, ok := context.Functions.Entries[name]
	return f, ok
}

func GetCompilerFunctionArgumentsSize(function AwooCompilerFunction) uint32 {
	offset := uint32(0)
	for _, argument := range function.Arguments {
		offset += argument.Size
	}
	return offset
}
