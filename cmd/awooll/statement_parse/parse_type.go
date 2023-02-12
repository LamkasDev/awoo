package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructNodeType(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) node.AwooParserNodeResult {
	n := node.CreateNodeType(t)
	for t, ok := parser.PeekParser(cparser); ok && t.Type == token.TokenOperatorDereference; t, ok = parser.PeekParser(cparser) {
		parser.AdvanceParser(cparser)
		n = node.CreateNodePointer(t, n.Node)
	}

	return n
}
