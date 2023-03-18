package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

type AwooParserNodeDataTypeArray struct {
	Type AwooParserNode
	Size uint32
}

func GetNodeTypeArrayType(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataTypeArray).Type
}

func SetNodeTypeArrayType(n *AwooParserNode, arrType AwooParserNode) {
	d := n.Data.(AwooParserNodeDataTypeArray)
	d.Type = arrType
	n.Data = d
}

func GetNodeTypeArraySize(n *AwooParserNode) uint32 {
	return n.Data.(AwooParserNodeDataTypeArray).Size
}

func SetNodeTypeArraySize(n *AwooParserNode, size uint32) {
	d := n.Data.(AwooParserNodeDataTypeArray)
	d.Size = size
	n.Data = d
}

func CreateNodeTypeArray(t lexer_token.AwooLexerToken, arrType AwooParserNode) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeTypeArray,
			Token: t,
			Data: AwooParserNodeDataTypeArray{
				Type: arrType,
			},
		},
	}
}
