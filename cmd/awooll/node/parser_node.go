package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

const ParserNodeTypeIdentifier = 0x0000
const ParserNodeTypeType = 0x0001
const ParserNodeTypePrimitive = 0x0002
const ParserNodeTypeExpression = 0x0003
const ParserNodeTypeNegative = 0x0004

type AwooParserNode struct {
	Error error
	Type  uint16
	Token lexer_token.AwooLexerToken
	Data  interface{}
}

type AwooParserNodeResult struct {
	Node      AwooParserNode
	Error     error
	End       bool
	IsBracket bool
}
