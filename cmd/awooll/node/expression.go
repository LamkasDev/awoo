package node

import "github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"

type AwooParserNodeDataExpression struct {
	Left  AwooParserNode
	Right AwooParserNode
}

func GetNodeExpressionLeft(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataExpression).Left
}

func SetNodeExpressionLeft(n *AwooParserNode, value AwooParserNode) {
	d := n.Data.(AwooParserNodeDataExpression)
	d.Left = value
	n.Data = d
}

func GetNodeExpressionRight(n *AwooParserNode) AwooParserNode {
	return n.Data.(AwooParserNodeDataExpression).Right
}

func SetNodeExpressionRight(n *AwooParserNode, value AwooParserNode) {
	d := n.Data.(AwooParserNodeDataExpression)
	d.Right = value
	n.Data = d
}

func CreateNodeExpression(t lexer_token.AwooLexerToken, left AwooParserNode, right AwooParserNode) AwooParserNode {
	return AwooParserNode{
		Type:  ParserNodeTypeExpression,
		Token: t,
		Data: AwooParserNodeDataExpression{
			Left:  left,
			Right: right,
		},
	}
}
