package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
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

func CreateNodeIdentifier(t lexer_token.AwooLexerToken) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeIdentifier,
			Token: t,
			Data: AwooParserNodeDataIdentifier{
				Value: lexer_token.GetTokenIdentifierValue(&t),
			},
		},
	}
}
