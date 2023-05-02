package lexer

import (
	"strings"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

func CreateTokenKeyword(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, bool) {
	tokenPosition := lexer.Current.Position
	matchedString := ConstructChunkFast(lexer, string(lexer.Current.Character), func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsNumber(c)
	})
	tokenPosition = lexer_token.ExtendAwooLexerTokenPosition(tokenPosition, lexer_token.AwooLexerTokenPosition{
		Length: uint32(len(matchedString)) - tokenPosition.Length,
	})
	matchingKeyword, ok := lexer.Settings.Tokens.Keywords[strings.ToLower(matchedString)]
	if !ok {
		AdvanceLexerFor(lexer, -int32(tokenPosition.Length)+1)
		return lexer_token.AwooLexerToken{}, matchedString, false
	}

	return lexer_token.NewAwooLexerToken(tokenPosition, matchingKeyword), matchedString, true
}
