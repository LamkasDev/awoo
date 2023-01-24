package lexer

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

type AwooLexerResult struct {
	Error   error
	Context lexer_context.AwooLexerContext
	Tokens  []lexer_token.AwooLexerToken
}
