package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooParserStatementDataFunc struct {
	Identifier node.AwooParserNode
	Arguments  []elf.AwooElfSymbolTableEntry
	ReturnType *node.AwooParserNode
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

func GetStatementFuncArguments(s *AwooParserStatement) []elf.AwooElfSymbolTableEntry {
	return s.Data.(AwooParserStatementDataFunc).Arguments
}

func SetStatementFuncArguments(s *AwooParserStatement, arguments []elf.AwooElfSymbolTableEntry) {
	d := s.Data.(AwooParserStatementDataFunc)
	d.Arguments = arguments
	s.Data = d
}

func GetStatementFuncReturnType(s *AwooParserStatement) *node.AwooParserNode {
	return s.Data.(AwooParserStatementDataFunc).ReturnType
}

func GetStatementFuncReturnTypePrecise(s *AwooParserStatement) *types.AwooTypeId {
	tNode := GetStatementFuncReturnType(s)
	if tNode != nil {
		t := node.GetNodeTypeType(tNode)
		return &t
	}

	return nil
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

func AppendStatementFuncBody(s *AwooParserStatement, statement AwooParserStatement) {
	d := s.Data.(AwooParserStatementDataFunc)
	SetStatementGroupBody(&d.Body, append(GetStatementGroupBody(&d.Body), statement))
	s.Data = d
}

func CreateStatementFunc(identifier node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeFunc,
		Data: AwooParserStatementDataFunc{
			Identifier: identifier,
			Arguments:  []elf.AwooElfSymbolTableEntry{},
			Body:       CreateStatementGroup([]AwooParserStatement{}),
		},
	}
}
