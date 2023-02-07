package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

type AwooParserNodeDataPrimitive struct {
	Type  uint16
	Value interface{}
}

func GetNodePrimitiveType(n *AwooParserNode) uint16 {
	return n.Data.(AwooParserNodeDataPrimitive).Type
}

func SetNodePrimitiveType(n *AwooParserNode, t uint16) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Type = t
	n.Data = d
}

func GetNodePrimitiveValue(n *AwooParserNode) interface{} {
	return n.Data.(AwooParserNodeDataPrimitive).Value
}

func SetNodePrimitiveValue(n *AwooParserNode, value interface{}) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Value = value
	n.Data = d
}

func CreateNodePrimitive(t lexer_token.AwooLexerToken) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypePrimitive,
			Token: t,
			Data: AwooParserNodeDataPrimitive{
				Type:  lexer_token.GetTokenPrimitiveType(&t),
				Value: lexer_token.GetTokenPrimitiveValue(&t),
			},
		},
	}
}
