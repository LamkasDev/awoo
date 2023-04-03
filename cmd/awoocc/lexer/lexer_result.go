package lexer

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

type AwooLexerResult struct {
	Error   error
	Text    []rune
	Context lexer_context.AwooLexerContext
	Tokens  []lexer_token.AwooLexerToken
}
