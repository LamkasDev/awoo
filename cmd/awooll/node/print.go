package node

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/jwalton/gchalk"
)

func GetNodeDataText(context *lexer_context.AwooLexerContext, n *AwooParserNode) string {
	switch n.Type {
	case ParserNodeTypeIdentifier:
		return GetNodeIdentifierValue(n)
	case ParserNodeTypeType:
		return context.Types.All[GetNodeTypeType(n)].Key
	case ParserNodeTypePrimitive:
		return fmt.Sprintf("%v", GetNodePrimitiveValue(n))
	case ParserNodeTypeExpression:
		l := GetNodeExpressionLeft(n)
		r := GetNodeExpressionRight(n)
		if GetNodeExpressionIsBracket(n) {
			return fmt.Sprintf(
				"%s%v %v %v%s",
				gchalk.Red("("),
				GetNodeDataText(context, &l),
				context.Tokens.All[n.Token.Type].Key,
				GetNodeDataText(context, &r),
				gchalk.Red(")"),
			)
		}
		return fmt.Sprintf(
			"(%v %v %v)",
			GetNodeDataText(context, &l),
			context.Tokens.All[n.Token.Type].Key,
			GetNodeDataText(context, &r),
		)
	case ParserNodeTypeNegative:
		v := GetNodeNegativeValue(n)
		return fmt.Sprintf("-%v", GetNodeDataText(context, &v))
	}

	return "??"
}
