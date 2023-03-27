package compiler_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

type AwooCompilerFunctionContainer struct {
	Entries map[string]AwooCompilerFunction
}

type AwooCompilerFunction struct {
	Symbol    elf.AwooElfSymbolTableEntry
	Arguments []statement.AwooParserStatementFuncArgument
}

func PushCompilerFunction(context *AwooCompilerContext, entry AwooCompilerFunction) {
	context.Functions.Entries[entry.Symbol.Name] = entry
}

func SetSizeOfCompilerFunction(context *AwooCompilerContext, name string, size uint32) {
	entry := context.Functions.Entries[name]
	entry.Symbol.Size = size
	context.Functions.Entries[name] = entry
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
