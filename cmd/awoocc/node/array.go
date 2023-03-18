package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

type AwooParserNodeDataArray struct {
	Elements []AwooParserNode
}

func GetNodeArrayElements(n *AwooParserNode) []AwooParserNode {
	return n.Data.(AwooParserNodeDataArray).Elements
}

func SetNodeArrayElements(n *AwooParserNode, elements []AwooParserNode) {
	d := n.Data.(AwooParserNodeDataArray)
	d.Elements = elements
	n.Data = d
}

func CreateNodeArray(t lexer_token.AwooLexerToken, elements []AwooParserNode) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeArray,
			Token: t,
			Data: AwooParserNodeDataArray{
				Elements: elements,
			},
		},
	}
}
