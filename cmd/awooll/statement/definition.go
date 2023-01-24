package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooParserStatementDataDefinition struct {
	Type       node.AwooParserNode
	Identifier node.AwooParserNode
	Value      node.AwooParserNode
}

func GetStatementDefinitionType(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinition).Type
}

func SetStatementDefinitionType(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinition)
	d.Type = n
	s.Data = d
}

func GetStatementDefinitionIdentifier(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinition).Identifier
}

func SetStatementDefinitionIdentifier(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinition)
	d.Identifier = n
	s.Data = d
}

func GetStatementDefinitionValue(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinition).Value
}

func SetStatementDefinitionValue(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinition)
	d.Value = n
	s.Data = d
}

func CreateStatementDefinition(n node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeDefinition,
		Data: AwooParserStatementDataDefinition{
			Type: n,
		},
	}
}
