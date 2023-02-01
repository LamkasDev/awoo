package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

// Value is NodeType
type AwooParserStatementDataDefinitionType struct {
	Identifier node.AwooParserNode
	Value      node.AwooParserNode
}

func GetStatementDefinitionTypeIdentifier(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinitionType).Identifier
}

func SetStatementDefinitionTypeIdentifier(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinitionType)
	d.Identifier = n
	s.Data = d
}

func GetStatementDefinitionTypeValue(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinitionType).Value
}

func SetStatementDefinitionTypeValue(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinitionType)
	d.Value = n
	s.Data = d
}

func CreateStatementDefinitionType(identifier node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeDefinitionType,
		Data: AwooParserStatementDataDefinitionType{
			Identifier: identifier,
		},
	}
}
