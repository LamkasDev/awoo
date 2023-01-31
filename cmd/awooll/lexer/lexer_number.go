package lexer

import (
	"strconv"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func CreateTokenNumber(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	base, baseMatchedString, baseValidator := ResolveBase(lexer)
	matchedString := ConstructChunk(lexer, string(lexer.Current), baseValidator)
	number, err := strconv.ParseInt(matchedString, base, 64)
	if err != nil {
		return lexer_token.AwooLexerToken{}, matchedString, err
	}

	return lexer_token.CreateTokenPrimitive(lexer.Position, types.AwooTypeInt64, number, base), (baseMatchedString + matchedString), nil
}

func BaseValidator(c rune) bool {
	return unicode.IsNumber(c)
}

func Base16Validator(c rune) bool {
	return unicode.IsNumber(c) || (c >= 'a' && c <= 'f')
}

func ResolveBase(lexer *AwooLexer) (int, string, ConstructChunkValidator) {
	if lexer.Current == '0' {
		c, ok := PeekLexer(lexer)
		if !ok {
			return 10, "", BaseValidator
		}
		switch unicode.ToLower(c) {
		case 'b':
			AdvanceLexerFor(lexer, 2)
			return 2, "0b", BaseValidator
		case 'o':
			AdvanceLexerFor(lexer, 2)
			return 8, "0o", BaseValidator
		case 'x':
			AdvanceLexerFor(lexer, 2)
			return 16, "0x", Base16Validator
		}
	}

	return 10, "", BaseValidator
}
