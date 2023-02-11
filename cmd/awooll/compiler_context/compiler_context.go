package compiler_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
)

type AwooCompilerContext struct {
	CurrentAddress uint16
	Parser         parser_context.AwooParserContext
	Scopes         AwooCompilerScopeContainer
	Registers      AwooCompilerRegisters
	Functions      AwooCompilerFunctionContainer
}

func GetProgramHeaderSize() uint16 {
	return 4
}
