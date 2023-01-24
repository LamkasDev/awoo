package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

const ParserNodeTypeIdentifier = 0x0000
const ParserNodeTypeType = 0x0001
const ParserNodeTypePrimitive = 0x0002

type AwooParserNode struct {
	Error error
	Type  uint16
	Token lexer_token.AwooLexerToken
	Data  interface{}
}
