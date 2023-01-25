package node

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/jwalton/gchalk"
)

type AwooParserNodeDataIdentifier struct {
	Value string
}

func GetNodeIdentifierValue(n *AwooParserNode) string {
	return n.Data.(AwooParserNodeDataIdentifier).Value
}

func SetNodeIdentifierValue(n *AwooParserNode, value string) {
	d := n.Data.(AwooParserNodeDataIdentifier)
	d.Value = value
	n.Data = d
}

func CreateNodeIdentifierSafe(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken) (AwooParserNode, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	_, ok := parser_context.GetContextVariable(context, identifier)
	if !ok {
		return AwooParserNode{}, fmt.Errorf("unknown identifier %s", gchalk.Red(identifier))
	}
	return CreateNodeIdentifier(t), nil
}

func CreateNodeIdentifier(t lexer_token.AwooLexerToken) AwooParserNode {
	return AwooParserNode{
		Type:  ParserNodeTypeIdentifier,
		Token: t,
		Data: AwooParserNodeDataIdentifier{
			Value: lexer_token.GetTokenIdentifierValue(&t),
		},
	}
}
