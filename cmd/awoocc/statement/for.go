package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
)

type AwooParserStatementDataFor struct {
	Initialization AwooParserStatement
	Condition      node.AwooParserNode
	Advancement    AwooParserStatement
	Body           AwooParserStatement
}

func GetStatementForInitialization(s *AwooParserStatement) AwooParserStatement {
	return s.Data.(AwooParserStatementDataFor).Initialization
}

func SetStatementForInitialization(s *AwooParserStatement, initialization AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataFor)
	d.Initialization = initialization
	s.Data = d
}

func GetStatementForCondition(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataFor).Condition
}

func SetStatementForCondition(s *AwooParserStatement, condition node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataFor)
	d.Condition = condition
	s.Data = d
}

func GetStatementForAdvancement(s *AwooParserStatement) AwooParserStatement {
	return s.Data.(AwooParserStatementDataFor).Advancement
}

func SetStatementForAdvancement(s *AwooParserStatement, advancement AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataFor)
	d.Advancement = advancement
	s.Data = d
}

func GetStatementForBody(s *AwooParserStatement) AwooParserStatement {
	return s.Data.(AwooParserStatementDataFor).Body
}

func SetStatementForBody(s *AwooParserStatement, b AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataFor)
	d.Body = b
	s.Data = d
}

func CreateStatementFor(initialization AwooParserStatement) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeFor,
		Data: AwooParserStatementDataFor{
			Initialization: initialization,
		},
	}
}
