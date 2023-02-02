package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
)

type AwooParserStatementDataIf struct {
	Value node.AwooParserNode
}

func CreateStatementIf(value node.AwooParserNode) AwooParserStatement {
	return AwooParserStatement{
		Type: ParserStatementTypeDefinitionVariable,
		Data: AwooParserStatementDataDefinitionVariable{
			Value: value,
		},
	}
}
