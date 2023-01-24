package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooParserStatementDataAssignment struct {
	Identifier node.AwooParserNode
	Value      node.AwooParserNode
}

func GetStatementAssignmentIdentifier(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataAssignment).Identifier
}

func SetStatementAssignmentIdentifier(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataAssignment)
	d.Identifier = n
	s.Data = d
}

func GetStatementAssignmentValue(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataAssignment).Value
}

func SetStatementAssignmentValue(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataAssignment)
	d.Value = n
	s.Data = d
}

func CreateStatementAssignment(n node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeAssignment,
		Data: AwooParserStatementDataAssignment{
			Identifier: n,
		},
	}
}
