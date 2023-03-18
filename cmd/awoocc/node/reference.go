package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

func CreateNodeReference(t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNodeResult {
	return CreateNodeSingle(ParserNodeTypeReference, t, value)
}

func CreateNodeDereference(t lexer_token.AwooLexerToken, value AwooParserNode) AwooParserNodeResult {
	return CreateNodeSingle(ParserNodeTypeDereference, t, value)
}
