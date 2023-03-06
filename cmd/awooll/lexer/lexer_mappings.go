package lexer

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooPrintNode func(clexer *AwooLexer, n *node.AwooParserNode) string

type AwooLexerMappings struct {
	PrintNode map[uint16]AwooPrintNode
}
