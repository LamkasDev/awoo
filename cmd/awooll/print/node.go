package print

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

func GetNodeDataText(context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	switch n.Type {
	case node.ParserNodeTypeIdentifier:
		return node.GetNodeIdentifierValue(n)
	case node.ParserNodeTypeType:
		return context.Tokens.All[node.GetNodeTypeType(n)].Name
	case node.ParserNodeTypePrimitive:
		return fmt.Sprintf("%v", node.GetNodePrimitiveValue(n))
	}

	return "??"
}
