package lexer

import (
	"strings"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
)

func CreateTokenLetter(lexer *AwooLexer) (lexer_token.AwooLexerToken, string) {
	matchedString := ConstructChunk(lexer, string(lexer.Current), func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsNumber(c)
	})
	matchingKeyword, ok := lexer.Context.Tokens.Keywords[strings.ToLower(matchedString)]
	if ok {
		return lexer_token.CreateToken(lexer.Position, matchingKeyword), matchedString
	}
	matchingType, ok := lexer.Context.Types.Lookup[matchedString]
	if ok {
		return lexer_token.CreateTokenType(lexer.Position, matchingType.Type), matchedString
	}

	return lexer_token.CreateTokenIdentifier(lexer.Position, matchedString), matchedString
}
