package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

func CreateNodeNegative(t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNodeResult {
	return CreateNodeSingle(ParserNodeTypeNegative, t, value)
}
