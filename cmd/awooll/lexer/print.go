package lexer

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

func PrintNode(context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	entry, ok := clexer.Settings.Mappings.PrintNode[n.Type]
	if ok {
		return entry(context, n)
	}

	return "??"
}
