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
		return context.Types.All[node.GetNodeTypeType(n)].Key
	case node.ParserNodeTypePrimitive:
		return fmt.Sprintf("%v", node.GetNodePrimitiveValue(n))
	case node.ParserNodeTypeExpression:
		l := node.GetNodeExpressionLeft(n)
		r := node.GetNodeExpressionRight(n)
		return fmt.Sprintf(
			"(%v %v %v)",
			GetNodeDataText(context, &l),
			context.Tokens.All[n.Token.Type].Key,
			GetNodeDataText(context, &r),
		)
	case node.ParserNodeTypeNegative:
		v := node.GetNodeNegativeValue(n)
		return fmt.Sprintf("-(%v)", GetNodeDataText(context, &v))
	}

	return "??"
}
