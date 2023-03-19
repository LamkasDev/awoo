package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooParserNodeDataType struct {
	Value types.AwooTypeId
}

func GetNodeTypeType(n *AwooParserNode) types.AwooTypeId {
	return n.Data.(AwooParserNodeDataType).Value
}

func SetNodeTypeType(n *AwooParserNode, value types.AwooTypeId) {
	d := n.Data.(AwooParserNodeDataType)
	d.Value = value
	n.Data = d
}

func CreateNodeType(t lexer_token.AwooLexerToken) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeType,
			Token: t,
			Data: AwooParserNodeDataType{
				Value: lexer_token.GetTokenTypeId(&t),
			},
		},
	}
}
