package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

func CreateNodeNegative(t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNodeResult {
	return CreateNodeSingle(ParserNodeTypeNegative, t, value)
}
