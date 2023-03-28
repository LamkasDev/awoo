package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooParserStatementDataFunc struct {
	Identifier node.AwooParserNode
	Arguments  []AwooParserStatementFuncArgument
	ReturnType *node.AwooParserNode
	Body       AwooParserStatement
}

type AwooParserStatementFuncArgument struct {
	Name        string
	Size        arch.AwooRegister
	Type        types.AwooTypeId
	TypeDetails *types.AwooTypeId
}

func GetStatementFuncIdentifier(s *AwooParserStatement) node.AwooParserNode {
	return s.Data.(AwooParserStatementDataFunc).Identifier
}

func SetStatementFuncIdentifier(s *AwooParserStatement, identifier node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataFunc)
	d.Identifier = identifier
	s.Data = d
}

func GetStatementFuncArguments(s *AwooParserStatement) []AwooParserStatementFuncArgument {
	return s.Data.(AwooParserStatementDataFunc).Arguments
}

func SetStatementFuncArguments(s *AwooParserStatement, arguments []AwooParserStatementFuncArgument) {
	d := s.Data.(AwooParserStatementDataFunc)
	d.Arguments = arguments
	s.Data = d
}

func GetStatementFuncReturnType(s *AwooParserStatement) *node.AwooParserNode {
	return s.Data.(AwooParserStatementDataFunc).ReturnType
}

func SetStatementFuncReturnType(s *AwooParserStatement, returnType *node.AwooParserNode) {
	d := s.Data.(AwooParserStatementDataFunc)
	d.ReturnType = returnType
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
			Arguments:  []AwooParserStatementFuncArgument{},
		},
	}
}
