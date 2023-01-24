package lexer_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

type AwooLexerContext struct {
	Tokens token.AwooTokenMap
	Types  types.AwooTypeMap
}
