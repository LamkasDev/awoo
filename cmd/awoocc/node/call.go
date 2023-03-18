package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

type AwooParserNodeDataCall struct {
	Value     string
	Arguments []AwooParserNode
}

func GetNodeCallValue(n *AwooParserNode) string {
	return n.Data.(AwooParserNodeDataCall).Value
}

func SetNodeCallValue(n *AwooParserNode, value string) {
	d := n.Data.(AwooParserNodeDataCall)
	d.Value = value
	n.Data = d
}

func GetNodeCallArguments(n *AwooParserNode) []AwooParserNode {
	return n.Data.(AwooParserNodeDataCall).Arguments
}

func SetNodeCallArguments(n *AwooParserNode, arguments []AwooParserNode) {
	d := n.Data.(AwooParserNodeDataCall)
	d.Arguments = arguments
	n.Data = d
}

func CreateNodeCall(t lexer_token.AwooLexerToken) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypeCall,
			Token: t,
			Data: AwooParserNodeDataCall{
				Value:     lexer_token.GetTokenIdentifierValue(&t),
				Arguments: []AwooParserNode{},
			},
		},
	}
}
