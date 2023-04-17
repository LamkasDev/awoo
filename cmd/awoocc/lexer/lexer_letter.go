package lexer

import (
	"strings"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
)

func CreateTokenLetter(lexer *AwooLexer) (lexer_token.AwooLexerToken, string) {
	tokenPosition := lexer.Current.Position
	matchedString := ConstructChunkFast(lexer, string(lexer.Current.Character), func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsNumber(c)
	})
	tokenPosition = lexer_token.ExtendAwooLexerTokenPosition(tokenPosition, lexer_token.AwooLexerTokenPosition{
		Length: uint32(len(matchedString)) - tokenPosition.Length,
	})
	matchingKeyword, ok := lexer.Settings.Tokens.Keywords[strings.ToLower(matchedString)]
	if ok {
		return lexer_token.NewAwooLexerToken(tokenPosition, matchingKeyword), matchedString
	}

	return lexer_token.CreateTokenIdentifier(tokenPosition, matchedString), matchedString
}
