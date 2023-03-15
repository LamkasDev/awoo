package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

type AwooParserNodeDataArrayIndex struct {
	Identifier string
	Index      AwooParserNode
}

func GetNodeArrayIndexIdentifier(n *AwooParserNode) string {
	return n.Data.(AwooParserNodeDataArrayIndex).Identifier
}

func SetNodeArrayIndexIdentifier(n *AwooParserNode, identifier string) {
	d := n.Data.(AwooParserNodeDataArrayIndex)
	d.Identifier = identifier
	n.Data = d
}

func GetNodeArrayIndexIndex(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataArrayIndex).Index
}

func SetNodeArrayIndexIndex(n *AwooParserNode, index AwooParserNode) {
	d := n.Data.(AwooParserNodeDataArrayIndex)
	d.Index = index
	n.Data = d
}

func CreateNodeArrayIndex(t lexer_token.AwooLexerToken, identifier string) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeArrayIndex,
			Token: t,
			Data: AwooParserNodeDataArrayIndex{
				Identifier: identifier,
			},
		},
	}
}
