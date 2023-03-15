package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

const ParserNodeTypeIdentifier = 0x000
const ParserNodeTypeType = 0x001
const ParserNodeTypePointer = 0x002
const ParserNodeTypePrimitive = 0x003
const ParserNodeTypeExpression = 0x004
const ParserNodeTypeNegative = 0x005
const ParserNodeTypeReference = 0x006
const ParserNodeTypeDereference = 0x007
const ParserNodeTypeCall = 0x008
const ParserNodeTypeArray = 0x009
const ParserNodeTypeArrayIndex = 0x00A

type AwooParserNode struct {
	Type  uint16
	Token lexer_token.AwooLexerToken
	Data  interface{}
}

type AwooParserNodeResult struct {
	Node       AwooParserNode
	End        bool
	EndBracket bool
}
