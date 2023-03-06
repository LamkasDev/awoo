package statement

import "github.com/LamkasDev/awoo-emu/cmd/awooll/node"

type AwooParserStatementDataReturn struct {
	Value *node.AwooParserNode
}

func GetStatementReturnValue(s *AwooParserStatement) *node.AwooParserNode {
	return s.Data.(AwooParserStatementDataReturn).Value
}

func SetStatementReturnValue(s *AwooParserStatement, value *node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataReturn)
	d.Value = value
	s.Data = d
}

func CreateStatementReturn(value *node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeReturn,
		Data: AwooParserStatementDataReturn{
			Value: value,
		},
	}
}
