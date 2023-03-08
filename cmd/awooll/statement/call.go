package statement

import "github.com/LamkasDev/awoo-emu/cmd/awooll/node"

type AwooParserStatementDataCall struct {
	Node node.AwooParserNode
}

func GetStatementCallNode(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataCall).Node
}

func SetStatementCallNode(s *AwooParserStatement, node node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataCall)
	d.Node = node
	s.Data = d
}

func CreateStatementCall(node node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeCall,
		Data: AwooParserStatementDataCall{
			Node: node,
		},
	}
}
