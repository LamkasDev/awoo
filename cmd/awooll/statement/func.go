package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooParserStatementDataFunc struct {
	Identifier node.AwooParserNode
	Body       AwooParserStatement
}

func GetStatementFuncIdentifier(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataFunc).Identifier
}

func SetStatementFuncIdentifier(s *AwooParserStatement, identifier node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataFunc)
	d.Identifier = identifier
	s.Data = d
}

func GetStatementFuncBody(s *AwooParserStatement) AwooParserStatement {
	return s.Data.(AwooParserStatementDataFunc).Body
}

func SetStatementFuncBody(s *AwooParserStatement, body AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataFunc)
	d.Body = body
	s.Data = d
}

func CreateStatementFunc(identifier node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeFunc,
		Data: AwooParserStatementDataFunc{
			Identifier: identifier,
		},
	}
}
