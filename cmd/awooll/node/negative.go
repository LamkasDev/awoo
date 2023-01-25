package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

type AwooParserNodeDataNegative struct {
	Value AwooParserNode
}

func GetNodeNegativeValue(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataNegative).Value
}

func SetNodeNegativeValue(n *AwooParserNode, value AwooParserNode) {
	d := n.Data.(AwooParserNodeDataNegative)
	d.Value = value
	n.Data = d
}

func CreateNodeNegative(t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNode {
	return AwooParserNode{
		Type:  ParserNodeTypeNegative,
		Token: t,
		Data: AwooParserNodeDataNegative{
			Value: value,
		},
	}
}
