package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

type AwooParserNodeDataSingle struct {
	Value AwooParserNode
}

func GetNodeSingleValue(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataSingle).Value
}

func SetNodeSingleValue(n *AwooParserNode, value AwooParserNode) {
	d := n.Data.(AwooParserNodeDataSingle)
	d.Value = value
	n.Data = d
}

func CreateNodeSingle(nodeType uint16, t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  nodeType,
			Token: t,
			Data: AwooParserNodeDataSingle{
				Value: value,
			},
		},
	}
}
