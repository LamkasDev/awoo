package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooParserStatementDataDefinitionVariable struct {
	Type       node.AwooParserNode
	Identifier node.AwooParserNode
	Value      *node.AwooParserNode
}

func GetStatementDefinitionVariableType(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinitionVariable).Type
}

func SetStatementDefinitionVariableType(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinitionVariable)
	d.Type = n
	s.Data = d
}

func GetStatementDefinitionVariableIdentifier(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinitionVariable).Identifier
}

func SetStatementDefinitionVariableIdentifier(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinitionVariable)
	d.Identifier = n
	s.Data = d
}

func GetStatementDefinitionVariableValue(s *AwooParserStatement) *node.AwooParserNode {
	return s.Data.(AwooParserStatementDataDefinitionVariable).Value
}

func SetStatementDefinitionVariableValue(s *AwooParserStatement, n *node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataDefinitionVariable)
	d.Value = n
	s.Data = d
}

func CreateStatementDefinitionVariable(variableType node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeDefinitionVariable,
		Data: AwooParserStatementDataDefinitionVariable{
			Type: variableType,
		},
	}
}
