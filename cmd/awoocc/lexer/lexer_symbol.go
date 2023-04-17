package lexer

import (
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

func CreateTokenCouple(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, bool) {
	tokenPosition := lexer.Current.Position
	matchedString := ConstructChunkFast(lexer, string(lexer.Current.Character), func(c rune) bool {
		return unicode.IsPunct(c) || unicode.IsSymbol(c)
	})
	tokenPosition = lexer_token.ExtendAwooLexerTokenPosition(tokenPosition, lexer_token.AwooLexerTokenPosition{
		Length: uint32(len(matchedString)) - tokenPosition.Length,
	})
	couple, ok := lexer.Settings.Tokens.Couple[matchedString]
	if ok {
		return lexer_token.NewAwooLexerToken(tokenPosition, couple), matchedString, true
	}

	return lexer_token.AwooLexerToken{}, matchedString, false
}
