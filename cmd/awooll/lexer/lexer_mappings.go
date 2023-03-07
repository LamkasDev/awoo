package lexer

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooPrintNode func(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string

type AwooLexerMappings struct {
	PrintNode map[uint16]AwooPrintNode
}
