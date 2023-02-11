package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooParserStatementDataIf struct {
	Value node.AwooParserNode
	Body  AwooParserStatement
	Else  []AwooParserStatement
}

func GetStatementIfValue(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataIf).Value
}

func SetStatementIfValue(s *AwooParserStatement, n node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataIf)
	d.Value = n
	s.Data = d
}

func GetStatementIfBody(s *AwooParserStatement) AwooParserStatement {
	return s.Data.(AwooParserStatementDataIf).Body
}

func SetStatementIfBody(s *AwooParserStatement, b AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataIf)
	d.Body = b
	s.Data = d
}

func GetStatementIfElse(s *AwooParserStatement) []AwooParserStatement {
	return s.Data.(AwooParserStatementDataIf).Else
}

func SetStatementIfElse(s *AwooParserStatement, next []AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataIf)
	d.Else = next
	s.Data = d
}

func CreateStatementIf(value node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeIf,
		Data: AwooParserStatementDataIf{
			Value: value,
			Else:  []AwooParserStatement{},
		},
	}
}
