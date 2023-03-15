package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

type AwooParserNodeDataArray struct {
	Type AwooParserNode
	Size uint16
}

func GetNodeArrayType(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataArray).Type
}

func SetNodeArrayType(n *AwooParserNode, arrType AwooParserNode) {
	d := n.Data.(AwooParserNodeDataArray)
	d.Type = arrType
	n.Data = d
}

func GetNodeArraySize(n *AwooParserNode) uint16 {
	return n.Data.(AwooParserNodeDataArray).Size
}

func SetNodeArraySize(n *AwooParserNode, size uint16) {
	d := n.Data.(AwooParserNodeDataArray)
	d.Size = size
	n.Data = d
}

func CreateNodeArray(t lexer_token.AwooLexerToken, arrType AwooParserNode) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeArray,
			Token: t,
			Data: AwooParserNodeDataArray{
				Type: arrType,
			},
		},
	}
}
