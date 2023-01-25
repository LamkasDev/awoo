package compiler_context

import "github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"

type AwooCompilerContext struct {
	Parser parser_context.AwooParserContext
	Scopes AwooCompilerScopeContainer
}
