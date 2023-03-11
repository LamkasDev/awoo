package lexer

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/jwalton/gchalk"
)

func PrintNodeIdentifier(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	return node.GetNodeIdentifierValue(n)
}

func PrintNodeType(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	return context.Types.All[node.GetNodeTypeType(n)].Key
}

func PrintNodePointer(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	s := node.GetNodeSingleValue(n)
	return fmt.Sprintf("*%s", PrintNode(settings, context, &s))
}

func PrintNodePrimitive(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	return fmt.Sprintf("%v", node.GetNodePrimitiveValue(n))
}

func PrintNodeExpression(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	l := node.GetNodeExpressionLeft(n)
	r := node.GetNodeExpressionRight(n)
	if node.GetNodeExpressionIsBracket(n) {
		return fmt.Sprintf(
			"%s%s %s %s%s",
			gchalk.Red("("),
			PrintNode(settings, context, &l),
			settings.Tokens.All[n.Token.Type].Name,
			PrintNode(settings, context, &r),
			gchalk.Red(")"),
		)
	}
	return fmt.Sprintf(
		"(%s %s %s)",
		PrintNode(settings, context, &l),
		settings.Tokens.All[n.Token.Type].Name,
		PrintNode(settings, context, &r),
	)
}

func PrintNodeNegative(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	v := node.GetNodeSingleValue(n)
	return fmt.Sprintf("-%s", PrintNode(settings, context, &v))
}

func PrintNodeReference(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	v := node.GetNodeSingleValue(n)
	return fmt.Sprintf("&%s", PrintNode(settings, context, &v))
}

func PrintNodeDereference(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	v := node.GetNodeSingleValue(n)
	return fmt.Sprintf("*%s", PrintNode(settings, context, &v))
}

func PrintNodeCall(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	v := node.GetNodeCallValue(n)
	return fmt.Sprintf("%s()", v)
}

func PrintNode(settings *AwooLexerSettings, context *lexer_context.AwooLexerContext, n *node.AwooParserNode) string {
	entry, ok := settings.Mappings.PrintNode[n.Type]
	if ok {
		return entry(settings, context, n)
	}

	return "??"
}
